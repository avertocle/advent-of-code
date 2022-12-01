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

func ExtractInt2DFromString1D(lines []string, sep string, cols []int, defaultVal int) [][]int {
	ans := make([][]int, len(lines))
	var temp []int
	for i, line := range lines {
		temp = ExtractInt1DFromString0D(line, sep, -1)
		ans[i] = make([]int, len(cols))
		for j, col := range cols {
			ans[i][j] = temp[col]
		}
	}
	return ans
}

func ExtractInt1DFromString1D(lines []string, sep string, col int, defaultVal int) []int {
	ans := make([]int, len(lines))
	var temp []int
	for i, line := range lines {
		temp = ExtractInt1DFromString0D(line, sep, defaultVal)
		ans[i] = temp[col]
	}
	return ans
}

func ExtractInt1DFromString0D(line string, sep string, defaultVal int) []int {
	var err error
	tokens := strings.Split(line, sep)
	ans := make([]int, len(tokens))
	for i, t := range tokens {
		ans[i], err = strconv.Atoi(t)
		if err != nil {
			fmt.Printf("strconv.Atoi failed for (%v) | err = %v | using default (%v) \n", t, err, defaultVal)
			ans[i] = defaultVal
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
