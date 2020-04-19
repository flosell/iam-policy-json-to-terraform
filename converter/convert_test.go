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
	}

	for _, fixture := range fixtures {
		t.Run(fixture.in, func(t *testing.T) {
			input, ferr := ioutil.ReadFile(fixture.in)
			if ferr != nil {
				t.Fatal(ferr)
			}
			expectedOutput, ferr := ioutil.ReadFile(fixture.out)
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
		})
	}

}
