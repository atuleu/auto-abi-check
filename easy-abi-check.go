package main

import (
	"os"
)

var major_version = "0"
var minor_version = "1"
var patch_version = "0~dev"

var version_string = major_version + "." + minor_version + "." + patch_version

func main() {
	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}
}
