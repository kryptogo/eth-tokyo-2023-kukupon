import React, { useEffect, useState } from "react";
import { SimpleAccountAPI } from "@account-abstraction/sdk";
import { useAccountAPI } from "@/service/useOperation";

const useRedeemCode = (code: string) => {
  const { data, isLoading, isSuccess } = useAccountAPI(code);
  const [address, setAddress] = useState("");

  useEffect(() => {
    if (!data) return;

    const getAccountAddress = async (accountAPI: SimpleAccountAPI) => {
      const address = await accountAPI.getAccountAddress();
      setAddress(address);
    };

    getAccountAddress(data);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isSuccess]);

  return address;
};

export default useRedeemCode;
