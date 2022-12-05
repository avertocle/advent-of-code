package main

import (
	"fmt"
	"github.com/avertocle/contests/io"
	"github.com/avertocle/contests/io/ds"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"github.com/avertocle/contests/metrics"
	"log"
	"strings"
)

//added a . at end of all lines in input file because ide removes trailing space
const inputFilePath = "input.txt"

const zeroStr = "[0]"

func main() {
	metrics.ProgStart()
	iPiles, iMoves := getInputOrDie()

	//io.PrettyArray2DString(iPiles)
	//io.PrettyArray2DInt(iMoves)

	ans01 := solvePartOne(iPiles, iMoves)
	fmt.Printf("answer part-01 = %v\n", ans01)

	ans02 := solvePartTwo(iPiles, iMoves)
	fmt.Printf("answer part-02 = %v\n", ans02)
}

func solvePartOne(iPiles [][]string, iMoves [][]int) string {
	stacks := makeStacks(iPiles)
	for _, m := range iMoves {
		makeSingleMovePartOne(m[0], m[1]-1, m[2]-1, stacks)
	}
	ans := makeAnsString(stacks)
	return ans
}

func solvePartTwo(iPiles [][]string, iMoves [][]int) string {
	stacks := makeStacks(iPiles)
	for _, m := range iMoves {
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

/* input :
[][]string : piles
[][]int : moves
*/
func getInputOrDie() ([][]string, [][]int) {

	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}

	piles := make([][]string, 0)
	iMoves := make([][]int, 0)
	for _, l := range lines {
		if len(l) == 0 {
			continue
		} else if strings.HasPrefix(l, " ") || strings.HasPrefix(l, "[") {
			piles = append(piles, parsePileLine(l))
		} else if strings.HasPrefix(l, "m") {
			iMoves = append(iMoves, parseMovesLine(l))
		}
	}

	io.PrettyArray2DString(piles)

	// transform piles
	iPiles := make([][]string, len(piles[0]))
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
		iPiles[j] = temp
	}
	return iPiles, iMoves
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
