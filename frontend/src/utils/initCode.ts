import { ethers } from "ethers";

export function getInitCode(
  simpleAccountFactory: string,
  ownerAddress: string
) {
  // Define the function signature and parameter types
  const createAccountFunction = {
    name: "createAccount",
    type: "function",
    stateMutability: "nonpayable",
    inputs: [
      { name: "owner", type: "address" },
      { name: "index", type: "uint256" },
    ],
  };

  // Define the function parameter values
  const index = 0;

  // Create an Interface instance
  const contractInterface = new ethers.utils.Interface([createAccountFunction]);

  // ABI-encode the function and parameters
  const encodedFunctionData = contractInterface.encodeFunctionData(
    createAccountFunction.name,
    [ownerAddress, index]
  );

  // Concatenate the two strings
  const concatenatedStrings = ethers.utils.hexConcat([
    simpleAccountFactory,
    encodedFunctionData,
  ]);
  return concatenatedStrings;
}
