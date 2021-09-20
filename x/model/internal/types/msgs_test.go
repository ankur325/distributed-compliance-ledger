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

//nolint:testpackage
package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
)

func getTestModel() Model {
	return Model{
		VID:                                      testconstants.VID,
		PID:                                      testconstants.PID,
		DeviceTypeID:                             testconstants.DeviceTypeID,
		ProductName:                              testconstants.ProductName,
		ProductLabel:                             testconstants.ProductLabel,
		PartNumber:                               testconstants.PartNumber,
		CommissioningCustomFlow:                  testconstants.CommissioningCustomFlow,
		CommissioningCustomFlowURL:               testconstants.CommissioningCustomFlowURL,
		CommissioningModeInitialStepsHint:        testconstants.CommissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: testconstants.CommissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      testconstants.CommissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: testconstants.CommissioningModeSecondaryStepsInstruction,
		UserManualURL: testconstants.UserManualURL,
		SupportURL:    testconstants.SupportURL,
		ProductURL:    testconstants.ProductURL,
	}
}

func getTestModelForUpdate() Model {
	return Model{
		VID:                        testconstants.VID,
		PID:                        testconstants.PID,
		DeviceTypeID:               testconstants.DeviceTypeID,
		ProductLabel:               testconstants.ProductLabel,
		CommissioningCustomFlowURL: testconstants.CommissioningCustomFlowURL,
		UserManualURL:              testconstants.UserManualURL,
		SupportURL:                 testconstants.SupportURL,
		ProductURL:                 testconstants.ProductURL,
	}
}

func TestNewMsgAddModel(t *testing.T) {
	model := getTestModel()
	msg := NewMsgAddModel(model, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "add_model_info")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

//nolint:funlen
func TestMsgAddModelValidation(t *testing.T) {
	require.Nil(t, NewMsgAddModel(getTestModel(), testconstants.Signer).ValidateBasic())

	// Invalid VID
	model := getTestModel()
	model.VID = 0
	require.NotNil(t, NewMsgAddModel(model, testconstants.Signer).ValidateBasic())

	// Invalid PID
	model = getTestModel()
	model.PID = 0
	require.NotNil(t, NewMsgAddModel(model, testconstants.Signer).ValidateBasic())

	// DeviceTypeID is Mandator
	model = getTestModel()
	model.DeviceTypeID = 0
	require.Nil(t, NewMsgAddModel(model, testconstants.Signer).ValidateBasic())

	// Name is mandatory
	model = getTestModel()
	model.ProductName = ""
	require.NotNil(t, NewMsgAddModel(model, testconstants.Signer).ValidateBasic())

	// ProductLabel is mandatory
	model = getTestModel()
	model.ProductLabel = ""
	require.NotNil(t, NewMsgAddModel(model, testconstants.Signer).ValidateBasic())

	// PartNumber is mandatory
	model = getTestModel()
	model.PartNumber = ""
	require.NotNil(t, NewMsgAddModel(model, testconstants.Signer).ValidateBasic())

	// Missing non mandatory
	model = getTestModel()
	model.CommissioningModeInitialStepsHint = 0
	model.CommissioningCustomFlow = 0
	model.CommissioningModeInitialStepsHint = 0
	model.CommissioningModeInitialStepsInstruction = ""
	require.Nil(t, NewMsgAddModel(model, testconstants.Signer).ValidateBasic())

	// Signer is nil
	model = getTestModel()
	require.NotNil(t, NewMsgAddModel(model, nil).ValidateBasic())
	// Signer is nil
	require.NotNil(t, NewMsgAddModel(model, []byte{}).ValidateBasic())
}

func TestMsgAddModelGetSignBytes(t *testing.T) {
	msg := NewMsgAddModel(getTestModel(), testconstants.Signer)
	expected := `{"type":"model/AddModel","value":{"Model":{"CDVersionNumber":312,"chipBlob":"Chip Blob Text","cid":12345,"commissioningCustomFlow":1,"commissioningCustomFlowURL":"https://sampleflowurl.dclmodel","commissioningModeInitialStepsHint":2,"commissioningModeInitialStepsInstruction":"commissioningModeInitialStepsInstruction details","commissioningModeSecondaryStepsHint":3,"commissioningModeSecondaryStepsInstruction":"commissioningModeSecondaryStepsInstruction steps","description":"Device Description","firmwareDigests":"Firmware Digest String","hardwareVersion":21,"hardwareVersionString":"2.1","otaBlob":"OTABlob Text","otaChecksum":"0xe3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855","otaChecksumType":"SHA-256","otaURL":"http://ota.firmware.com","pid":22,"productName":"Device Name","productURL":"https://url.producturl.dclmodel","releaseNotesURL":"https://url.releasenotes.dclmodel","revoked":false,"sku":"RCU2205A","softwareVersion":1,"softwareVersionString":"1.0","supportURL":"https://url.supporturl.dclmodel","userManualURL":"https://url.usermanual.dclmodel","vendorBlob":"Vendor Blob Text","vid":1},"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}` //nolint:lll
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestNewMsgUpdateModel(t *testing.T) {
	msg := NewMsgUpdateModel(getTestModelForUpdate(), testconstants.Signer)
	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "update_model_info")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestMsgUpdateModelValidation(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgUpdateModel
	}{

		{true, NewMsgUpdateModel(
			getTestModelForUpdate(), testconstants.Signer)},
		{false, NewMsgUpdateModel(
			getTestModelForUpdate(), nil)},
		{false, NewMsgUpdateModel(
			getTestModelForUpdate(), []byte{})},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgUpdateModelGetSignBytes(t *testing.T) {
	msg := NewMsgUpdateModel(getTestModelForUpdate(), testconstants.Signer)

	expected := `{"type":"model/UpdateModel","value":{"Model":{"CDVersionNumber":312,"chipBlob":"Chip Blob Text","cid":12345,"commissioningCustomFlowURL":"https://sampleflowurl.dclmodel","description":"Device Description","firmwareDigests":"Firmware Digest String","hardwareVersion":0,"hardwareVersionString":"","otaBlob":"OTABlob Text","otaChecksum":"0xe3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855","otaChecksumType":"SHA-256","otaURL":"http://ota.firmware.com","pid":22,"productName":"","productURL":"https://url.producturl.dclmodel","releaseNotesURL":"https://url.releasenotes.dclmodel","revoked":false,"sku":"","softwareVersion":0,"softwareVersionString":"","supportURL":"https://url.supporturl.dclmodel","userManualURL":"https://url.usermanual.dclmodel","vendorBlob":"Vendor Blob Text","vid":1},"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}` //nolint:lll

	require.Equal(t, expected, string(msg.GetSignBytes()))
}