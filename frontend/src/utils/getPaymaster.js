const { PaymasterAPI } = require("@account-abstraction/sdk");
const { UserOperationStruct } = require("@account-abstraction/contracts");

class WhitelistingPaymasterAPI extends PaymasterAPI {
  constructor() {
    super();
  }

  async getPaymasterAndData(userOp) {
    console.log("userOp: ", userOp);
    return process.env.NEXT_PUBLIC_PAYMASTER_ADDRESS;
  }
}

module.exports.getWhitelistingPaymaster = () => new WhitelistingPaymasterAPI();
