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
	CodeProposedCertificateAlreadyExists           = errors.Register(ModuleName, 401, "ProposedCertificateAlreadyExists")
	CodeProposedCertificateDoesNotExist            = errors.Register(ModuleName, 402, "ProposedCertificateDoesNotExist")
	CodeCertificateAlreadyExists                   = errors.Register(ModuleName, 403, "CertificateAlreadyExists")
	CodeCertificateDoesNotExist                    = errors.Register(ModuleName, 404, "CertificateDoesNotExist")
	CodeProposedCertificateRevocationAlreadyExists = errors.Register(ModuleName, 405, "ProposedCertificateRevocationAlreadyExists")
	CodeProposedCertificateRevocationDoesNotExist  = errors.Register(ModuleName, 406, "ProposedCertificateRevocationDoesNotExist ")
	CodeRevokedCertificateDoesNotExist             = errors.Register(ModuleName, 407, "RevokedCertificateDoesNotExist")
	CodeInappropriateCertificateType               = errors.Register(ModuleName, 408, "InappropriateCertificateType")
	CodeInvalidCertificate                         = errors.Register(ModuleName, 409, "InvalidCertificate")
)

func ErrProposedCertificateAlreadyExists(subject string, subjectKeyID string) error {
	return errors.Wrap(CodeProposedCertificateAlreadyExists,
		fmt.Sprintf("Proposed X509 root certificate associated with the combination "+
			"of subject=%v and subjectKeyID=%v already exists on the ledger", subject, subjectKeyID))
}

func ErrProposedCertificateDoesNotExist(subject string, subjectKeyID string) error {
	return errors.Wrap(CodeProposedCertificateDoesNotExist,
		fmt.Sprintf("No proposed X509 root certificate associated "+
			"with the combination of subject=%v and subjectKeyID=%v on the ledger. "+
			"The cerificate either does not exists or already approved.", subject, subjectKeyID))
}

func ErrCertificateAlreadyExists(issuer string, serialNumber string) error {
	return errors.Wrap(CodeCertificateAlreadyExists,
		fmt.Sprintf("X509 certificate associated with the combination of "+
			"issuer=%v and serialNumber=%v already exists on the ledger", issuer, serialNumber))
}

func ErrCertificateDoesNotExist(subject string, subjectKeyID string) error {
	return errors.Wrap(CodeCertificateDoesNotExist,
		fmt.Sprintf("No X509 certificate associated with the "+
			"combination of subject=%v and subjectKeyID=%v on the ledger", subject, subjectKeyID))
}

func ErrProposedCertificateRevocationAlreadyExists(subject string, subjectKeyID string) error {
	return errors.Wrap(CodeProposedCertificateRevocationAlreadyExists,
		fmt.Sprintf("Proposed X509 root certificate revocation associated with the combination "+
			"of subject=%v and subjectKeyID=%v already exists on the ledger", subject, subjectKeyID))
}

func ErrProposedCertificateRevocationDoesNotExist(subject string, subjectKeyID string) error {
	return errors.Wrap(CodeProposedCertificateRevocationDoesNotExist,
		fmt.Sprintf("No proposed X509 root certificate revocation associated "+
			"with the combination of subject=%v and subjectKeyID=%v on the ledger.", subject, subjectKeyID))
}

func ErrRevokedCertificateDoesNotExist(subject string, subjectKeyID string) error {
	return errors.Wrap(CodeRevokedCertificateDoesNotExist,
		fmt.Sprintf("No revoked X509 certificate associated with the "+
			"combination of subject=%v and subjectKeyID=%v on the ledger", subject, subjectKeyID))
}

func ErrInappropriateCertificateType(error interface{}) error {
	return errors.Wrap(CodeInappropriateCertificateType, fmt.Sprintf("%v", error))
}

func ErrCodeInvalidCertificate(error interface{}) error {
	return errors.Wrap(CodeInvalidCertificate, fmt.Sprintf("%v", error))
}
