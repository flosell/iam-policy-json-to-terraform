package main

import (
	"github.com/flosell/iam-policy-json-to-terraform/converter"
	"github.com/gopherjs/gopherjs/js"
)

func ConvertString(policyName string, jsonString string) string {
	result, err := converter.Convert(policyName, []byte(jsonString))
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	js.Global.Set("convert", ConvertString)
}
