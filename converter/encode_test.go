package converter

import (
	"testing"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)

func TestEncodeDataSourceStruct(t *testing.T) {
	data_source := HclDataSource{
		Type: "aws_iam_policy_document",
		Name: "deny_access_without_mfa",
		Statements: []HclStatement{
			HclStatement{
				Sid: "BlockMostAccessUnlessSignedInWithMFA",
				Effect: "Deny",
			},
		},
	}
	actual , err := Encode(data_source)
	if err != nil {
		t.Fatal(err)
	}
	expected, ferr := ioutil.ReadFile("fixtures/simple.tf")

	if ferr != nil {
		t.Fatal(ferr)
	}

	assert.EqualValues(t,
		string(expected),
		string(actual),
	)
}
