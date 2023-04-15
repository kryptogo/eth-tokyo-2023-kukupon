const ethers = require('ethers');
const simpleAccountFactory = '0x71D63edCdA95C61D6235552b5Bc74E32d8e2527B';

function getInitCode(ownerAddress) {
  // Define the function signature and parameter types
  const createAccountFunction = {
    name: 'createAccount',
    type: 'function',
    stateMutability: 'nonpayable',
    inputs: [
      { name: 'owner', type: 'address' },
      { name: 'index', type: 'uint256' },
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

console.log(getInitCode('0xd1AdE280FEbfBB0a929B99F95ab94FADcfDCF456'));
