package converter

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestConvertFromJsonToTerraformHcl(t *testing.T) {
	var fixtures = []struct {
		in  string
		out string
	}{
		{"fixtures/simple.json", "fixtures/simple.tf"},
		{"fixtures/empty-policy.json", "fixtures/empty-policy.tf"},
		{"fixtures/single-statement.json", "fixtures/single-statement.tf"},
		{"fixtures/tf-interpolations.json", "fixtures/tf-interpolations.tf"},
	}

	for _, fixture := range fixtures {
		t.Run(fixture.in, func(t *testing.T) {
			input := readFixtureFile(t, fixture.in)
			expectedOutput := readFixtureFile(t, fixture.out)

			actualOutput, err := Convert("policy", input)

			if err != nil {
				t.Fatal(err)
			}

			assert.EqualValues(t,
				string(expectedOutput),
				actualOutput,
			)
		})
	}

}

func TestConversionErrors(t *testing.T) {
	var testcases = []struct {
		fixtureFile         string
		partialErrorMessage string
		expectedError       error
	}{
		{"fixtures/error-broken.json", "unexpected end of JSON input", nil},
		{"fixtures/error-no-statements.json", "", ErrorLackOfStatements},
		{"fixtures/error-cloudformation-snippet.json", "", ErrorLooksLikeCloudformation},
	}

	for _, testcase := range testcases {
		t.Run(testcase.fixtureFile, func(t *testing.T) {
			input := readFixtureFile(t, testcase.fixtureFile)

			_, err := Convert("policy", input)

			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), testcase.partialErrorMessage)

			if testcase.expectedError != nil {
				assert.ErrorIs(t, err, testcase.expectedError)
			}
		})
	}
}

func TestEscapeDollarSigns(t *testing.T) {
	var testcases = []struct {
		in  string
		out string
	}{
		// basic strings
		{"", ""},
		{"hello world", "hello world"},
		// aws policy variables
		{"${aws:username}", "$${aws:username}"},
		{"home/${aws:username}/*", "home/$${aws:username}/*"},
		{"arn:aws:s3:::bucket/${aws:PrincipalTag/department}", "arn:aws:s3:::bucket/$${aws:PrincipalTag/department}"},
		{"${s3:x-amz-acl}", "$${s3:x-amz-acl}"},

		// aws special character escaping
		{"${*}", "$${*}"},
		{"${?}", "$${?}"},
		{"${$}", "$${$}"},

		// terraform interpolations
		{"${aws_kms_key.key.arn}", "${aws_kms_key.key.arn}"},
		{"hello ${var.greeting}!", "hello ${var.greeting}!"},
		{"arn:aws:s3:::foo/${join(var.separator,local.path_elements)}", "arn:aws:s3:::foo/${join(var.separator,local.path_elements)}"},
	}

	for _, testcase := range testcases {
		t.Run(testcase.in, func(t *testing.T) {
			assert.EqualValues(t,
				testcase.out,
				escapePolicyVariables(testcase.in),
			)
		})

	}
}

func readFixtureFile(t *testing.T, file string) []byte {
	input, ferr := ioutil.ReadFile(file)
	if ferr != nil {
		t.Fatal(ferr)
	}
	return input
}
