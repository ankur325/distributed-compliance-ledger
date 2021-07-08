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
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Codespace sdk.CodespaceType = ModuleName

	CodeComplianceInfoDoesNotExist sdk.CodeType = 301
	CodeInconsistentDates          sdk.CodeType = 302
	CodeAlreadyCertifyed           sdk.CodeType = 303
	CodeModelInfoDoesNotExist      sdk.CodeType = 304
)

func ErrComplianceInfoDoesNotExist(vid interface{}, pid interface{},
	softwareVersion interface{}, hardwareVersion interface{}, certificationType interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeComplianceInfoDoesNotExist,
		fmt.Sprintf("No certification information about the model with vid=%v, "+
			"pid=%v softwareVersion=%v, hardwareVersion=%v and "+
			"certification_type=%v on the ledger. This means that the model is either not certified yet or "+
			"certified by default (off-ledger).", vid, pid, softwareVersion, hardwareVersion, certificationType))
}

func ErrInconsistentDates(error interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeInconsistentDates, fmt.Sprintf("%v", error))
}

func ErrAlreadyCertifyed(vid interface{}, pid interface{},
	softwareVersion interface{}, hardwareVersion interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeAlreadyCertifyed,
		fmt.Sprintf("Model with vid=%v, pid=%v, softwareVersion=%v, hardwareVersion=%v already certified on the ledger ",
			vid, pid, softwareVersion, hardwareVersion))
}

func ErrModelInfoDoesNotExist(vid interface{}, pid interface{},
	softwareVersion interface{}, hardwareVersion interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeModelInfoDoesNotExist,
		fmt.Sprintf("Model with vid=%v, pid=%v, softwareVersion=%v, hardwareVersion=%v does not exist on the ledger",
			vid, pid, softwareVersion, hardwareVersion))
}
