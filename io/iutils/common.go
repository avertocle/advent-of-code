package iutils

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
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

func GetInputFileList(dirPath string) ([]string, error) {
	var files []fs.FileInfo
	var err error
	if files, err = ioutil.ReadDir(dirPath); err != nil {
		return nil, fmt.Errorf("ioutil.ReadDir failed | %v", err)
	}
	ans := make([]string, 0)
	for _, f := range files {
		if strings.HasPrefix(f.Name(), "input_") {
			ans = append(ans, f.Name())
		}
	}
	return ans, nil
}
