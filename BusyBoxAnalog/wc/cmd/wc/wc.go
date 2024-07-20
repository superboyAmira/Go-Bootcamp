package main

import (
	"fmt"
	"goday02/src/wc/internal/handler"
	"goday02/src/wc/internal/input"
	"sync"
)

func main() {
	settings, err := input.GetSettings()
	if err != nil {
		return
	}

	var wg sync.WaitGroup
	infoAll := make(map[string]handler.ResultInfo)
	resultChan := make(chan handler.ResultInfo, len(settings.FileSequence))
	for _, path := range settings.FileSequence {
		wg.Add(1)
		go handler.HandleProccess(path, settings, &wg, resultChan)
		infoAll[path] = <-resultChan
	}
	wg.Wait()
	close(resultChan)

	for path_i, info := range infoAll {
		if settings.Lflag && settings.Mflag && settings.Wflag {
			fmt.Printf("%d %d %d\t%s\n", info.Lines, info.Chars, info.Words, path_i)
		} else if settings.Lflag && settings.Mflag {
			fmt.Printf("%d %d\t%s\n", info.Lines, info.Chars, path_i)
		} else if settings.Lflag && settings.Wflag {
			fmt.Printf("%d %d\t%s\n", info.Lines, info.Words, path_i)
		} else if settings.Mflag && settings.Wflag {
			fmt.Printf("%d %d\t%s\n", info.Chars, info.Words, path_i)
		} else if settings.Lflag {
			fmt.Printf("%d\t%s\n", info.Lines, path_i)
		} else if settings.Mflag {
			fmt.Printf("%d\t%s\n", info.Chars, path_i)
		} else if settings.Wflag {
			fmt.Printf("%d\t%s\n", info.Words, path_i)
		}
	}
}