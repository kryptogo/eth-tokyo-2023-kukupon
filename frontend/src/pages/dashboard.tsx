import Button from "@/components/Button";
import useCouponStore from "@/store/useCouponStore";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
const Dashboard = () => {
  const coupons = useCouponStore((state) => state.coupons);

  //copy text
  const copyText = (text: string) => {
    navigator.clipboard.writeText(text);
    toast("Copied!");
  };

  return (
    <div className="flex flex-col items-center w-full h-screen bg-animeTwice">
      <div className=" rounded-[10px]  text-[56px] font-bold text-white p-4 mt-28 min-w-[400px] text-center flex justify-center items-center flex-col relative">
        <div
          className="bg-contain bg-no-repeat rounded-full min-w-fit w-full max-w-[240px] position-absolute top-0 left-0"
          style={{
            backgroundImage: `url(${"https://crypto.news/app/uploads/2022/04/1inch.jpg"})`,
          }}
        ></div>
        KryptoGO
        <div className="bg-primary rounded-full h-2 w-full max-w-[240px] "></div>
      </div>
      <div className="h-12"></div>
      <div className="bg-[#171717] opacity-80 rounded-[10px] w-3/4 flex flex-col p-4 overflow-y-scroll max-h-[800px]">
        {Array.from(Object.keys(coupons)).map((title) => (
          <>
            {coupons?.[title]?.map((code) => (
              <div
                className="bg-[#323232] w-full rounded-[10px] p-4 font-bold flex mb-2"
                key={code}
              >
                <div className="mr-2 text-white">{`${title} COUPON`}</div>
                <div className=" bg-[#8E8E8E] w-full rounded-[10px] flex items-center p-2">
                  <div className="text-white">{code}</div>
                  <div className="flex-1"></div>
                  <Button
                    text="Copy"
                    size="small"
                    onClick={(e) => copyText(code)}
                  ></Button>
                </div>
              </div>
            ))}
          </>
        ))}
      </div>
      <ToastContainer />
    </div>
  );
};

export default Dashboard;
