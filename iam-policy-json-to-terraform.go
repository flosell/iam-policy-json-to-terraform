package main

import "fmt"
import (
	"github.com/flosell/iam-policy-json-to-terraform/encoder"
	"log"
	"bufio"
	"os"
	"io/ioutil"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	b, err := ioutil.ReadAll(reader)

	if err != nil {
		log.Fatal("unable to read stdin: ", err)
	}

	converted, err := encoder.Convert(b)

	if err != nil {
		log.Fatal("unable to convert: ", err)
	}

	fmt.Print(converted)

}
