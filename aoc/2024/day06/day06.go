package day06

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

const DirPath = "../2024/day06"

var gInput [][]byte

func SolveP1() string {
	ans := 0
	guard, visited, isOut, _ := initStuff()
	for {
		markVisited(guard, visited, true)
		if guard, isOut, _ = moveGuard(guard, visited); isOut {
			break
		}
	}
	ans = len(visited)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	for i := 0; i < len(gInput); i++ {
		for j := 0; j < len(gInput[0]); j++ {
			if gInput[i][j] == '#' || gInput[i][j] == '^' {
				continue
			}
			gInput[i][j] = '#'
			guard, visited, isOut, isInLoop := initStuff()
			for {
				markVisited(guard, visited, false)
				if guard, isOut, isInLoop = moveGuard(guard, visited); isOut {
					break
				} else if isInLoop {
					ans++
					break
				}
			}
			gInput[i][j] = '.'
		}
	}
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

func initStuff() ([]int, map[string]bool, bool, bool) {
	// guard is an array of 3 elements : [row, col, direction]
	guard := append(arrz.Find2D(gInput, '^')[0], 'u')
	return guard, make(map[string]bool), false, false
}

func moveGuard(guard []int, visited map[string]bool) ([]int, bool, bool) {
	newGuardPos := moveFwd(guard)
	if isOutOfBounds(newGuardPos) {
		return nil, true, false
	} else if hasLooped(newGuardPos, visited) {
		return nil, false, true
	} else if hasObstacle(newGuardPos) {
		newGuardPos = turnRight(guard)
	}
	if hasLooped(newGuardPos, visited) {
		return nil, false, true
	}
	return newGuardPos, false, false
}

func markVisited(gp []int, visited map[string]bool, ignoreDirection bool) {
	if ignoreDirection {
		visited[fmt.Sprintf("%v", gp[:2])] = true
	} else {
		visited[fmt.Sprintf("%v", gp)] = true
	}
}

func hasLooped(gp []int, visited map[string]bool) bool {
	return visited[fmt.Sprintf("%v", gp)]
}

func isOutOfBounds(gp []int) bool {
	return gp[0] < 0 || gp[0] >= len(gInput) || gp[1] < 0 || gp[1] >= len(gInput[0])
}

func hasObstacle(gp []int) bool {
	return gInput[gp[0]][gp[1]] == '#'
}

func turnRight(gp []int) []int {
	m := map[int]int{'u': 'r', 'd': 'l', 'l': 'u', 'r': 'd'}
	return []int{gp[0], gp[1], m[gp[2]]}
}

func moveFwd(gp []int) []int {
	switch gp[2] {
	case 'u':
		return []int{gp[0] - 1, gp[1], 'u'}
	case 'd':
		return []int{gp[0] + 1, gp[1], 'd'}
	case 'l':
		return []int{gp[0], gp[1] - 1, 'l'}
	case 'r':
		return []int{gp[0], gp[1] + 1, 'r'}
	default:
		errz.HardAssert(false, "invalid guardPos : %v", gp)
		return nil
	}
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)

}
