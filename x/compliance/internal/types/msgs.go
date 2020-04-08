package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

type MsgAddModelInfo struct {
	ID                       string         `json:"id"`
	Name                     string         `json:"name"`
	Description              string         `json:"description"`
	SKU                      string         `json:"sku"`
	FirmwareVersion          string         `json:"firmware_version"`
	HardwareVersion          string         `json:"hardware_version"`
	CertificateID            string         `json:"certificate_id,omitempty"`
	CertifiedDate            time.Time      `json:"certified_date,omitempty"`
	TisOrTrpTestingCompleted bool           `json:"tis_or_trp_testing_completed"`
	Signer                   sdk.AccAddress `json:"signer"`
}

func NewMsgAddModelInfo(id string, name string, description string, sku string,
	firmwareVersion string, hardwareVersion string, certificateID string, certifiedDate time.Time,
	tisOrTrpTestingCompleted bool, signer sdk.AccAddress) MsgAddModelInfo {
	return MsgAddModelInfo{
		ID:                       id,
		Name:                     name,
		Description:              description,
		SKU:                      sku,
		FirmwareVersion:          firmwareVersion,
		HardwareVersion:          hardwareVersion,
		CertificateID:            certificateID,
		CertifiedDate:            certifiedDate,
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

func (m MsgAddModelInfo) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress(m.Signer.String())
	}

	if len(m.ID) == 0 ||
		len(m.Name) == 0 ||
		len(m.Description) == 0 ||
		len(m.SKU) == 0 ||
		len(m.FirmwareVersion) == 0 ||
		len(m.HardwareVersion) == 0 {
		return sdk.ErrUnknownRequest("Id, Name, Description, SKU, FirmwareVersion and HardwareVersion " +
			"cannot be empty")
	}

	return nil
}

func (m MsgAddModelInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgAddModelInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

type MsgUpdateModelInfo struct {
	ID                          string         `json:"id"`
	NewName                     string         `json:"new_name"`
	NewOwner                    sdk.AccAddress `json:"new_owner"`
	NewDescription              string         `json:"new_description"`
	NewSKU                      string         `json:"new_sku"`
	NewFirmwareVersion          string         `json:"new_firmware_version"`
	NewHardwareVersion          string         `json:"new_hardware_version"`
	NewCertificateID            string         `json:"new_certificate_id,omitempty"`
	NewCertifiedDate            time.Time      `json:"new_certified_date,omitempty"`
	NewTisOrTrpTestingCompleted bool           `json:"new_tis_or_trp_testing_completed"`
	Signer                      sdk.AccAddress `json:"signer"`
}

func NewMsgUpdateModelInfo(id string, newName string, newOwner sdk.AccAddress, newDescription string, newSKU string,
	newFirmwareVersion string, newHardwareVersion string, newCertificateID string, newCertifiedDate time.Time,
	newTisOrTrpTestingCompleted bool, signer sdk.AccAddress) MsgUpdateModelInfo {
	return MsgUpdateModelInfo{
		ID:                          id,
		NewName:                     newName,
		NewOwner:                    newOwner,
		NewDescription:              newDescription,
		NewSKU:                      newSKU,
		NewFirmwareVersion:          newFirmwareVersion,
		NewHardwareVersion:          newHardwareVersion,
		NewCertificateID:            newCertificateID,
		NewCertifiedDate:            newCertifiedDate,
		NewTisOrTrpTestingCompleted: newTisOrTrpTestingCompleted,
		Signer:                      signer,
	}
}

func (m MsgUpdateModelInfo) Route() string {
	return RouterKey
}

func (m MsgUpdateModelInfo) Type() string {
	return "update_model_info"
}

func (m MsgUpdateModelInfo) ValidateBasic() sdk.Error {
	if m.NewOwner.Empty() {
		return sdk.ErrInvalidAddress(m.NewOwner.String())
	}

	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress(m.Signer.String())
	}

	if len(m.ID) == 0 ||
		len(m.NewName) == 0 ||
		len(m.NewDescription) == 0 ||
		len(m.NewSKU) == 0 ||
		len(m.NewFirmwareVersion) == 0 ||
		len(m.NewHardwareVersion) == 0 {
		return sdk.ErrUnknownRequest("Id, NewName, NewDescription, NewSKU, NewFirmwareVersion and NewHardwareVersion " +
			"cannot be empty")
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
	ID     string         `json:"id"`
	Signer sdk.AccAddress `json:"signer"`
}

func NewMsgDeleteModelInfo(id string, signer sdk.AccAddress) MsgDeleteModelInfo {
	return MsgDeleteModelInfo{ID: id, Signer: signer}
}

func (m MsgDeleteModelInfo) Route() string {
	return RouterKey
}

func (m MsgDeleteModelInfo) Type() string {
	return "delete_model_info"
}

func (m MsgDeleteModelInfo) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress(m.Signer.String())
	}

	if len(m.ID) == 0 {
		return sdk.ErrUnknownRequest("Id cannot be empty")
	}

	return nil
}

func (m MsgDeleteModelInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgDeleteModelInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
