/*
- some funcs recieve error as an arg to be chainable, bad design but looks cool
- no error handling unless absolutely required
#competitive-programming #in-that-zone
*/

package io

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func AsIntArray(lines []string, err error) ([]int, error) {
	if err != nil {
		return nil, fmt.Errorf("input had error | %v", err)
	}
	input := make([]int, 0)
	for i := 0; i < len(lines); i++ {
		v, err := strconv.Atoi(lines[i])
		if err != nil {
			err = fmt.Errorf("strconv.Atoi failed for (%v) | %v", v, err)
			return nil, err
		}
		input = append(input, v)
	}
	return input, nil
}

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

func FetchInputFromFileRaw() ([]byte, error) {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		err = fmt.Errorf("ioutil.ReadFile failed | %v", err)
		return nil, err
	}
	return input, nil
}

func FetchInputFromWeb(day string) ([]byte, error) {
	url := fmt.Sprintf("https://adventofcode.com/2021/day/%v/input", day)
	var res *http.Response
	var inputStr []byte
	var err error
	if res, err = http.Get(url); err != nil {
		err = fmt.Errorf("http.Get failed | %v", err)
		return nil, err
	} else if inputStr, err = ioutil.ReadAll(res.Body); err != nil {
		err = fmt.Errorf("ioutil.ReadAll failed | %v", err)
		return nil, err
	} else {
		defer res.Body.Close()
		return inputStr, nil
	}
}

func SplitToIntArray(line string, sep string) []int {
	tokens := strings.Split(line, sep)
	if sep == " " {
		tokens = strings.Fields(line)
	}
	ans := make([]int, len(tokens))
	var err error
	for i, t := range tokens {
		ans[i], err = strconv.Atoi(t)
		if err != nil {
			fmt.Printf("strconv.Atoi failed for (%v) (%v) | %v", i, t, err)
			fmt.Println(strings.Join(tokens, "|"))
		}
	}
	return ans
}

func MaxInt(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func MinInt(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func CountIn2dByteIf(grid [][]byte, f func(byte, int, int) bool) int {
	count := 0
	for i, row := range grid {
		for j, cell := range row {
			if f(cell, i, j) {
				count++
			}
		}
	}
	return count
}

func Init2DByte(rows, cols int) [][]byte {
	ans := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		ans[i] = make([]byte, cols)
	}
	return ans
}

func Init2DInit(rows, cols int) [][]int {
	ans := make([][]int, rows)
	for i := 0; i < rows; i++ {
		ans[i] = make([]int, cols)
	}
	return ans
}
