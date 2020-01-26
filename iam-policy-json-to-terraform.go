package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/flosell/iam-policy-json-to-terraform/converter"
)

// AppVersion : current version
const AppVersion = "1.2.0"

func main() {
	policyName := flag.String("name", "policy", "name of the policy in generated hcl")
	version := flag.Bool("version", false, "prints the version")
	flag.Parse()

	versionValue := *version
	if versionValue {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

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
