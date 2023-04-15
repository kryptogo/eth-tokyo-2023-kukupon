import React from "react";
import Image from "next/image";
import { ConnectButton } from "@rainbow-me/rainbowkit";

const redeem = () => {
  return (
    <div className="w-full h-screen bg-primaryPink text-secondaryPink">
      <div className="w-[314px] h-[518px] absolute right-0 top-1/2 -translate-y-1/2 scale-75 origin-right z-2">
        <Image
          src="/images/cutee-sitting.png"
          alt="cutee"
          fill
          className="object-cover"
        />
      </div>
      <div className="w-[501px] h-[665px] absolute left-0 bottom-0 scale-75 origin-bottom-left z-0">
        <Image
          src="/images/cutee-squashing.png"
          alt="cutee"
          fill
          className="object-cover"
        />
      </div>

      <header className="container mx-auto h-[70px] flex items-center justify-between px-4">
        <div className="flex items-center gap-4">
          <div className="w-[47px] h-[47px] relative">
            <Image
              src="/images/donki-logo.png"
              alt="donki-logo"
              fill
              className="object-cover"
            />
          </div>
          <h1 className="text-[24px]">DONDONKI</h1>
        </div>

        <ConnectButton />
      </header>

      <main className="container h-[calc(100vh-70px)] flex flex-col items-center justify-center mx-auto text-center relative z-1 space-y-6">
        <div>
          <h1 className="text-[84px] font-black">DONDONKI</h1>
          <h1 className="text-[84px] font-black">CUTEEEEE GAMEFI</h1>
        </div>

        <div className="relative">
          <input
            type="text"
            className="rounded-[10px] w-[457px] h-[81px] bg-[#D9D9D9] pl-[85px]"
          />
          <Image
            src="/images/kukuponpon.png"
            width={72}
            height={72}
            alt="kukuponpon"
            className="absolute -translate-y-1/2 left-4 top-1/2"
          />
        </div>

        <button className="w-[142px] h-[37px] rounded-[10px] bg-secondaryPink text-white">
          Play Now
        </button>
      </main>
    </div>
  );
};

export default redeem;
