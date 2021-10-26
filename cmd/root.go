package cmd

import (
	"flag"
	"fmt"
	"os"

	models "github.com/hahwul/licaner/pkg/models"
	"github.com/logrusorgru/aurora"
)

func Init() models.Options {
	pipeOption := flag.Bool("pipe", false, "Use pipeline")
	versionOption := flag.Bool("version", false, "Version of licaner")
	_ = pipeOption

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, aurora.White("Usage: %s [flags]\n").String(), os.Args[0])
		fmt.Fprintf(os.Stderr, aurora.White("Flags:\n").String())
		flag.PrintDefaults()
	}

	flag.Parse()

	options := models.Options{
		Pipe: *pipeOption,
	}

	// Show version
	if *versionOption {
		fmt.Println("version")
		return options
	}

	return options
}
