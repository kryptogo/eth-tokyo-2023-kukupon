package service

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// VerifySignature checks the signature of the given message.
func VerifySignature(from, sigHex, msg string) bool {
	sig := hexutil.MustDecode(sigHex)

	msgHash := accounts.TextHash([]byte(msg))
	sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

	recovered, err := crypto.SigToPub(msgHash, sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)
	return strings.EqualFold(from, recoveredAddr.Hex())
}
