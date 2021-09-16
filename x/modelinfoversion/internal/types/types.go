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
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//nolint:maligned
type ModelVersion struct {
	VID                          uint16 `json:"vid"`
	PID                          uint16 `json:"pid"`
	SoftwareVersion              uint32 `json:"softwareVersion"`
	SoftwareVersionString        string `json:"softwareVersionString"`
	CDVersionNumber              uint16 `json:"CDVersionNumber"`
	FirmwareDigests              string `json:"firmwareDigests,omitempty"`
	SoftwareVersionValid         bool   `json:"softwareVersionValid"`
	OtaURL                       string `json:"otaURL,omitempty"`
	OtaFileSize                  uint64 `json:"otaFileSize"`
	OtaChecksum                  string `json:"otaChecksum,omitempty"`
	OtaChecksumType              uint16 `json:"otaChecksumType"`
	MinApplicableSoftwareVersion uint32 `json:"minApplicableSoftwareVersion"`
	MaxApplicableSoftwareVersion uint32 `json:"maxApplicableSoftwareVersion"`
	ReleaseNotesURL              string `json:"releaseNotesURL,omitempty"`
}

type ModelVersions struct {
	VID              uint16   `json:"vid"`
	PID              uint16   `json:"pid"`
	SoftwareVersions []uint32 `json:"softwareVersions"`
}

/*
	Model Version Info stored in KVStore
*/
//nolint:maligned
type ModelVersionInfo struct {
	ModelVersion
	Owner sdk.AccAddress `json:"owner"`
}

func (d ModelVersionInfo) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
