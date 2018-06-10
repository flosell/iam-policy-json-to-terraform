package converter

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestConvertFromJsonToTerraformHcl(t *testing.T) {
	input, ferr := ioutil.ReadFile("fixtures/simple.json")
	if ferr != nil {
		t.Fatal(ferr)
	}
	expectedOutput, ferr := ioutil.ReadFile("fixtures/simple.tf")
	if ferr != nil {
		t.Fatal(ferr)
	}

	actualOutput, err := Convert("policy", input)

	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t,
		string(expectedOutput),
		actualOutput,
	)
}
