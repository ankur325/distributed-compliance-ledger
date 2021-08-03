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
	CodeValidatorAlreadyExist = errors.Register(ModuleName, 601, "Validator already exists")
	CodeValidatorDoesNotExist = errors.Register(ModuleName, 602, "Validator does not exists")
	CodePoolIsFull            = errors.Register(ModuleName, 603, "Poll is full")
	CodeAccountAlreadyHasNode = errors.Register(ModuleName, 604, "Account already has node")
)

func ErrValidatorExists(address interface{}) error {
	return errors.Wrap(CodeValidatorAlreadyExist,
		fmt.Sprintf("Validator associated with the validator_address=%v already exists on the ledger", address))
}

func ErrValidatorDoesNotExist(address interface{}) error {
	return errors.Wrap(CodeValidatorDoesNotExist,
		fmt.Sprintf("No validator associated with the validator_address=%v on the ledger", address))
}

func ErrPoolIsFull() error {
	return errors.Wrap(CodePoolIsFull,
		fmt.Sprintf("Pool ledger already contains maximum number of active nodes: \"%v\"", MaxNodes))
}

func ErrAccountAlreadyHasNode(address interface{}) error {
	return errors.Wrap(CodeAccountAlreadyHasNode,
		fmt.Sprintf("There is already node stored on the ledger managed by an account"+
			" associated with the address=\"%v\"", address))
}
