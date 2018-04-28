package converter

import (
	"io/ioutil"
	"github.com/stretchr/testify/assert"
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

	actualOutput, err := Convert(input)

	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t,
		expectedOutput,
		actualOutput,
	)
}

