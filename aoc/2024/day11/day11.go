package day11

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/mapz"
	"github.com/avertocle/contests/io/stringz"
)

const DirPath = "../2024/day11"

var gInput []int

func SolveP1() string {
	ans := runSim(25)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := runSim(75)
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

func runSim(simCount int) int64 {
	stoneMap := mapz.FromArr1D(gInput, int64(1))
	for i := 0; i < simCount; i++ {
		stoneMap = runSimStep(stoneMap)
	}
	ans := mapz.SumValues(stoneMap)
	return ans
}

func runSimStep(stoneMap map[int]int64) map[int]int64 {
	newStoneMap := make(map[int]int64)
	for stone, count := range stoneMap {
		if stone == 0 {
			newStoneMap[1] += count
		} else if len(fmt.Sprintf("%d", stone))%2 == 0 {
			stone0, stone1 := makeSplitStones(stone)
			newStoneMap[stone0] += count
			newStoneMap[stone1] += count
		} else {
			newStoneMap[stone*2024] += count
		}
	}
	return newStoneMap
}

func makeSplitStones(stone int) (int, int) {
	str := fmt.Sprintf("%d", stone)
	mid := len(str) / 2
	stone0, stone1 := stringz.AtoI(str[0:mid], -1), stringz.AtoI(str[mid:], -1)
	return stone0, stone1
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractInt1DFromString0D(lines[0], " ", -1)
}
