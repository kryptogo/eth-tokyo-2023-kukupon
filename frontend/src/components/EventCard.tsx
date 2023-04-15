import React, { useEffect, useState } from "react";
import dynamic from "next/dynamic";
import { useIDKit } from "@worldcoin/idkit";
import { useSignMessage } from "wagmi";
import { useMutation } from "@tanstack/react-query";

import Button from "./Button";

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
  const {
    isSuccess: signMessageSuccess,
    isLoading: signMessageLoading,
    signMessage,
  } = useSignMessage({
    message: id,
  });

  const {
    mutate: verify,
    isLoading: verifyLoading,
    isSuccess: verifySuccess,
  } = useMutation((data) =>
    fetch(`${process.env.NEXT_PUBLIC_API_URL}/verify_campaign`, {
      method: "POST",
    })
  );

  const handleSuccess = () => {};
  const handleClick = () => {};

  useEffect(() => {
    if (!signMessageSuccess) return;
  }, [signMessageSuccess]);

  return (
    <div
      className="w-[280px] h-[394px] border border-white rounded-[10px] flex items-center justify-center"
      style={{
        backgroundImage: `url(${backgroundImage})`,
      }}
    >
      <div className="flex flex-col items-center gap-[120px]">
        <div className="flex flex-col items-center gap-5">
          <h2 className="text-[42px] font-bold">{title}</h2>
          <ul>
            {description.map((item, index) => (
              <li key={index}>{item}</li>
            ))}
          </ul>
        </div>
        <Button text="Verify" onClick={handleClick} />

        <IDKitWidget
          app_id={process.env.NEXT_PUBLIC_APP_ID!}
          action={id}
          onSuccess={handleSuccess}
        />
      </div>
    </div>
  );
};

export default EventCard;
