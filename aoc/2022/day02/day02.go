package day02

import (
	"fmt"
	"github.com/avertocle/contests/io/iutils"
	"log"
	"strings"
)

// gInput : [][]byte : each row contains a move pair
var gInput [][]byte

type result int

const (
	lose result = iota
	draw
	win
)

func SolveP1() string {
	score := 0
	var res result
	var myMove, oppMove byte
	for _, row := range gInput {
		oppMove = row[0]
		myMove = row[1]
		res = calcMyMatchResultPart01(oppMove, myMove)
		score += calcMyRoundScore(myMove, res)
	}
	return fmt.Sprintf("%v", score)
}

func SolveP2() string {
	score := 0
	var res result
	var myMove, oppMove byte
	for _, row := range gInput {
		oppMove = row[0]
		res = calcResultPartO2(row[1])
		myMove = calcMyMovePart02(oppMove, res)
		score += calcMyRoundScore(myMove, res)
	}
	return fmt.Sprintf("%v", score)
}

/***** PART 01 Functions *****/

func calcMyMatchResultPart01(oppMove, myMove byte) result {
	s := string([]byte{oppMove, myMove})
	var res result
	if strings.Compare(s, "AY") == 0 ||
		strings.Compare(s, "BZ") == 0 ||
		strings.Compare(s, "CX") == 0 {
		res = win
	} else if strings.Compare(s, "AX") == 0 ||
		strings.Compare(s, "BY") == 0 ||
		strings.Compare(s, "CZ") == 0 {
		res = draw
	} else {
		res = lose
	}
	//fmt.Printf("%v %v\n", s, res)
	return res
}

/***** PART 02 Functions *****/

func calcResultPartO2(resCode byte) result {
	if resCode == 'X' {
		return lose
	} else if resCode == 'Y' {
		return draw
	} else if resCode == 'Z' {
		return win
	} else {
		return -1
	}
}

func calcMyMovePart02(oppMove byte, res result) byte {
	n := 10*int(oppMove-'A'+1) + int(res)
	m := make(map[int]byte)
	m[10] = 'C'
	m[11] = 'A'
	m[12] = 'B'
	m[20] = 'A'
	m[21] = 'B'
	m[22] = 'C'
	m[30] = 'B'
	m[31] = 'C'
	m[32] = 'A'
	return m[n]
}

/***** Common Functions *****/

func calcMyRoundScore(move byte, res result) int {
	score1 := calcOutcomeScore(res)
	score2 := calcMoveScore(move)
	//fmt.Printf("%v, %v\n", score1, score2)
	return score1 + score2
}

func calcOutcomeScore(res result) int {
	return int(res) * 3
}

func calcMoveScore(move byte) int {
	if move == 'A' || move == 'B' || move == 'C' {
		return int(move-'A') + 1
	} else if move == 'X' || move == 'Y' || move == 'Z' {
		return int(move-'X') + 1
	} else {
		fmt.Printf("calcMoveScore : invalid move : %v\n", move)
		return -1
	}
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	gInput = iutils.ExtractByte2DFromString1D(lines, " ", nil, 0)
}
