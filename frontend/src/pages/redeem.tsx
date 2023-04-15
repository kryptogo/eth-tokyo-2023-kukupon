import { ConnectButton } from '@rainbow-me/rainbowkit';
import Image from 'next/image';

const redeem = () => {
  return (
    <div className="w-full h-screen bg-primaryPink text-secondaryPink text-cartoon">
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
          <h1 className="text-[24px] font-bold font-cartoon">DONDONKI</h1>
        </div>

        <ConnectButton />
      </header>

      <main className="container h-[calc(100vh-70px)] flex flex-col items-center justify-center mx-auto text-center relative z-1 space-y-6">
        <div>
          <h1 className="text-[84px] font-black font-cartoon">DONDONKI</h1>
          <h1 className="text-[84px] font-black font-cartoon">
            CUTEEEEE GAMEFI
          </h1>
        </div>

        <div className="relative">
          <input
            type="text"
            className="rounded-[10px] w-[457px] h-[81px] bg-[#D9D9D9] pl-[125px]"
          />
          <Image
            src="/images/kukupon.png"
            width={120}
            height={120}
            alt="kukupon"
            className="absolute -translate-y-1/2 left-4 top-1/2"
          />
        </div>

        <button className="w-[142px] h-[37px] rounded-[10px] bg-secondaryPink text-white">
          Play Now
        </button>
      </main>

      {/* FOR MINT NFT */}
      {/* <main className="container h-[calc(100vh-70px)] flex flex-col items-center justify-center mx-auto text-center relative z-1 space-y-6">
        <div>
          <h1 className="text-[84px] font-black font-cartoon">All Done</h1>
          <h4 className="text-[24px] font-black font-cartoon">
            Just Have fun Try Mint your first **FREE** NFT (for real 😎)
          </h4>
        </div>

        <TapMe></TapMe>

        <button className="w-[401px] h-[88px] font-cartoon rounded-[10px] bg-secondaryPink font-black text-white text-[32px]">
          Mint NFT
        </button>
      </main> */}

      {/* FOR LOADING STATE*/}
      {/* <div className="loading absolute top-0  w-full h-full text-secondaryPink text-cartoon flex items-center justify-center">
        <div className="top-0 absolute bg-black w-full h-full opacity-70"></div>
        <div className="rounded-[10px] bg-white z-10 flex relative w-64 h-72">
          <Image
            src="/images/loading_girl.png"
            width={180}
            height={120}
            style={{ top: -100, left: 40 }}
            alt="kukupon"
            className="absolute "
          />
          <div className="absolute top-12 left-0 w-full h-full flex items-center justify-center font-cartoon font-bold text-center p-4">
            We are creating your wallet. Give me a second I’ll give you a new
            world
          </div>
        </div>
      </div> */}
    </div>
  );
};

export default redeem;
