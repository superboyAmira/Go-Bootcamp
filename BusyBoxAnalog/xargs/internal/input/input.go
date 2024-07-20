package input

import (
	"bufio"
	"os"
)

func Input() (ret []string, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		ret = append(ret, line)
	}
	if scanner.Err() != nil {
		err = scanner.Err()
		return
	}
	// for split in future
	ret = append(ret, "")
	ret = append(ret, os.Args[1:]...)
	return
}
