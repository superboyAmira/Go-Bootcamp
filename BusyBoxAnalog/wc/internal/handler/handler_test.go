package handler

import (
	"goday02/src/wc/internal/input"
	"sync"
	"testing"
)

func TestHandleProcessRus(t *testing.T) {
	path := "../../test/rus.txt"
	expected := ResultInfo{
		Lines: 5,
		Chars: 1169,
		Words: 166,
	}

	var wg sync.WaitGroup
	settings := &input.Settings{Lflag: true, Mflag: true, Wflag: true}
	resultChan := make(chan ResultInfo, 1)

	wg.Add(1)
	go HandleProccess(path, settings, &wg, resultChan)

	wg.Wait()
	close(resultChan)

	actual := <-resultChan
	if actual.Lines != expected.Lines {
		t.Errorf("HandleProccess(%s) Lines = %d; want %d", path, actual.Lines, expected.Lines)
	}
	if actual.Chars != expected.Chars {
		t.Errorf("HandleProccess(%s) Chars = %d; want %d", path, actual.Chars, expected.Chars)
	}
	if actual.Words != expected.Words {
		t.Errorf("HandleProccess(%s) Words = %d; want %d", path, actual.Words, expected.Words)
	}
}

func TestHandleProcessEng(t *testing.T) {
	path := "../../test/eng.txt"
	expected := ResultInfo{
		Lines: 5,
		Chars: 1088,
		Words: 185,
	}

	var wg sync.WaitGroup
	settings := &input.Settings{Lflag: true, Mflag: true, Wflag: true}
	resultChan := make(chan ResultInfo, 1)

	wg.Add(1)
	go HandleProccess(path, settings, &wg, resultChan)

	wg.Wait()
	close(resultChan)

	actual := <-resultChan
	if actual.Lines != expected.Lines {
		t.Errorf("HandleProccess(%s) Lines = %d; want %d", path, actual.Lines, expected.Lines)
	}
	if actual.Chars != expected.Chars {
		t.Errorf("HandleProccess(%s) Chars = %d; want %d", path, actual.Chars, expected.Chars)
	}
	if actual.Words != expected.Words {
		t.Errorf("HandleProccess(%s) Words = %d; want %d", path, actual.Words, expected.Words)
	}
}
