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
	CodeAccountAlreadyExists                  = errors.Register(ModuleName, 101, "AccountAlreadyExists")
	CodeAccountDoesNotExist                   = errors.Register(ModuleName, 102, "AccountDoesNotExist")
	CodePendingAccountAlreadyExists           = errors.Register(ModuleName, 103, "PendingAccountAlreadyExists")
	CodePendingAccountDoesNotExist            = errors.Register(ModuleName, 104, "PendingAccountDoesNotExist")
	CodePendingAccountRevocationAlreadyExists = errors.Register(ModuleName, 105, "PendingAccountRevocationAlreadyExists")
	CodePendingAccountRevocationDoesNotExist  = errors.Register(ModuleName, 106, "PendingAccountRevocationDoesNotExist")
)

func ErrAccountAlreadyExists(address interface{}) error {
	return errors.Wrap(CodeAccountAlreadyExists,
		fmt.Sprintf("Account associated with the address=%v already exists on the ledger", address))
}

func ErrAccountDoesNotExist(address interface{}) error {
	return errors.Wrap(CodeAccountDoesNotExist,
		fmt.Sprintf("No account associated with the address=%v on the ledger", address))
}

func ErrPendingAccountAlreadyExists(address interface{}) error {
	return errors.Wrap(CodePendingAccountAlreadyExists,
		fmt.Sprintf("Pending account associated with the address=%v already exists on the ledger", address))
}

func ErrPendingAccountDoesNotExist(address interface{}) error {
	return errors.Wrap(CodePendingAccountDoesNotExist,
		fmt.Sprintf("No pending account associated with the address=%v on the ledger", address))
}

func ErrPendingAccountRevocationAlreadyExists(address interface{}) error {
	return errors.Wrap(CodePendingAccountRevocationAlreadyExists,
		fmt.Sprintf("Pending account revocation associated with the address=%v already exists on the ledger", address))
}

func ErrPendingAccountRevocationDoesNotExist(address interface{}) error {
	return errors.Wrap(CodePendingAccountRevocationDoesNotExist,
		fmt.Sprintf("No pending account revocation associated with the address=%v on the ledger", address))
}
