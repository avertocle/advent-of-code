package iutils

import (
	"bufio"
	"fmt"
	"os"
)

func FromFile(path string, skipEmpty bool) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("os.Open failed | %v", err)
	}
	defer file.Close()

	// optionally, resize scanner's capacity for lines over 64K, see next example
	scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
		if !skipEmpty {
			input = append(input, line)
		} else if len(line) > 0 {
			input = append(input, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner.Err() failed | %v", err)
	}
	return input, nil
}
