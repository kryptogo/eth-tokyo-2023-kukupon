import { useMutation } from "@tanstack/react-query";
import { useIDKit } from "@worldcoin/idkit";
import dynamic from "next/dynamic";
import React, { useEffect } from "react";
import { useAccount, useSignMessage } from "wagmi";

import useCouponStore from "@/store/useCouponStore";
import Button from "./Button";
import { useRouter } from "next/router";

interface IEventCard {
  id: string;
  title: string;
  description: string[];
  backgroundImage: string;
}

const IDKitWidget = dynamic(
  // @ts-ignore
  () => import("@worldcoin/idkit").then((mod) => mod.IDKitWidget),
  { ssr: false }
);

const EventCard: React.FC<IEventCard> = ({
  id,
  title,
  description,
  backgroundImage,
}) => {
  const { open, setOpen } = useIDKit();
  const router = useRouter();
  const { address, isConnecting, isDisconnected } = useAccount();
  const [addCoupons] = useCouponStore((state) => [state.addCoupons]);

  const {
    data: signature,
    isSuccess: signMessageSuccess,
    isLoading: signMessageLoading,
    signMessage,
  } = useSignMessage({
    message: id,
  });

  const {
    data,
    mutate: verify,
    isLoading: verifyLoading,
    isSuccess: verifySuccess,
  } = useMutation(
    (data: { campaign_id: string; from: string; signature: string }) =>
      fetch(`${process.env.NEXT_PUBLIC_API_URL}/verify_campaign`, {
        method: "POST",
        body: JSON.stringify(data),
      })
  );

  const {
    data: coupons,
    mutate: getCoupons,
    isLoading: getCouponLoading,
    isSuccess: getCouponSuccess,
  } = useMutation((data: { campaign_id: string }) =>
    fetch(`${process.env.NEXT_PUBLIC_API_URL}/get_coupon`, {
      method: "POST",
      body: JSON.stringify(data),
      credentials: "include",
    }).then((res) => res.json())
  );

  const handleSuccess = () => {
    getCoupons({
      campaign_id: id,
    });
  };
  const handleClick = () => {
    signMessage();
  };

  useEffect(() => {
    if (!signMessageSuccess) return;

    const payload = {
      campaign_id: id,
      from: address as string,
      signature: signature as string,
    };

    verify(payload);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [signMessageSuccess]);

  useEffect(() => {
    if (!verifySuccess || !data.ok) return;

    setOpen(true);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [verifySuccess]);

  useEffect(() => {
    if (!getCouponSuccess) return;

    addCoupons(title, coupons.coupons);
    router.push("/dashboard");

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [getCouponSuccess]);

  return (
    <div
      className="grayscaleBg w-full h-[450px] border border-white rounded-lg flex items-center justify-center bg-cover bg-no-repeat p-4 hover:scale-[1.02] transiton-transform ease-in-out"
      style={{
        backgroundImage: `url(${backgroundImage})`,
      }}
    >
      <div className="flex flex-col items-center justify-between h-full p-4">
        <div className="flex flex-col items-center gap-5">
          <h2 className="text-[42px] font-bold text-white">{title}</h2>
          {description.map((item, index) => (
            <div
              key={item}
              className="flex items-start gap-2 text-lg text-white"
            >
              <span>－</span>
              <span key={index}>{item}</span>
            </div>
          ))}
        </div>
        <div className="flex flex-col items-center space-y-4 text-center">
          {signMessageLoading && (
            <p className="text-white">Waiting for signing... </p>
          )}
          {verifyLoading && (
            <p className="text-white">Checking your qualifications...</p>
          )}

          <Button
            text="Verify"
            onClick={handleClick}
            status={signMessageLoading || verifyLoading ? "loading" : "normal"}
          />

          <IDKitWidget
            app_id={process.env.NEXT_PUBLIC_APP_ID!}
            action={id}
            onSuccess={handleSuccess}
          />
        </div>
      </div>
    </div>
  );
};

export default EventCard;
