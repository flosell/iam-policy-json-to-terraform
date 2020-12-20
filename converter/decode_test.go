package converter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshalJsonStringAndArrayToArray(t *testing.T) {
	var testcases = []struct {
		in        string
		out       jsonStatements
		errString string
	}{
		// behaving as an array
		{`[]`, jsonStatements{}, ""},
		{`[{"Sid":"helloworld"}]`, jsonStatements{jsonStatement{Sid: "helloworld"}}, ""},

		// behaving as an array - errors
		{`[{"Sid":"helloworld"}`, nil, "unexpected end of JSON input"},
		{`["helloworld"]`, nil, "json: cannot unmarshal string into Go value of type converter.jsonStatement"},

		// behaving as a single statement
		{`{"Sid":"helloworld"}`, jsonStatements{jsonStatement{Sid: "helloworld"}}, ""},
		// behaving as a single statement - errors
		{`{"Sid":42}`, nil, "json: cannot unmarshal number into Go struct field jsonStatement.Sid of type string"},
	}

	for _, testcase := range testcases {
		t.Run(testcase.in, func(t *testing.T) {
			var out jsonStatements
			err := json.Unmarshal([]byte(testcase.in), &out)

			if testcase.errString != "" {
				assert.EqualError(t, err, testcase.errString)
			} else if err != nil {
				t.Fatal(err)
			} else {
				assert.Equal(t, testcase.out, out)
			}
		})
	}

}
func TestEscapeHclSnippetsInThem(t *testing.T) {
	var testcases = []struct {
		in  string
		out string
	}{
		// no escaping
		{`"arn:aws:s3:::foo"`, `"arn:aws:s3:::foo"`},
		// multiple hcl snippets
		{`"foo${var.bar}", "foo${var.baz}"`, `"foo${var.bar}", "foo${var.baz}"`},
		// single escape
		{`"arn:aws:s3:::foo/${join("/",local.path_elements)}"`, `"arn:aws:s3:::foo/${join(\"/\",local.path_elements)}"`},
		// ${ unclosed parens should be ignored
		{`hello${world`, `hello${world`},
		// $ at end of string
		{`US$`, `US$`},
		// ${ at end of string
		{`hello${`, `hello${`},
		// nested expressions
		{`"foo${join("/${var.sep}",local.bar)}"`, `"foo${join(\"/${var.sep}\",local.bar)}"`},
		{`"foo${join("/${chomp("foo")}",local.bar)}"`, `"foo${join(\"/${chomp(\"foo\")}\",local.bar)}"`},
		// unbalanced curly braces
		{"}", "}"},
		{"${}}", "${}}"},
	}

	for _, testcase := range testcases {
		t.Run(testcase.in, func(t *testing.T) {
			out := escapeHclSnippetsInJSON([]byte(testcase.in))

			assert.Equal(t, testcase.out, string(out))
		})
	}

}
