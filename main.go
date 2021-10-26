package main

import (
	cmd "github.com/hahwul/licaner/cmd"
	scan "github.com/hahwul/licaner/pkg/scan"
)

func main() {
	options := cmd.Init()
	scan.Scan(options)
}
