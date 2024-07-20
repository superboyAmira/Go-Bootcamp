package main

import (
	"goday02/src/xargs/internal/input"
	"os"
	"os/exec"
)

// ./find -f -ext 'go' ../../../find/test_dir | ./xargs ./wc -l
func main() {
	slice, err := input.Input()
	if err != nil {
		return
	}

	var commandName string
	var args []string
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == "" {
			commandName = slice[i+1]
			args = slice[(i + 2):]
			break
		}
	}
	for _, line := range slice {
		if line == "" {
			break
		}
		args = append(args, line)
	}

	cmd := exec.Command(commandName, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return
	}
}
