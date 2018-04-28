package encoder

import (
	"io/ioutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	expected := []Statement{
		{
			Sid: "BlockMostAccessUnlessSignedInWithMFA",
		},
	}
	jsonString, ferr := ioutil.ReadFile("fixtures/simple.json")

	if ferr != nil {
		t.Fatal(ferr)
	}

	actual, err := Decode(jsonString)

	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t,
		expected,
		actual,
	)
}
