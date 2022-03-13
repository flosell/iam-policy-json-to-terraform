package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWebWrapperConvertsString(t *testing.T) {
	result := ConvertString("someName", `
{
	"Statement": []
}`)
	assert.Equal(t, "data \"aws_iam_policy_document\" \"someName\" {}\n", result)
}

func TestWebWrapperPanicsIfThingsCanBeParsedSoWeGetANiceJavaScriptException(t *testing.T) {
	assert.Panicsf(t, func() { ConvertString("someName", "{") }, "Expecting panic")
}
