package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/flosell/iam-policy-json-to-terraform/converter"
	"github.com/mattn/go-isatty"
	"io/ioutil"
	"log"
	"os"
)

// AppVersion : current version
const AppVersion = "1.7.0"

func main() {
	policyName := flag.String("name", "policy", "name of the policy in generated hcl")
	version := flag.Bool("version", false, "prints the version")
	flag.Parse()

	versionValue := *version
	if versionValue {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	if isatty.IsTerminal(os.Stdin.Fd()) {
		os.Stderr.WriteString("Paste a valid IAM policy and press the EOF afterwards.\n")
		os.Stderr.WriteString("Alternatively, you can pipe input directly into the command.\n")
		os.Stderr.WriteString("> ")
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
