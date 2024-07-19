package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func  HandleStr(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(path)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	cnt := 0
	for scanner.Scan() {
		cnt++
	}
	if scanner.Err() != nil {
		return
	}
	fmt.Printf("%d\t%s", cnt, path)
}

func HandleChar(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(path)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	chrcnt := 0
	for scanner.Scan() {
		chrcnt += len([]rune(scanner.Text()))
		
	}
	if scanner.Err() != nil {
		return
	}
	fmt.Printf("%d\t%s", chrcnt, path)
}

func HandleWord(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(path)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	strcnt := 0
	for scanner.Scan() {
		strcnt += len(strings.Fields(scanner.Text()))
	}
	if scanner.Err() != nil {
		return
	}
	fmt.Printf("%d\t%s", strcnt, path)
}