package input

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

/*
recieves error as an arg to be chainable, bad design but looks cool
#compatitive-programming #in-that-zone
*/
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

func FromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("os.Open failed | %v", err)
	}
	defer file.Close()

	// optionally, resize scanner's capacity for lines over 64K, see next example
	scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
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
		fmt.Printf("%s\n", inputStr)
		return inputStr, nil
	}
}
