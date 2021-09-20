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

package vendorinfo

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/internal/types"
)

type GenesisState struct {
	VendorInfoRecords []VendorInfo `json:"vendor_info_records"`
}

func NewGenesisState() GenesisState {
	return GenesisState{VendorInfoRecords: []VendorInfo{}}
}

//nolint:cognit
func ValidateGenesis(data GenesisState) error {
	//TODO ADD validation

	// for _, record := range data.VendorInfoRecords {
	// 	if record.Model.VID == 0 {
	// 		return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: Invalid VID. Value: %v", record))
	// 	}

	// 	if record.Model.PID == 0 {
	// 		return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: Invalid PID. Value: %v", record))
	// 	}

	// 	if record.Model.ProductName == "" {
	// 		return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: Missed ProductName. Value: %v", record))
	// 	}

	// 	if record.Model.Description == "" {
	// 		return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: Missed Description. Value: %v", record))
	// 	}

	// 	if record.Model.SKU == "" {
	// 		return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: Missed SKU. Value: %v", record))
	// 	}

	// 	if record.Model.HardwareVersion == 0 {
	// 		return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: Missed HardwareVersion. Value: %v", record))
	// 	}

	// 	if record.Model.HardwareVersionString == "" {
	// 		return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: Missed HardwareVersionString. Value: %v", record))
	// 	}

	// 	if record.Model.SoftwareVersion == 0 {
	// 		return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: Missed SoftwareVersion. Value: %v", record))
	// 	}

	// 	if record.Model.SoftwareVersionString == "" {
	// 		return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: Missed SoftwareVersionString. Value: %v", record))
	// 	}

	// 	if record.Owner.Empty() {
	// 		return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: Missed Owner. Value: %v", record))
	// 	}

	// 	if record.Model.OtaURL != "" || record.Model.OtaChecksum != "" || record.Model.OtaChecksumType != "" {
	// 		if record.Model.OtaURL == "" || record.Model.OtaChecksum == "" || record.Model.OtaChecksumType == "" {
	// 			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: The fields OtaURL, OtaChecksum and "+
	// 				"OtaChecksumType must be either specified together, or not specified together. Value: %v", record))
	// 		}
	// 	}
	// }

	return nil
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.VendorInfoRecords {
		keeper.SetVendorInfo(ctx, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []VendorInfo

	k.IterateVendorInfos(ctx, func(modelInfo types.VendorInfo) (stop bool) {
		records = append(records, modelInfo)

		return false
	})

	return GenesisState{VendorInfoRecords: records}
}
