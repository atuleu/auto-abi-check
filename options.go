package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)

type VersionCmd struct {
}

func (c *VersionCmd) Execute(args []string) error {
	fmt.Println("auto-abi-check version ", version_string)
	return nil
}

var opts = LibraryDescriptionOptions{}

var parser = flags.NewParser(&opts, flags.Default)

func init() {
	parser.Description = "A wrapper around abi-compliance-check to automate C/C++ compliance check"
	parser.AddCommand("version",
		"prints the current version",
		"prints the current version, simply.",
		&VersionCmd{})
}
