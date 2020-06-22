package schema

import (
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
)

type FinalTranscriptCredential struct {
	*verifiable.Credential
}

type FinalTranscriptSubject struct {
	ID   string  `json:"id,omitempty"`
	Type string  `json:"type"`
	GPA  float64 `json:"gpa"`
}
