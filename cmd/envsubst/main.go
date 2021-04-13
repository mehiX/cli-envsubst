package main

/**
Read an input stream, replace shell variable with respective enviornment variables, write the result on the output stream.
**/

import (
	"flag"
	"fmt"
	"os"

	"github.com/mehiX/cli-envsubst/transform"
	"github.com/mehiX/cli-envsubst/transformers"
)

// receive a string from standard input, process it, print it back to standard output
func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Println("Read an input stream, replace shell variable with respective enviornment variables, write the result on the output stream")
		flag.PrintDefaults()
	}

	flag.Parse()

	transform.Do(os.Stdin, os.Stdout, transformers.LookupEnvCaseIns)

}
