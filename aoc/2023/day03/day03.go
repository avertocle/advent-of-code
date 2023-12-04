package day03

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"strconv"
)

var gInput [][]byte

func SolveP1() string {
	ans := 0
	// gInput is padded with 1x '.' layer
	for i := 1; i < len(gInput)-1; i++ {
		for j := 1; j < len(gInput[0])-1; j++ {
			if !isDigit(gInput[i][j]) {
				continue
			}
			valBytes := make([]byte, 0)
			isPartNumber := false
			for ; j < len(gInput[0])-1; j++ {
				valBytes = append(valBytes, gInput[i][j])
				if hasSymbolsAround(gInput, i, j) {
					isPartNumber = true
				}
				if !isDigit(gInput[i][j+1]) {
					break
				}
			}
			if isPartNumber {
				val, _ := strconv.Atoi(string(valBytes))
				ans += val
			}
		}
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	parts := make([]*partz, 0)
	// gInput is padded with 1x '.' layer
	for i := 1; i < len(gInput)-1; i++ {
		for j := 1; j < len(gInput[0])-1; j++ {
			if !isDigit(gInput[i][j]) {
				continue
			}
			valBytes := make([]byte, 0)
			isPartNumber := false
			gears := make([][]int, 0)
			for ; j < len(gInput[0])-1; j++ {
				valBytes = append(valBytes, gInput[i][j])
				if hasSymbolsAround(gInput, i, j) {
					isPartNumber = true
				}
				gears = append(gears, findGearsAround(gInput, i, j)...)
				if !isDigit(gInput[i][j+1]) {
					break
				}
			}
			if isPartNumber {
				cleanedGears := cleanGearList(gears)
				//fmt.Println(string(valBytes), isPartNumber, cleanedGears)
				val, _ := strconv.Atoi(string(valBytes))
				parts = append(parts, &partz{
					val:   val,
					gears: cleanedGears,
				})
			}
		}
	}

	foo := make(map[string][]*partz)
	for _, part := range parts {
		for _, gear := range part.gears {
			key := fmt.Sprintf("r%vc%v", gear[0], gear[1])
			if foo[key] == nil {
				foo[key] = make([]*partz, 0)
			}
			foo[key] = append(foo[key], part)
		}
	}

	for _, parts := range foo {
		if len(parts) != 2 {
			continue
		}
		//fmt.Printf("%v.", len(parts))
		prod := 1
		for _, part := range parts {
			prod *= part.val
		}
		ans += prod
	}

	return fmt.Sprintf("%v", ans)
}

func cleanGearList(gears [][]int) [][]int {
	cleaned := make([][]int, 0)
	set := make(map[string]bool)
	for _, gear := range gears {
		key := fmt.Sprintf("r%vc%v", gear[0], gear[1])
		if !set[key] {
			cleaned = append(cleaned, gear)
			set[key] = true
		}
	}
	return cleaned
}

/***** Common Functions *****/

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isSymbol(b byte) bool {
	return !isDigit(b) && b != '.'
}

func isGear(b byte) bool {
	return b == '*'
}

func hasSymbolsAround(arr [][]byte, row, col int) bool {
	return isSymbol(arr[row-1][col-1]) || isSymbol(arr[row-1][col]) || isSymbol(arr[row-1][col+1]) ||
		isSymbol(arr[row][col-1]) || isSymbol(arr[row][col+1]) ||
		isSymbol(arr[row+1][col-1]) || isSymbol(arr[row+1][col]) || isSymbol(arr[row+1][col+1])
}

func findGearsAround(arr [][]byte, row, col int) [][]int {
	gears := make([][]int, 0)
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if isGear(arr[i][j]) {
				gears = append(gears, []int{i, j})
			}
		}
	}
	return gears
}

/***** P1 Functions *****/

/***** P2 Functions *****/

type partz struct {
	val   int
	gears [][]int
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	input := iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
	gInput = bytez.Pad2D(input, len(input), len(input[0]), 1, '.')
	//bytez.PPrint2D(gInput)
}
