/*
- some funcs recieve error as an arg to be chainable, bad design but looks cool
- no error handling unless absolutely required
#competitive-programming #in-that-zone
*/

package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func String1DToInt2D(lines []string, sep string) [][]int {
	ans := make([][]int, len(lines))
	var temp []int
	for i, line := range lines {
		temp = StringToInt1D(line, sep)
		ans[i] = temp
	}
	return ans
}

func String1DToInt1D(lines []string, sep string) []int {
	ans := make([]int, len(lines))
	var temp []int
	for i, line := range lines {
		temp = StringToInt1D(line, sep)
		ans[i] = temp[0]
	}
	return ans
}

func StringToInt1D(line string, sep string) []int {
	var err error
	tokens := strings.Split(line, sep)
	ans := make([]int, len(tokens))
	for i, t := range tokens {
		ans[i], err = strconv.Atoi(t)
		if err != nil {
			err = fmt.Errorf("strconv.Atoi failed for (%v) | %v", t, err)
		}
	}
	return ans
}

func String1DToByte2D(lines []string) [][]byte {
	input := make([][]byte, len(lines))
	for i, row := range lines {
		input[i] = []byte(row)
	}
	return input
}
