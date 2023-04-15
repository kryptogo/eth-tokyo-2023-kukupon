import { ethers } from "ethers";
import { useQuery } from "@tanstack/react-query";

import {
  getWhitelistingPaymaster,
  getSimpleAccount,
  getKeyFromCoupon,
  getGasFee,
  getHttpRpcClient,
} from "../utils";

const config = {
  rpcUrl: process.env.NEXT_PUBLIC_JSON_RPC_URL!,
  entrypoint: process.env.NEXT_PUBLIC_ENTRYPOINT_ADDRESS!,
  simpleAccountFactory: process.env.NEXT_PUBLIC_SIMPLE_ACCOUNT_FACTORY!,
  bundlerUrl: process.env.NEXT_PUBLIC_BUNDLER_URL!,
};

const provider = new ethers.providers.JsonRpcProvider(config.rpcUrl, {
  chainId: 80001,
  name: "mumbai",
});
const paymasterAPI = getWhitelistingPaymaster();

type TxInfo = {
  target: string;
  value?: string;
  data: string;
};

export const sendUserOperation = async (code: string, txInfo: TxInfo) => {
  const signingKey = getKeyFromCoupon(code);
  const accountAPI = getSimpleAccount(
    provider,
    signingKey,
    config.entrypoint,
    config.simpleAccountFactory,
    paymasterAPI
  );

  const op = await accountAPI.createSignedUserOp({
    ...txInfo,
    ...(await getGasFee(provider)),
  });

  const client = await getHttpRpcClient(
    provider,
    config.bundlerUrl,
    config.entrypoint
  );
  const uoHash = await client.sendUserOpToBundler(op);
  const txHash = await accountAPI.getUserOpReceipt(uoHash);

  return txHash;
};

const getAcountAPI = async (code: string) => {
  const signingKey = getKeyFromCoupon(code);
  const accountAPI = getSimpleAccount(
    provider,
    signingKey,
    config.entrypoint,
    config.simpleAccountFactory,
    paymasterAPI
  );

  return accountAPI;
};

export const useAccountAPI = (code: string) =>
  useQuery(["accountAPI", code], () => getAcountAPI(code));
