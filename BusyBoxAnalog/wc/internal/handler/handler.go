package handler

import (
	"bufio"
	"goday02/src/wc/internal/input"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

type ResultInfo struct {
	Lines int
	Chars int
	Words int
}

func HandleProccess(path string, settings *input.Settings, wg *sync.WaitGroup, result chan<- ResultInfo) {
	defer wg.Done()
	var res ResultInfo
	if settings.Lflag {
		res.Lines = handleStr(path)
	}
	if settings.Mflag {
		res.Chars = handleChar(path)
	}
	if settings.Wflag {
		res.Words = handleWord(path)
	}
	result <- res
}

func handleStr(path string) int {
	file, err := os.Open(path)
	if err != nil {
		return -1
	}
	scanner := bufio.NewScanner(file)
	cnt := 0
	for scanner.Scan() {
		cnt++
	}
	if scanner.Err() != nil {
		return -1
	}
	return cnt
}

func handleChar(path string) int {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return -1
	}
	charCount := utf8.RuneCount(bytes)
	return charCount
}

func handleWord(path string) int {
	file, err := os.Open(path)
	if err != nil {
		return -1
	}
	scanner := bufio.NewScanner(file)
	strcnt := 0
	for scanner.Scan() {
		strcnt += len(strings.Fields(scanner.Text()))
	}
	if scanner.Err() != nil {
		return -1
	}
	return strcnt
}
