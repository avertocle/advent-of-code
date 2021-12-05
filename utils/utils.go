package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func FetchInputFromFileAsIntArray(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("os.Open failed | %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	input := make([]int, 0)
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			err = fmt.Errorf("strconv.Atoi failed for (%v) | %v", v, err)
			return nil, err
		}
		input = append(input, v)
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
