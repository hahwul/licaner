package cmd

import (
	"flag"
	"fmt"
	"os"

	models "github.com/hahwul/licaner/pkg/models"
	"github.com/logrusorgru/aurora"
)

func Init() (models.Options, bool) {
	versionOption := flag.Bool("version", false, "Version of licaner")

	urlOption := flag.String("url", "", "Scan from URL")
	pipeOption := flag.Bool("pipe", false, "Scan from pipeline")
	urlsfileOption := flag.String("urls-file", "", "Scan from URLs file")
	rawfileOption := flag.String("raw-file", "", "Scan from Raw HTTP Request File")
	formatOption := flag.String("format", "line", "Output format")
	outputOption := flag.String("output", "", "Output filename")
	nocbOption := flag.Bool("no-cache-busting", false, "Not use cache busting")
	_ = pipeOption

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, aurora.White("Usage: %s [flags]\n").String(), os.Args[0])
		fmt.Fprintf(os.Stderr, aurora.White("Flags:\n").String())
		flag.PrintDefaults()
	}

	flag.Parse()

	options := models.Options{
		Pipeline:       *pipeOption,
		URL:            *urlOption,
		URLsFile:       *urlsfileOption,
		RawFile:        *rawfileOption,
		Format:         *formatOption,
		Output:         *outputOption,
		NoCacheBusting: *nocbOption,
	}

	// Show version
	if *versionOption {
		fmt.Println("version")
		return options, false
	}

	return options, true
}
