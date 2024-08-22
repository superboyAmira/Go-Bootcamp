package main

import (
	"fmt"
	"os/exec"
)

func main() {
	testSleep := exec.Command("/bin/sh", "-c", "go test -v ../../ex00/sleepSort/sort_test.go ../../ex00/sleepSort/sort.go")
	testCrawler := exec.Command("/bin/sh", "-c", "go test -v ../../ex01/crawler/crawler_test.go ../../ex01/crawler/crawler.go")
	testOctopus := exec.Command("/bin/sh", "-c", "go test -v ../../ex02/octopus/octopus_test.go ../../ex02/octopus/octopus.go")
	
	out, err := testSleep.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(out))
	out, err = testCrawler.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(out))
	out, err = testOctopus.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(out))
}
