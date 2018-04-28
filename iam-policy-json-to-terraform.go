package main

import "fmt"
import (
	"log"
	"bufio"
	"os"
	"io/ioutil"
	"github.com/flosell/iam-policy-json-to-terraform/converter"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	b, err := ioutil.ReadAll(reader)

	if err != nil {
		log.Fatal("unable to read stdin: ", err)
	}

	converted, err := converter.Convert(b)

	if err != nil {
		log.Fatal("unable to convert: ", err)
	}

	fmt.Print(converted)

}
