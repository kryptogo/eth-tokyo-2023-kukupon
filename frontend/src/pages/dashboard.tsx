import Button from '@/components/Button';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
const dashboard = () => {
  //create 10 list item wrap with div
  const cupons: string[] = [
    'coupon code 1',
    'coupon code 2',
    'coupon code 3',
    'coupon code 4',
    'coupon code 5',
    'coupon code 6',
    'coupon code 7',
    'coupon code 8',
    'coupon code 9',
    'coupon code 10',
    'coupon code 11',
    'coupon code 12',
  ];

  //copy text
  const copyText = (text: string) => {
    navigator.clipboard.writeText(text);
    toast('Copied!');
  };

  return (
    <div className="bg-animeTwice w-full h-screen flex flex-col items-center">
      <div className=" rounded-[10px]  text-[56px] font-bold text-white p-4 mt-28 min-w-[400px] text-center flex justify-center items-center flex-col relative">
        <div
          className="bg-contain bg-no-repeat rounded-full min-w-fit w-full max-w-[240px] position-absolute top-0 left-0"
          style={{
            backgroundImage: `url(${'https://crypto.news/app/uploads/2022/04/1inch.jpg'})`,
          }}
        ></div>
        1Inch
        <div className="bg-primary rounded-full h-2 w-full max-w-[240px] "></div>
      </div>
      <div className="h-12"></div>
      <div className="bg-[#171717] opacity-80 rounded-[10px] w-3/4 flex flex-col p-4 overflow-y-scroll max-h-[800px]">
        {cupons.map((coupon) => (
          <div
            className="bg-[#323232] w-full rounded-[10px] p-4 font-bold flex mb-2"
            key={coupon}
          >
            <div className="text-white  mr-2">{coupon}</div>
            <div className=" bg-[#8E8E8E] w-full rounded-[10px] flex items-center p-2">
              <div className="text-white">asjd0j12Ajoi1...123kl;Sajkld</div>
              <div className="flex-1"></div>
              <Button
                text="Copy"
                size="small"
                onClick={(e) => copyText(coupon)}
              ></Button>
            </div>
          </div>
        ))}
      </div>
      <ToastContainer />
    </div>
  );
};

export default dashboard;
