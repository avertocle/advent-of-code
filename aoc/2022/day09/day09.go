package day09

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"math"
	"strings"
)

var inpDir []byte
var inpSteps []int
var inpLen int

func SolveP1() string {
	tRoute := make(map[string]int)
	var hx, hy, tx, ty int
	tRoute[toKey(tx, ty)] = 1
	for i := 0; i < inpLen; i++ {
		hx, hy = finalHPos(hx, hy, inpDir[i], inpSteps[i])
		for !areHTtouching(hx, hy, tx, ty) {
			tx, ty = nextTPos(tx, ty, hx, hy)
			tRoute[toKey(tx, ty)] = 1
		}
	}
	ans := len(tRoute)
	//printTroute(tRoute)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	tRoute := make(map[string]int)
	rx := make([]int, 10)
	ry := make([]int, 10)
	tRoute[toKey(rx[9], ry[9])] = 1
	for i := 0; i < inpLen; i++ {
		rx[0], ry[0] = finalHPos(rx[0], ry[0], inpDir[i], inpSteps[i])
		for !areHTtouching(rx[0], ry[0], rx[1], ry[1]) {
			for t := 1; t < len(rx); t++ {
				if !areHTtouching(rx[t-1], ry[t-1], rx[t], ry[t]) {
					rx[t], ry[t] = nextTPos(rx[t], ry[t], rx[t-1], ry[t-1])
				}
				tRoute[toKey(rx[9], ry[9])] = 1
			}
		}
	}
	ans := len(tRoute)
	//printTroute(tRoute)
	return fmt.Sprintf("%v", ans)
}

func finalHPos(x, y int, direc byte, steps int) (int, int) {
	switch direc {
	case 'L':
		return x - steps, y
	case 'R':
		return x + steps, y
	case 'U':
		return x, y + steps
	case 'D':
		return x, y - steps
	default:
		fmt.Printf("wrong move direction for H : %v", direc)
		return math.MaxInt, math.MaxInt
	}
}

// assumes they're not touching, check if they're touching before calling this
func nextTPos(tx, ty, hx, hy int) (int, int) {
	if tx == hx {
		if ty < hy {
			return tx, ty + 1
		} else {
			return tx, ty - 1
		}
	} else if ty == hy {
		if tx < hx {
			return tx + 1, ty
		} else {
			return tx - 1, ty
		}
	} else {
		if tx < hx && ty < hy {
			return tx + 1, ty + 1
		} else if tx < hx && ty > hy {
			return tx + 1, ty - 1
		} else if tx > hx && ty < hy {
			return tx - 1, ty + 1
		} else {
			return tx - 1, ty - 1
		}
	}
}

func areHTtouching(hx, hy, tx, ty int) bool {
	dx := intz.Abs(hx - tx)
	dy := intz.Abs(hy - ty)
	if (dx == 0 && dy <= 1) ||
		(dy == 0 && dx <= 1) ||
		(dx == 1 && dy == 1) {
		return true
	} else {
		return false
	}
}

func toKey(x, y int) string {
	return fmt.Sprintf("%v,%v", x, y)
}

func fromKey(k string) (int, int) {
	t := strings.Split(k, ",")
	return stringz.AtoI(t[0], -1), stringz.AtoI(t[1], -1)
}

func printTroute(m map[string]int) {
	size := 400
	arr := bytez.Init2D(size, size, '.')
	var x, y int
	for k, _ := range m {
		x, y = fromKey(k)
		arr[size/2-y][size/2+x] = '#'
	}
	arr[size/2][size/2] = 's'
	fmt.Println()
	fmt.Println()
	fmt.Println("deprecated")
	bytez.PPrint2D(arr)
}

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	inpDir = iutils.ExtractByte1DFromString1D(lines, " ", 0, 1)
	inpSteps = iutils.ExtractInt1DFromString1D(lines, " ", 1, 0)
	inpLen = len(inpDir)
	//fmt.Printf("%v\n", string(inpDir))
	//fmt.Printf("%v\n", inpSteps)

}
