package day11

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/numz"
	"math"
)

var gInputUni [][]byte
var gInputGalaxies [][]int

const DirPath = "../2023/day11"

func SolveP1() string {
	ans := 0
	expandedGalaxies := getExpandedGalaxies(gInputUni, gInputGalaxies, 2)
	ans = calcSumOfShortestPaths(expandedGalaxies)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	expandedGalaxies := getExpandedGalaxies(gInputUni, gInputGalaxies, 1000000)
	ans = calcSumOfShortestPaths(expandedGalaxies)
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

func calcSumOfShortestPaths(galaxies [][]int) int {
	ans := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			spLen := calcShortestPath(galaxies[i], galaxies[j])
			ans += spLen
		}
	}
	return ans
}

func calcShortestPath(start, end []int) int {
	return int(math.Abs(float64(start[0]-end[0])) + math.Abs(float64(start[1]-end[1])))
}

func getExpandedGalaxies(refUni [][]byte, galaxies [][]int, units int) [][]int {
	emptyRows, emptyCols := getEmptyRows(refUni), getEmptyCols(refUni)
	expGalaxies := make([][]int, 0)
	for _, galaxy := range galaxies {
		emptyRowsAbove := intz.CountLesser1D(emptyRows, galaxy[0])
		emptyColsLeft := intz.CountLesser1D(emptyCols, galaxy[1])
		expGalaxies = append(expGalaxies, []int{
			galaxy[0] + numz.Max(emptyRowsAbove*(units-1), 0),
			galaxy[1] + numz.Max(emptyColsLeft*(units-1), 0),
		})
		//fmt.Println(galaxy, emptyRowsAbove, emptyColsLeft, expGalaxies[len(expGalaxies)-1])
	}
	return expGalaxies
}

func getEmptyRows(refUni [][]byte) []int {
	emptyRows := make([]int, 0)
	for i := 0; i < len(refUni); i++ {
		isEmpty := true
		for j := 0; j < len(refUni[0]); j++ {
			if refUni[i][j] != '.' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyRows = append(emptyRows, i)
		}
	}
	return emptyRows
}

func getEmptyCols(refUni [][]byte) []int {
	emptyCols := make([]int, 0)
	for j := 0; j < len(refUni[0]); j++ {
		isEmpty := true
		for i := 0; i < len(refUni); i++ {
			if refUni[i][j] != '.' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyCols = append(emptyCols, j)
		}
	}
	return emptyCols
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInputUni = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
	gInputGalaxies = bytez.Find2D(gInputUni, '#')
}
