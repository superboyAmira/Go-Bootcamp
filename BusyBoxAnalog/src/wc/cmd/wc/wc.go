package main

import (
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

	for _, path := range settings.FileSequence {
		if settings.Lflag {
			wg.Add(1)
			go handler.HandleStr(path, &wg)
		}
		if settings.Mflag {
			wg.Add(1)
			go handler.HandleChar(path, &wg)
		}
		if settings.Wflag {
			wg.Add(1)
			go handler.HandleWord(path, &wg)
		}
	}
	wg.Wait()	
}