package main

import "fmt"
import (
	"github.com/flosell/iam-policy-json-to-terraform/encoder"
	"log"
)

func main() {

	policy_document := encoder.DataSource{
		Type: "aws_iam_policy_document",
		Name: "deny_access_without_mfa",
		Statements: []encoder.Statement{
			encoder.Statement{
				Sid: "BlockMostAccessUnlessSignedInWithMFA",
			},
		},
	}

	encoded, err := encoder.Encode(policy_document)

	if err != nil {
		log.Fatal("unable to encode: ", err)
	}

	fmt.Print(encoded)

}
