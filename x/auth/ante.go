// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth/internal/types"
)

// SignatureVerificationGasConsumer is the type of function that is used to both consume gas when verifying signatures
// and also to accept or reject different types of PubKey's. This is where apps can define their own PubKey.
type SignatureVerificationGasConsumer = func(meter sdk.GasMeter, pubkey crypto.PubKey) error

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers.
//nolint:funlen
func NewAnteHandler(ak Keeper, sigGasConsumer SignatureVerificationGasConsumer) sdk.AnteHandler {
	return func(
		ctx sdk.Context, tx sdk.Tx, simulate bool,
	) (newCtx sdk.Context, err error) {
		// all transactions must be of type auth.StdTx.
		stdTx, ok := tx.(auth.StdTx)
		if !ok {
			// Set a gas meter with limit 0 as to prevent an infinite gas meter attack
			// during runTx.
			newCtx = SetGasMeter(simulate, ctx, 0)

			return newCtx, errors.Wrap(errors.ErrInvalidRequest, "tx must be StdTx")
		}

		newCtx = SetGasMeter(simulate, ctx, stdTx.Fee.Gas)

		// AnteHandlers must have their own defer/recover in order for the BaseApp
		// to know how much gas was used! This is because the GasMeter is created in
		// the AnteHandler, but if it panics the context won't be set properly in
		// runTx's recover call.
		defer func() {
			if r := recover(); r != nil {
				switch rType := r.(type) {
				case sdk.ErrorOutOfGas:
					log := fmt.Sprintf(
						"out of gas in location: %v; gasWanted: %d, gasUsed: %d",
						rType.Descriptor, stdTx.Fee.Gas, newCtx.GasMeter().GasConsumed(),
					)
					err = errors.Wrap(errors.ErrOutOfGas, log)
				default:
					panic(r)
				}
			}
		}()

		if err := tx.ValidateBasic(); err != nil {
			return newCtx, err
		}

		newCtx.GasMeter().ConsumeGas(types.TxSizeCostPerByte*sdk.Gas(len(newCtx.TxBytes())), "txSize")

		if err = ValidateMemo(stdTx); err != nil {
			return newCtx, err
		}

		// signatures contains the sequence number, account number, and signatures.
		signers := stdTx.GetSigners()
		signatures := stdTx.Signatures

		isGenesis := ctx.BlockHeight() == 0

		for i, signature := range signatures {
			account, error := GetSignerAcc(newCtx, ak, signers[i])
			if error != nil {
				return newCtx, error
			}

			// check signature, return account with incremented nonce.
			signBytes := GetSignBytes(newCtx.ChainID(), stdTx, account, isGenesis)
			account, err = processSig(newCtx, account, signature, signBytes, simulate, sigGasConsumer)

			if err != nil {
				return newCtx, err
			}

			ak.SetAccount(newCtx, account)
		}

		return newCtx, nil // continue...
	}
}

// GetSignerAcc returns an account for a given address that is expected to sign a transaction.
func GetSignerAcc(ctx sdk.Context, keeper Keeper, address sdk.AccAddress) (acc types.Account, err error) {
	if !keeper.IsAccountPresent(ctx, address) {
		return acc, types.ErrAccountDoesNotExist(address)
	}

	acc = keeper.GetAccount(ctx, address)

	return acc, nil
}

// ValidateMemo validates the memo size.
func ValidateMemo(stdTx auth.StdTx) error {
	memoLength := len(stdTx.GetMemo())
	if uint64(memoLength) > types.MaxMemoCharacters {
		return errors.Wrap(errors.ErrMemoTooLarge,
			fmt.Sprintf(
				"maximum number of characters is %d but received %d characters",
				types.MaxMemoCharacters, memoLength,
			),
		)
	}
	return nil
}

// verify the signature and increment the sequence.
func processSig(
	ctx sdk.Context, acc types.Account, sig auth.StdSignature, signBytes []byte, simulate bool,
	sigGasConsumer SignatureVerificationGasConsumer,
) (updatedAcc Account, err error) {
	if err := sigGasConsumer(ctx.GasMeter(), acc.PubKey); err != nil {
		return acc, err
	}

	if !simulate && !acc.PubKey.VerifyBytes(signBytes, sig.Signature) {
		return acc, errors.Wrap(errors.ErrUnauthorized, "Signature verification failed; verify correct account sequence and chain-id")
	}

	acc.Sequence++

	return acc, nil
}

// DefaultSigVerificationGasConsumer is the default implementation of SignatureVerificationGasConsumer. It consumes gas
// for signature verification based upon the public key type. The cost is fetched from the given params and is matched
// by the concrete type.
func DefaultSigVerificationGasConsumer(meter sdk.GasMeter, pubkey crypto.PubKey) error {
	switch pubkey := pubkey.(type) {
	case ed25519.PubKeyEd25519:
		meter.ConsumeGas(types.DefaultSigVerifyCostED25519, "ante verify: ed25519")

		return errors.Wrap(errors.ErrInvalidPubKey, "ED25519 public keys are unsupported")

	case secp256k1.PubKeySecp256k1:
		meter.ConsumeGas(types.DefaultSigVerifyCostSecp256k1, "ante verify: secp256k1")

		return nil

	default:
		return errors.Wrap(errors.ErrInvalidPubKey, fmt.Sprintf("unrecognized public key type: %T", pubkey))
	}
}

// SetGasMeter returns a new context with a gas meter set from a given context.
func SetGasMeter(simulate bool, ctx sdk.Context, gasLimit uint64) sdk.Context {
	// In various cases such as simulation and during the genesis block, we do not
	// meter any gas utilization.
	if simulate || ctx.BlockHeight() == 0 {
		return ctx.WithGasMeter(sdk.NewInfiniteGasMeter())
	}

	return ctx.WithGasMeter(sdk.NewGasMeter(gasLimit))
}

// GetSignBytes returns a slice of bytes to sign over for a given transaction
// and an account.
func GetSignBytes(chainID string, stdTx auth.StdTx, acc types.Account, genesis bool) []byte {
	var accNum uint64
	if !genesis {
		accNum = acc.AccountNumber
	}

	return auth.StdSignBytes(
		chainID, accNum, acc.Sequence, stdTx.Fee, stdTx.Msgs, stdTx.Memo,
	)
}
