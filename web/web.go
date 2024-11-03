package main

import (
	"fmt"
	"github.com/flosell/iam-policy-json-to-terraform/converter"
	"syscall/js"
)

func ConvertString(this js.Value, args []js.Value) interface{} {
	policyName := args[0].String()
	jsonString := args[1].String()
	fmt.Printf("policyName %s jsonString %s", policyName, jsonString)
	result, err := converter.Convert(policyName, []byte(jsonString))
	if err != nil {
		panic(err)
	}
	return js.ValueOf(result)
}

func main() {
	wait := make(chan struct{}, 0)
	js.Global().Set("ConvertString", js.FuncOf(ConvertString))
	<-wait
}
