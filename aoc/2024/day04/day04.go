package day04

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

const DirPath = "../2024/day04"

var gInput [][]byte

func SolveP1() string {
	ans := 0
	for i := 0; i < len(gInput); i++ {
		for j := 0; j < len(gInput[0]); j++ {
			if gInput[i][j] == 'X' {
				ans += findAllXmasStartingAtX(i, j)
			}
		}
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	for i := 0; i < len(gInput); i++ {
		for j := 0; j < len(gInput[0]); j++ {
			if gInput[i][j] == 'A' && checkMasInXShapeCentredAtA(i, j) {
				ans++
			}
		}
	}
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func findAllXmasStartingAtX(r, c int) int {
	count := 0
	checks := []func(int, int) bool{checkXmasLeft, checkXmasRight, checkXmasUp, checkXmasDown,
		checkXmasTopLeft, checkXmasTopRight, checkXmasBottomLeft, checkXmasBottomRight}
	for _, check := range checks {
		if check(r, c) {
			count++
		}
	}
	return count
}

func checkXmasLeft(r, c int) bool {
	return c >= 3 && string(gInput[r][c-3:c+1]) == "SAMX"
}

func checkXmasRight(r, c int) bool {
	return c <= len(gInput[0])-4 && string(gInput[r][c:c+4]) == "XMAS"
}

func checkXmasUp(r, c int) bool {
	return r >= 3 && string([]byte{gInput[r-3][c], gInput[r-2][c], gInput[r-1][c], gInput[r][c]}) == "SAMX"
}

func checkXmasDown(r, c int) bool {
	return r <= len(gInput)-4 && string([]byte{gInput[r][c], gInput[r+1][c], gInput[r+2][c], gInput[r+3][c]}) == "XMAS"
}

func checkXmasTopLeft(r, c int) bool {
	return r >= 3 && c >= 3 && string([]byte{gInput[r-3][c-3], gInput[r-2][c-2], gInput[r-1][c-1], gInput[r][c]}) == "SAMX"
}

func checkXmasTopRight(r, c int) bool {
	return r >= 3 && c <= len(gInput[0])-4 && string([]byte{gInput[r-3][c+3], gInput[r-2][c+2], gInput[r-1][c+1], gInput[r][c]}) == "SAMX"
}

func checkXmasBottomLeft(r, c int) bool {
	return r <= len(gInput)-4 && c >= 3 && string([]byte{gInput[r+3][c-3], gInput[r+2][c-2], gInput[r+1][c-1], gInput[r][c]}) == "SAMX"
}

func checkXmasBottomRight(r, c int) bool {
	return r <= len(gInput)-4 && c <= len(gInput[0])-4 && string([]byte{gInput[r+3][c+3], gInput[r+2][c+2], gInput[r+1][c+1], gInput[r][c]}) == "SAMX"
}

/***** P2 Functions *****/

func checkMasInXShapeCentredAtA(r, c int) bool {
	return r >= 1 && c >= 1 && c <= len(gInput[0])-2 && r <= len(gInput)-2 && checkMasLeftDiagonal(r, c) && checkMasRightDiagonal(r, c)
}

func checkMasLeftDiagonal(r, c int) bool {
	str := string([]byte{gInput[r-1][c-1], gInput[r][c], gInput[r+1][c+1]})
	return str == "MAS" || str == "SAM"
}

func checkMasRightDiagonal(r, c int) bool {
	str := string([]byte{gInput[r-1][c+1], gInput[r][c], gInput[r+1][c-1]})
	return str == "MAS" || str == "SAM"
}

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
}

/*
M*M
*A*
S*S

M*S
*A*
M*S

S*M
*A*
S*M

S*S
*A*
M*M
*/
