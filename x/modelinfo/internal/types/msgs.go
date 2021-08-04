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

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const RouterKey = ModuleName

//nolint:maligned
type MsgAddModelInfo struct {
	VID                      uint16         `json:"vid"`
	PID                      uint16         `json:"pid"`
	CID                      uint16         `json:"cid,omitempty"`
	Version                  string         `json:"version,omitempty"`
	Name                     string         `json:"name"`
	Description              string         `json:"description"`
	SKU                      string         `json:"sku"`
	HardwareVersion          string         `json:"hardware_version"`
	FirmwareVersion          string         `json:"firmware_version"`
	OtaURL                   string         `json:"ota_url,omitempty"`
	OtaChecksum              string         `json:"ota_checksum,omitempty"`
	OtaChecksumType          string         `json:"ota_checksum_type,omitempty"`
	Custom                   string         `json:"custom,omitempty"`
	TisOrTrpTestingCompleted bool           `json:"tis_or_trp_testing_completed"`
	Signer                   sdk.AccAddress `json:"signer"`
}

func NewMsgAddModelInfo(
	vid uint16,
	pid uint16,
	cid uint16,
	version string,
	name string,
	description string,
	sku string,
	hardwareVersion string,
	firmwareVersion string,
	otaURL string,
	otaChecksum string,
	otaChecksumType string,
	custom string,
	tisOrTrpTestingCompleted bool,
	signer sdk.AccAddress,
) MsgAddModelInfo {
	return MsgAddModelInfo{
		VID:                      vid,
		PID:                      pid,
		CID:                      cid,
		Version:                  version,
		Name:                     name,
		Description:              description,
		SKU:                      sku,
		HardwareVersion:          hardwareVersion,
		FirmwareVersion:          firmwareVersion,
		OtaURL:                   otaURL,
		OtaChecksum:              otaChecksum,
		OtaChecksumType:          otaChecksumType,
		Custom:                   custom,
		TisOrTrpTestingCompleted: tisOrTrpTestingCompleted,
		Signer:                   signer,
	}
}

func (m MsgAddModelInfo) Route() string {
	return RouterKey
}

func (m MsgAddModelInfo) Type() string {
	return "add_model_info"
}

func (m MsgAddModelInfo) ValidateBasic() error {
	if m.Signer.Empty() {
		return errors.Wrap(errors.ErrInvalidAddress, "Invalid Signer: it cannot be empty")
	}

	if m.VID == 0 {
		return errors.Wrap(errors.ErrInvalidRequest, "Invalid VID: it must be non-zero 16-bit unsigned integer")
	}

	if m.PID == 0 {
		return errors.Wrap(errors.ErrInvalidRequest, "Invalid PID: it must be non-zero 16-bit unsigned integer")
	}

	if len(m.Name) == 0 {
		return errors.Wrap(errors.ErrInvalidRequest, "Invalid Name: it cannot be empty")
	}

	if len(m.Description) == 0 {
		return errors.Wrap(errors.ErrInvalidRequest, "Invalid Description: it cannot be empty")
	}

	if len(m.SKU) == 0 {
		return errors.Wrap(errors.ErrInvalidRequest, "Invalid SKU: it cannot be empty")
	}

	if len(m.HardwareVersion) == 0 {
		return errors.Wrap(errors.ErrInvalidRequest, "Invalid HardwareVersion: it cannot be empty")
	}

	if len(m.FirmwareVersion) == 0 {
		return errors.Wrap(errors.ErrInvalidRequest, "Invalid FirmwareVersion: it cannot be empty")
	}

	if m.OtaURL != "" || m.OtaChecksum != "" || m.OtaChecksumType != "" {
		if m.OtaURL == "" || m.OtaChecksum == "" || m.OtaChecksumType == "" {
			return errors.Wrap(errors.ErrInvalidRequest, "Invalid MsgAddModelInfo: the fields OtaURL, OtaChecksum and "+
				"OtaChecksumType must be either specified together, or not specified together")
		}
	}

	return nil
}

func (m MsgAddModelInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgAddModelInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

//nolint:maligned
type MsgUpdateModelInfo struct {
	VID                      uint16         `json:"vid"`
	PID                      uint16         `json:"pid"`
	CID                      uint16         `json:"cid,omitempty"`
	Description              string         `json:"description,omitempty"`
	OtaURL                   string         `json:"ota_url,omitempty"`
	Custom                   string         `json:"custom,omitempty"`
	TisOrTrpTestingCompleted bool           `json:"tis_or_trp_testing_completed"`
	Signer                   sdk.AccAddress `json:"signer"`
}

func NewMsgUpdateModelInfo(
	vid uint16,
	pid uint16,
	cid uint16,
	description string,
	otaURL string,
	custom string,
	tisOrTrpTestingCompleted bool,
	signer sdk.AccAddress,
) MsgUpdateModelInfo {
	return MsgUpdateModelInfo{
		VID:                      vid,
		PID:                      pid,
		CID:                      cid,
		Description:              description,
		OtaURL:                   otaURL,
		Custom:                   custom,
		TisOrTrpTestingCompleted: tisOrTrpTestingCompleted,
		Signer:                   signer,
	}
}

func (m MsgUpdateModelInfo) Route() string {
	return RouterKey
}

func (m MsgUpdateModelInfo) Type() string {
	return "update_model_info"
}

func (m MsgUpdateModelInfo) ValidateBasic() error {
	if m.Signer.Empty() {
		return errors.Wrap(errors.ErrInvalidAddress, "Invalid Signer: it cannot be empty")
	}

	if m.VID == 0 {
		return errors.Wrap(errors.ErrInvalidRequest, "Invalid VID: it must be non-zero 16-bit unsigned integer")
	}

	if m.PID == 0 {
		return errors.Wrap(errors.ErrInvalidRequest, "Invalid PID: it must be non-zero 16-bit unsigned integer")
	}

	return nil
}

func (m MsgUpdateModelInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgUpdateModelInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

type MsgDeleteModelInfo struct {
	VID    uint16         `json:"vid"`
	PID    uint16         `json:"pid"`
	Signer sdk.AccAddress `json:"signer"`
}

func NewMsgDeleteModelInfo(vid uint16, pid uint16, signer sdk.AccAddress) MsgDeleteModelInfo {
	return MsgDeleteModelInfo{
		VID:    vid,
		PID:    pid,
		Signer: signer,
	}
}

func (m MsgDeleteModelInfo) Route() string {
	return RouterKey
}

func (m MsgDeleteModelInfo) Type() string {
	return "delete_model_info"
}

func (m MsgDeleteModelInfo) ValidateBasic() error {
	if m.Signer.Empty() {
		return errors.Wrap(errors.ErrInvalidAddress, "Invalid Signer: it cannot be empty")
	}

	if m.VID == 0 {
		return errors.Wrap(errors.ErrInvalidRequest, "Invalid VID: it must be non-zero 16-bit unsigned integer")
	}

	if m.PID == 0 {
		return errors.Wrap(errors.ErrInvalidRequest, "Invalid PID: it must be non-zero 16-bit unsigned integer")
	}

	return nil
}

func (m MsgDeleteModelInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgDeleteModelInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
