package input

import (
	"flag"
)

func GetFile() (path []string, archive *string, err error) {
	a := flag.String("a", "", "path to archive")

	flag.CommandLine.Init("flag settings didn`t stop executing a app", flag.ContinueOnError)
	flag.Parse()

	if *a != "" {
		archive = a
	} else {
		archive = nil
	}
	path = flag.Args()
	return
}
