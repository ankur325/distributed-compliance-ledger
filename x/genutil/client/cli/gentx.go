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

package cli

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	tmos "github.com/tendermint/tendermint/libs/os"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/genutil"
	validator "github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/client/cli"
)

// GenTxCmd builds the application's gentx command.
//nolint:gocognit,funlen
func GenTxCmd(ctx *server.Context, cdc *codec.Codec, mbm module.BasicManager,
	defaultNodeHome, defaultCLIHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gentx",
		Short: "Generate a genesis transaction to create a validator",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			config := ctx.Config
			config.SetRoot(viper.GetString(flags.FlagHome))

			nodeID, valPubKey, err := genutil.InitializeNodeValidatorFiles(ctx.Config)
			if err != nil {
				return err
			}

			// Read --nodeID, if empty take it from priv_validator.json
			if nodeIDString := viper.GetString(validator.FlagNodeID); nodeIDString != "" {
				nodeID = nodeIDString
			}
			// Read --pubkey, if empty take it from priv_validator.json
			if valPubKeyString := viper.GetString(validator.FlagPubKey); valPubKeyString != "" {
				valPubKey, err = sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, valPubKeyString)
				if err != nil {
					return err
				}
			}

			genDoc, err := tmtypes.GenesisDocFromFile(config.GenesisFile())
			if err != nil {
				return err
			}

			var genesisState map[string]json.RawMessage
			if err = cdc.UnmarshalJSON(genDoc.AppState, &genesisState); err != nil {
				return err
			}

			if err = mbm.ValidateGenesis(genesisState); err != nil {
				return err
			}

			inBuf := bufio.NewReader(cmd.InOrStdin())
			kb, err := keys.NewKeyring(sdk.KeyringServiceName(),
				"test", viper.GetString(flagClientHome), inBuf)
			if err != nil {
				return errors.Wrap(err, "failed to initialize keybase")
			}

			name := viper.GetString(flags.FlagName)
			key, err := kb.Get(name)
			if err != nil {
				return errors.Wrap(err, "failed to read from keybase")
			}

			// Set flags for creating gentx
			viper.Set(flags.FlagHome, viper.GetString(flagClientHome))
			validator.PrepareFlagsForTxCreateValidator(config, nodeID, genDoc.ChainID, valPubKey)

			err = genutil.ValidateAccountInGenesis(genesisState, key.GetAddress(), cdc)
			if err != nil {
				return err
			}

			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			// create a 'create-validator' message
			txBldr, msg, err := validator.BuildCreateValidatorMsg(cliCtx, txBldr, true)
			if err != nil {
				return errors.Wrap(err, "failed to build create-validator message")
			}

			if key.GetType() == keys.TypeOffline || key.GetType() == keys.TypeMulti {
				fmt.Println("Offline key passed in. Use `tx sign` command to sign:")
				return utils.PrintUnsignedStdTx(txBldr, cliCtx, []sdk.Msg{msg})
			}

			// write the unsigned transaction to the buffer
			w := bytes.NewBuffer([]byte{})
			cliCtx = cliCtx.WithOutput(w)

			if err = utils.PrintUnsignedStdTx(txBldr, cliCtx, []sdk.Msg{msg}); err != nil {
				return errors.Wrap(err, "failed to print unsigned std tx")
			}

			// read the transaction
			stdTx, err := readUnsignedGenTxFile(cdc, w)
			if err != nil {
				return errors.Wrap(err, "failed to read unsigned gen tx file")
			}

			// sign the transaction and write it to the output file
			signedTx, err := utils.SignStdTx(txBldr, cliCtx, name, stdTx, false, true)
			if err != nil {
				return errors.Wrap(err, "failed to sign std tx")
			}

			// Fetch output file from
			outputDocument := viper.GetString(flags.FlagOutputDocument)
			if outputDocument == "" {
				outputDocument, err = makeOutputFilepath(config.RootDir, nodeID)
				if err != nil {
					return errors.Wrap(err, "failed to create output file path")
				}
			}

			if err := writeSignedGenTx(cdc, outputDocument, signedTx); err != nil {
				return errors.Wrap(err, "failed to write signed gen tx")
			}

			fmt.Fprintf(os.Stderr, "Genesis transaction written to %q\n", outputDocument)

			return nil
		},
	}

	ipDefault, _ := server.ExternalIP()
	cmd.Flags().String(validator.FlagIP, ipDefault, "The node's public IP")
	cmd.Flags().Bool(flags.FlagTrustNode, true, "Trust the Node")
	cmd.Flags().String(validator.FlagNodeID, "", "The node's NodeID")
	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "node's home directory")
	cmd.Flags().String(flagClientHome, defaultCLIHome, "client's home directory")
	cmd.Flags().String(flags.FlagFrom, "", "Name or address of private key with which to sign the gentx")
	cmd.Flags().String(flags.FlagOutputDocument, "",
		"write the genesis transaction JSON document to the given file instead of the default location")

	cmd.Flags().AddFlagSet(validator.InitValidatorFlags())

	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	_ = cmd.MarkFlagRequired(flags.FlagTrustNode)

	return cmd
}

func makeOutputFilepath(rootDir, nodeID string) (string, error) {
	writePath := filepath.Join(rootDir, "config", "gentx")
	if err := tmos.EnsureDir(writePath, 0700); err != nil {
		return "", err
	}
	return filepath.Join(writePath, fmt.Sprintf("gentx-%v.json", nodeID)), nil
}

func readUnsignedGenTxFile(cdc *codec.Codec, r io.Reader) (auth.StdTx, error) {
	var stdTx auth.StdTx

	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return stdTx, err
	}

	err = cdc.UnmarshalJSON(bytes, &stdTx)

	return stdTx, err
}

func writeSignedGenTx(cdc *codec.Codec, outputDocument string, tx auth.StdTx) error {
	outputFile, err := os.OpenFile(outputDocument, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	json, err := cdc.MarshalJSON(tx)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(outputFile, "%s\n", json)

	return err
}
