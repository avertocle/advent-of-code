package day05

import (
	"github.com/avertocle/contests/io/ds"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"strings"
)

//added a . at end of all lines in input file because ide removes trailing space
var gInpPiles [][]string
var gInpMoves [][]int

const zeroStr = "[0]"

func SolveP1() string {
	stacks := makeStacks(gInpPiles)
	for _, m := range gInpMoves {
		makeSingleMovePartOne(m[0], m[1]-1, m[2]-1, stacks)
	}
	ans := makeAnsString(stacks)
	return ans
}

func SolveP2() string {
	stacks := makeStacks(gInpPiles)
	for _, m := range gInpMoves {
		makeSingleMovePartTwo(m[0], m[1]-1, m[2]-1, stacks)
	}
	ans := makeAnsString(stacks)
	return ans
}

/***** Common Functions *****/

func makeStacks(iPiles [][]string) []*ds.Stack {
	stacks := make([]*ds.Stack, len(iPiles))
	for i, _ := range stacks {
		stacks[i] = ds.NewStack()
		stacks[i].PushAllRev(iPiles[i])
	}
	return stacks
}

func makeAnsString(stacks []*ds.Stack) string {
	ans := make([]byte, len(stacks))
	for i, s := range stacks {
		ans[i] = []byte(s.Peek())[1]
	}
	return string(ans)
}

/***** PART 01 Functions *****/

func makeSingleMovePartOne(count, src, des int, stacks []*ds.Stack) {
	temp := stacks[src].PopN(count)
	stacks[des].PushAll(temp)
}

/***** PART 02 Functions *****/

func makeSingleMovePartTwo(count, src, des int, stacks []*ds.Stack) {
	temp := stacks[src].PopN(count)
	stacks[des].PushAllRev(temp)
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)

	piles := make([][]string, 0)
	gInpMoves = make([][]int, 0)
	for _, l := range lines {
		if len(l) == 0 {
			continue
		} else if strings.HasPrefix(l, " ") || strings.HasPrefix(l, "[") {
			piles = append(piles, parsePileLine(l))
		} else if strings.HasPrefix(l, "m") {
			gInpMoves = append(gInpMoves, parseMovesLine(l))
		}
	}

	//stringz.PPrint2D(piles)

	// transform piles
	gInpPiles = make([][]string, len(piles[0]))
	var temp []string
	for j := 0; j < len(piles[0]); j++ {
		temp = make([]string, 0)
		for i := 0; i < len(piles)-1; i++ { // to ignore the number line
			if strings.Compare(piles[i][j], zeroStr) == 0 {
				continue
			} else {
				temp = append(temp, piles[i][j])
			}
		}
		gInpPiles[j] = temp
	}
}

func parsePileLine(l string) []string {
	ans := make([]string, 0)
	temp := ""
	for i := 0; i < len(l)-1; i += 4 { // len-1 to ignore last char (.)
		temp = l[i : i+3]
		if strings.Compare(temp, "   ") == 0 {
			temp = zeroStr
		}
		ans = append(ans, temp)
	}
	return ans
}

func parseMovesLine(l string) []int {
	tokens := strings.Split(l, " ")
	ans := []int{stringz.AtoiQ(tokens[1], -1),
		stringz.AtoiQ(tokens[3], -1),
		stringz.AtoiQ(tokens[5], -1)}
	return ans
}
