import Image from "next/image";
import { Inter } from "next/font/google";
import { ConnectButton } from "@rainbow-me/rainbowkit";

import Button from "@/components/Button";
import EventCard from "@/components/EventCard";

import kukuponpon from "/public/images/kukuponpon.png";
import ape from "/public/images/ape.png";

export default function Home() {
  return (
    <div className="w-full h-screen bg-[#212121] bg-anime bg-no-repeat bg-right">
      <div className="container mx-auto h-full">
        <header className="h-[100px] flex justify-end items-center">
          <ConnectButton />
        </header>
        <main className="h-[calc(100vh-100px)] grid grid-cols-2 items-stretch">
          <div className="w-full h-full flex flex-col justify-center gap-[120px] max-w-[500px]">
            <div className="text-white font-cartoon">
              <div className="font-bold">
                <div className="flex items-center justify-cneter">
                  <h1 className="text-[75px] leading-tight">KUKU</h1>
                  <Image
                    src={kukuponpon}
                    alt="kukuponpon logo"
                    className="w-[145px]"
                  />
                </div>
                <h1 className="text-[75px] leading-tight">PONPON</h1>
              </div>

              <h2 className="text-[24px]">
                Maximize your gamefi time and minimize gas fees with our coupons
              </h2>
            </div>
            <Button text="Become a PON Master" />
          </div>
        </main>
      </div>

      <div className="w-full h-screen py-6">
        <div className="container mx-auto grid grid-flow-row grid-cols-4 gap-[20px]">
          <EventCard
            title="123"
            description={["123", "456"]}
            backgroundImage={ape.src}
          />
          <EventCard
            title="123"
            description={["123", "456"]}
            backgroundImage={ape.src}
          />
          <EventCard
            title="123"
            description={["123", "456"]}
            backgroundImage={ape.src}
          />
          <EventCard
            title="123"
            description={["123", "456"]}
            backgroundImage={ape.src}
          />
          <EventCard
            title="123"
            description={["123", "456"]}
            backgroundImage={ape.src}
          />
          <EventCard
            title="123"
            description={["123", "456"]}
            backgroundImage={ape.src}
          />
          <EventCard
            title="123"
            description={["123", "456"]}
            backgroundImage={ape.src}
          />
        </div>
      </div>
    </div>
  );
}
