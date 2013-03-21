package main

import (
	"github.com/jessevdk/go-flags"
	"os"
	"fmt"
)


type Options struct {
	Version func() error `short:"V" long:"version" description:"Print the version"`

}

var options = &Options{
	Version: func() error {
		fmt.Println("easy-abi-check version ", version_string)
		os.Exit(0)
		return nil
	},
}

var parser = flags.NewParser(options, flags.Default)

func init() {
	parser.Description = "A wrapper around abi-compliance-check to automate C/C++ compliance check"
}
