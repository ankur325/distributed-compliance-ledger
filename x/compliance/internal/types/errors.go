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

	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	CodeComplianceInfoDoesNotExist = errors.Register(ModuleName, 301, "No certification information exists")
	CodeInconsistentDates          = errors.Register(ModuleName, 302, "Inconsistent Dates")
	CodeAlreadyCertifyed           = errors.Register(ModuleName, 303, "Already Certified")
	CodeModelInfoDoesNotExist      = errors.Register(ModuleName, 304, "No Model Info Present")
)

func ErrComplianceInfoDoesNotExist(vid interface{}, pid interface{}, certificationType interface{}) error {
	return errors.Wrap(CodeComplianceInfoDoesNotExist,
		fmt.Sprintf("No certification information about the model with vid=%v, pid=%v and "+
			"certification_type=%v on the ledger. This means that the model is either not certified yet or "+
			"certified by default (off-ledger).", vid, pid, certificationType))
}

func ErrInconsistentDates(error interface{}) error {
	return errors.Wrap(CodeInconsistentDates, fmt.Sprintf("%v", error))
}

func ErrAlreadyCertifyed(vid interface{}, pid interface{}) error {
	return errors.Wrap(CodeAlreadyCertifyed,
		fmt.Sprintf("Model with vid=%v, pid=%v already certified on the ledger", vid, pid))
}

func ErrModelInfoDoesNotExist(vid interface{}, pid interface{}) error {
	return errors.Wrap(CodeModelInfoDoesNotExist,
		fmt.Sprintf("Model with vid=%v, pid=%v does not exist on the ledger", vid, pid))
}
