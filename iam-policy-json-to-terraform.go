package main

import "fmt"
import (
	"bufio"
	"flag"
	"github.com/flosell/iam-policy-json-to-terraform/converter"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	policyName := flag.String("name", "policy", "name of the policy in generated hcl")
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)
	b, err := ioutil.ReadAll(reader)

	if err != nil {
		log.Fatal("unable to read stdin: ", err)
	}

	converted, err := converter.Convert(*policyName, b)

	if err != nil {
		log.Fatal("unable to convert: ", err)
	}

	fmt.Print(converted)

}
