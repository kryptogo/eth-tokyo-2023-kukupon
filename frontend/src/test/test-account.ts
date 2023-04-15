import { ethers } from "ethers";
import {
  getWhitelistingPaymaster,
  getSimpleAccount,
  getKeyFromCoupon,
  getGasFee,
  printOp,
  getHttpRpcClient,
} from "../utils";

require("dotenv").config();

const withPM = true;
const couponCode = "123";

const config = {
  rpcUrl: process.env.NEXT_PUBLIC_JSON_RPC_URL!,
  entrypoint: process.env.NEXT_PUBLIC_ENTRYPOINT_ADDRESS!,
  simpleAccountFactory: process.env.NEXT_PUBLIC_SIMPLE_ACCOUNT_FACTORY!,
  bundlerUrl: process.env.NEXT_PUBLIC_BUNDLER_URL!,
};

async function main() {
  const provider = new ethers.providers.JsonRpcProvider(config.rpcUrl, {
    chainId: 80001,
    name: "mumbai",
  });
  const paymasterAPI = withPM ? getWhitelistingPaymaster() : undefined;
  const signingKey = getKeyFromCoupon(couponCode);
  console.log("signingKey:", signingKey);

  const accountAPI = getSimpleAccount(
    provider,
    signingKey,
    config.entrypoint,
    config.simpleAccountFactory,
    paymasterAPI
  );
  const initCode = await accountAPI.getAccountInitCode();
  console.log("account init code:", initCode);
  console.log("simple account address:", await accountAPI.getAccountAddress());

  const op = await accountAPI.createSignedUserOp({
    target: ethers.utils.getAddress(
      "0xd2a8c5226e47f3c128657b15ba630ac98b0d5f4a"
    ),
    value: 0,
    data: "0x40d097c30000000000000000000000003fe5ac82b997255b8226abb6aefd91f405fe2e8e",
    ...(await getGasFee(provider)),
  });
  // const to = ethers.utils.getAddress(
  //   "0x0901549Bc297BCFf4221d0ECfc0f718932205e33"
  // );
  // const amount = ethers.utils.parseEther("0.01");
  // console.log(`Transferring 0.01 MATIC...`);
  // const op = await accountAPI.createSignedUserOp({
  //   target: to,
  //   value: amount,
  //   data: "0x",
  //   ...(await getGasFee(provider)),
  // });
  console.log(`Signed UserOperation: ${await printOp(op)}`);

  const client = await getHttpRpcClient(
    provider,
    config.bundlerUrl,
    config.entrypoint
  );
  const uoHash = await client.sendUserOpToBundler(op);
  console.log(`UserOpHash: ${uoHash}`);

  console.log("Waiting for transaction...");
  const txHash = await accountAPI.getUserOpReceipt(uoHash);
  console.log(`Transaction hash: ${txHash}`);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
