package input

import (
	"flag"
)

type CLIcfg struct {
	FileFlag bool
	DirFlag  bool
	SymFlag  bool
	ExtFlag  bool
	Path     string
	Ext      string
}

func ParseFile() (fileInfo CLIcfg) {
	f := flag.Bool("f", false, "to find only files")
	d := flag.Bool("d", false, "to find only directories")
	sl := flag.Bool("sl", false, "to find only sym links")
	ext := flag.String("ext", "", "set extension, only with -f")

	flag.CommandLine.Init("flag settings didn`t stop executing a app", flag.ContinueOnError)

	flag.Parse()

	if flag.NArg() != 0 {
		fileInfo.Path = flag.Arg(0)
	}

	if !*f && !*d && !*sl && *ext == "" {
		fileInfo.FileFlag = true
		fileInfo.DirFlag = true
		fileInfo.SymFlag = true
	} else {
		fileInfo.FileFlag = *f
		fileInfo.DirFlag = *d
		fileInfo.SymFlag = *sl
	}
	if *ext != "" && fileInfo.FileFlag {
		fileInfo.ExtFlag = true
		fileInfo.Ext = "." + *ext 
	} else {
		fileInfo.Ext = ""
		fileInfo.ExtFlag = false
	}

	return
}
