package input

import (
	"errors"
	"flag"
)

type Settings struct {
	Lflag        bool
	Mflag        bool
	Wflag        bool
	FileSequence []string
}


func GetSettings() (sett *Settings, err error) {
	l := flag.Bool("l", false, "to find cnt string")
	m := flag.Bool("m", false, "to find cnt chars")
	w := flag.Bool("w", false, "to find cnt words")

	flag.CommandLine.Init("flag settings didn`t stop executing a app", flag.ContinueOnError)

	flag.Parse()

	if flag.NArg() == 0 {
		err = errors.New("emty args")
		return nil, err
	}
	sett = &Settings{
		Lflag:        *l,
		Mflag:        *m,
		Wflag:        *w,
		FileSequence: flag.Args(),
	}
	
	return
}