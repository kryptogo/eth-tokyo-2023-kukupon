import { ConnectButton } from '@rainbow-me/rainbowkit';
import Image from 'next/image';
import { useAccount } from 'wagmi';

import Button from '@/components/Button';
import EventCard from '@/components/EventCard';

import { useCampaigns } from '@/service/campaign';
import kukuponpon from '/public/images/kukuponpon.png';

export default function Home() {
  const { address, isConnecting, isDisconnected } = useAccount();
  const { data: campaigns, isLoading } = useCampaigns();

  console.log(campaigns);

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
            <Button onClick={scrollToNext} text="Become a PON Master" />
          </div>
        </main>
      </div>

      <div className="w-full h-screen py-6">
        <div className="container mx-auto grid grid-flow-row grid-cols-4 gap-[20px]">
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
