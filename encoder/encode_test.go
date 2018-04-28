package encoder

import (
	"testing"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)

func TestEncoder(t *testing.T) {
	data_source := DataSource{
		Type: "aws_iam_policy_document",
		Name: "deny_access_without_mfa",
		Statements: []Statement{
			Statement{
				Sid: "BlockMostAccessUnlessSignedInWithMFA",
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
