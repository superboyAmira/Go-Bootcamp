package main

import (
	"fmt"
	"os/exec"
)

func main() {
	testEx00 := exec.Command("/bin/sh", "-c", "go test -v ../../internal/ex00/ex00_test.go ../../internal/ex00/ex00.go")
	testEx01 := exec.Command("/bin/sh", "-c", "go test -v ../../internal/ex01/ex01_test.go ../../internal/ex01/ex01.go")
	
	out, err := testEx00.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(out))
	out, err = testEx01.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(out))
}
