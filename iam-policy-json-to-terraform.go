package main

import "fmt"
import (
	"github.com/hashicorp/hcl/hcl/printer"
	"os"
	"github.com/flosell/iam-policy-json-to-terraform/encoder"
)

func main() {

	policy_document := encoder.PolicyDocument{
		Name: "deny_access_without_mfa",
		Statements: encoder.StatementList{
			Statements: []*encoder.Statement{
				&encoder.Statement{
					Sid: "BlockMostAccessUnlessSignedInWithMFA",

				},
			},
		},
	}

	printer.Fprint(os.Stdout, policy_document.Encode())
	fmt.Printf("\nHello, world.\n")

}
