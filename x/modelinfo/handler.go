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

package modelinfo

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelinfo/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelinfo/internal/types"
)

func NewHandler(keeper keeper.Keeper, authKeeper auth.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.MsgAddModelInfo:
			return handleMsgAddModelInfo(ctx, keeper, authKeeper, msg)
		case types.MsgUpdateModelInfo:
			return handleMsgUpdateModelInfo(ctx, keeper, authKeeper, msg)
			/*		case type.MsgDeleteModelInfo:
					return handleMsgDeleteModelInfo(ctx, keeper, authKeeper, msg)*/
		default:
			errMsg := fmt.Sprintf("unrecognized nameservice Msg type: %v", msg.Type())

			return nil, errors.Wrap(errors.ErrInvalidRequest, errMsg)
		}
	}
}

func handleMsgAddModelInfo(ctx sdk.Context, keeper keeper.Keeper, authKeeper auth.Keeper,
	msg types.MsgAddModelInfo) (*sdk.Result, error) {
	// check if model already exists
	if keeper.IsModelInfoPresent(ctx, msg.VID, msg.PID) {
		return nil, types.ErrModelInfoAlreadyExists(msg.VID, msg.PID)
	}

	// check sender has enough rights to add model
	if err := checkAddModelRights(ctx, authKeeper, msg.Signer); err != nil {
		return nil, err
	}

	modelInfo := types.NewModelInfo(
		msg.VID,
		msg.PID,
		msg.CID,
		msg.Version,
		msg.Name,
		msg.Description,
		msg.SKU,
		msg.HardwareVersion,
		msg.FirmwareVersion,
		msg.OtaURL,
		msg.OtaChecksum,
		msg.OtaChecksumType,
		msg.Custom,
		msg.TisOrTrpTestingCompleted,
		msg.Signer,
	)

	// store new model
	keeper.SetModelInfo(ctx, modelInfo)

	return &sdk.Result{}, nil
}

func handleMsgUpdateModelInfo(ctx sdk.Context, keeper keeper.Keeper, authKeeper auth.Keeper,
	msg types.MsgUpdateModelInfo) (*sdk.Result, error) {
	// check if model exists
	if !keeper.IsModelInfoPresent(ctx, msg.VID, msg.PID) {
		return nil, types.ErrModelInfoDoesNotExist(msg.VID, msg.PID)
	}

	modelInfo := keeper.GetModelInfo(ctx, msg.VID, msg.PID)

	// check if sender has enough rights to update model
	if err := checkUpdateModelRights(modelInfo.Owner, msg.Signer); err != nil {
		return nil, err
	}

	if msg.OtaURL != "" && modelInfo.OtaURL == "" {
		return nil, types.ErrOtaURLCannotBeSet(msg.VID, msg.PID)
	}

	// updates existing model value only if corresponding value in MsgUpdate is not empty

	if msg.CID != 0 {
		modelInfo.CID = msg.CID
	}

	if msg.Description != "" {
		modelInfo.Description = msg.Description
	}

	if msg.OtaURL != "" {
		modelInfo.OtaURL = msg.OtaURL
	}

	if msg.Custom != "" {
		modelInfo.Custom = msg.Custom
	}

	modelInfo.TisOrTrpTestingCompleted = msg.TisOrTrpTestingCompleted

	// store updated model
	keeper.SetModelInfo(ctx, modelInfo)

	return &sdk.Result{}, nil
}

//nolint:unused,deadcode
func handleMsgDeleteModelInfo(ctx sdk.Context, keeper keeper.Keeper, authKeeper auth.Keeper,
	msg types.MsgDeleteModelInfo) (*sdk.Result, error) {
	// check if model exists
	if !keeper.IsModelInfoPresent(ctx, msg.VID, msg.PID) {
		return nil, types.ErrModelInfoDoesNotExist(msg.VID, msg.PID)
	}

	modelInfo := keeper.GetModelInfo(ctx, msg.VID, msg.PID)

	// check if sender has enough rights to delete model
	if err := checkUpdateModelRights(modelInfo.Owner, msg.Signer); err != nil {
		return nil, err
	}

	// remove model from the store
	keeper.DeleteModelInfo(ctx, msg.VID, msg.PID)

	return &sdk.Result{}, nil
}

func checkAddModelRights(ctx sdk.Context, authKeeper auth.Keeper, signer sdk.AccAddress) error {
	// sender must have Vendor role to add new model
	if !authKeeper.HasRole(ctx, signer, auth.Vendor) {
		return errors.Wrap(errors.ErrUnauthorized, fmt.Sprintf("MsgAddModelInfo transaction should be "+
			"signed by an account with the %s role", auth.Vendor))
	}

	return nil
}

func checkUpdateModelRights(owner sdk.AccAddress, signer sdk.AccAddress) error {
	// sender must be equal to owner to edit model
	if !signer.Equals(owner) {
		return errors.Wrap(errors.ErrUnauthorized, fmt.Sprintf("MsgUpdateModelInfo tx should be signed by owner"))
	}

	return nil
}
