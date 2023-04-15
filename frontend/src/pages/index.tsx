import { ConnectButton } from "@kryptogo/kryptogokit";
import { useAccount } from "wagmi";

import Button from "@/components/Button";
import EventCard from "@/components/EventCard";
import { LottieComponent } from "@/components/LottieLogo";
import { useCampaigns } from "@/service/campaign";

export default function Home() {
  const { data: campaigns, isLoading } = useCampaigns();

  function scrollToNext() {
    const duration = 300; // duration of the animation in milliseconds
    const targetPosition = window.scrollY + window.innerHeight; // calculate the target position of the scroll

    function easeInOutQuad(t: number) {
      // easing function for smooth animation
      return t < 0.5 ? 2 * t * t : -1 + (4 - 2 * t) * t;
    }

    function animateScroll(timestamp: number) {
      const elapsedTime = timestamp - startTime;
      const newPosition =
        easeInOutQuad(elapsedTime / duration) *
          (targetPosition - startPosition) +
        startPosition;
      window.scrollTo(0, newPosition);

      if (elapsedTime < duration) {
        window.requestAnimationFrame(animateScroll);
      }
    }

    const startTime = performance.now();
    const startPosition = window.scrollY;
    window.requestAnimationFrame(animateScroll);
  }

  return (
    <div className="w-full h-screen bg-[#212121] bg-anime bg-contain bg-no-repeat bg-right">
      <div className="container h-full mx-auto">
        <header className="h-[100px] flex justify-end items-center">
          <ConnectButton />
        </header>
        <main className="h-[calc(100vh-100px)] grid grid-cols-1 justify-items-center lg:justify-items-start  text-center lg:text-left lg:grid-cols-2 items-stretch">
          <div className="w-full h-full flex flex-col justify-center gap-[120px] max-w-[300px] md:max-w-[500px]">
            <div className="text-white font-cartoon">
              <div className="font-bold">
                <div className="flex items-center justify-cneter">
                  <LottieComponent></LottieComponent>
                </div>
              </div>

              <h2 className="text-[24px]">
                Where Gift Cards Meet Blockchain Gaming
              </h2>
            </div>
            <Button onClick={scrollToNext} text="Claim Your First Coupon" />
          </div>
        </main>
      </div>

      <div className="w-full h-screen py-6w-full min-h-screen py-6">
        <div className="container mx-auto grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 lg:gap-6 xl:gap-8">
          {campaigns?.map((campaign) => (
            <EventCard
              id={campaign.id}
              key={campaign.id}
              title={campaign.sponsor}
              description={campaign.required_condition}
              backgroundImage={campaign.image}
            />
          ))}
        </div>
      </div>
    </div>
  );
}
