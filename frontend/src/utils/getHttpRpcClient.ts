import { HttpRpcClient } from "@account-abstraction/sdk/dist/src/HttpRpcClient";
import { ethers } from "ethers";

export async function getHttpRpcClient(
  provider: ethers.providers.JsonRpcProvider,
  bundlerUrl: string,
  entryPointAddress: string
) {
  const chainId = await provider.getNetwork().then((net) => net.chainId);
  return new HttpRpcClient(bundlerUrl, entryPointAddress, chainId);
}
