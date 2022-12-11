package day06

import (
	"fmt"
	"github.com/avertocle/contests/io/iutils"
	"log"
)

var gInput []byte

func ParseInput(inputFilePath string) {
	gInput = getInputOrDie(inputFilePath)
}

func SolveP1() string {
	ans := findFirstUniqSubseq(gInput, 4)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := findFirstUniqSubseq(gInput, 14)
	return fmt.Sprintf("%v", ans)
}

func PrintInputMetadata() {
	fmt.Printf("gInput length = %v\n", len(gInput))
}

/***** Common Functions *****/

func findFirstUniqSubseq(b1d []byte, window int) int {
	mStart := 0
	mEnd := mStart + window
	ans := -1
	for mEnd < len(b1d) {
		if areAllUniq(b1d[mStart:mEnd]) {
			ans = mEnd
			break
		}
		mStart++
		mEnd = mStart + window
	}
	return ans
}

func areAllUniq(b1d []byte) bool {
	set := make([]int, 26)
	var key byte
	for _, b := range b1d {
		key = b - 'a'
		set[key]++
		if set[key] > 1 {
			return false
		}
	}
	return true
}

/***** Input *****/

// gInput : []byte
func getInputOrDie(inputFilePath string) []byte {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	return []byte(lines[0])
}
