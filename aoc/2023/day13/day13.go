package day13

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
)

var gInput [][][]byte

const DirPath = "../2023/day13"

func SolveP1() string {
	ans := 0
	totalMr, totalMc := 0, 0
	for _, pattern := range gInput {
		mr, mc, err := findMirror(pattern)
		if err != nil {
			errz.HardAssert(false, "no mirror found [%v,%v] | %v", mr, mc, err)
		} else if len(mr) > 0 && len(mc) > 0 {
			errz.HardAssert(false, "both row anc col mirror found [%v,%v] | %v", mr, mc, err)
		} else if len(mr) > 0 {
			totalMr += mr[0]
		} else if len(mc) > 0 {
			totalMc += mc[0]
		}
	}
	ans = totalMr*100 + totalMc
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	totalMr, totalMc := 0, 0
	for p, pattern := range gInput {
		mrOrig, mcOrig, errOrig := findMirror(pattern)
		if errOrig != nil {
			errz.HardAssert(false, "mirror conflict [%v,%v] | %v", mrOrig, mcOrig, errOrig)
		}
		isFound := false
		for i := 0; i < len(pattern) && !isFound; i++ {
			for j := 0; j < len(pattern[i]) && !isFound; j++ {
				pattern[i][j] = flipByte(pattern[i][j])
				mr, mc, err := findMirror(pattern)
				if err == nil {
					for k := 0; k < len(mr); k++ {
						if mr[k] > 0 && (len(mrOrig) == 0 || mr[k] != mrOrig[0]) {
							isFound = true
							fmt.Printf("coord[%v, %v], orig[%v, %v], curr[%v,%v] | %v\n", i, j, mrOrig, mcOrig, mr, mc, err)
							totalMr += mr[k]
							break
						}
					}
					for k := 0; k < len(mc); k++ {
						if mc[k] > 0 && (len(mcOrig) == 0 || mc[k] != mcOrig[0]) {
							isFound = true
							fmt.Printf("coord[%v, %v], orig[%v, %v], curr[%v,%v] | %v\n", i, j, mrOrig, mcOrig, mr, mc, err)
							totalMc += mc[k]
							break
						}
					}
				}
				pattern[i][j] = flipByte(pattern[i][j])
			}
		}
		if !isFound {
			//arrz.PPrint2D(pattern)
			fmt.Println("no mirror found\n\n", p, mrOrig, mcOrig)
			//totalMr += mrOrig
			//totalMc += mcOrig
			//break
		}
	}
	fmt.Println(totalMr, totalMc)
	ans = totalMr*100 + totalMc
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

func findMirror(pattern [][]byte) ([]int, []int, error) {
	mrs := findHorizontalMirrors(pattern)
	mcs := findVerticalMirror(pattern)

	if len(mrs) == 0 && len(mcs) == 0 {
		return nil, nil, fmt.Errorf("no mirror found")
	} else {
		intz.Map1D(mrs, func(arr []int, i int) int { return arr[i] + 1 })
		intz.Map1D(mcs, func(arr []int, i int) int { return arr[i] + 1 })
		return mrs, mcs, nil
	}
}

func flipByte(b byte) byte {
	if b == '#' {
		return '.'
	} else if b == '.' {
		return '#'
	}
	errz.HardAssert(false, "invalid byte %v", string(b))
	return b
}

func findHorizontalMirrors(arr [][]byte) []int {
	ans := make([]int, 0)
	for i := 0; i < len(arr)-1; i++ {
		if areIdenticalRows(arr, i, i+1) {
			for d := 1; ; d++ {
				if i-d < 0 || i+d+1 >= len(arr) {
					ans = append(ans, i)
					break
				} else if !areIdenticalRows(arr, i-d, i+d+1) {
					break
				}
			}
		}
	}
	return ans
}

func findVerticalMirror(arr [][]byte) []int {
	temp := bytez.RotateClockwise2D(arr)
	return findHorizontalMirrors(temp)
}

func areIdenticalRows(arr [][]byte, r1, r2 int) bool {
	for j := 0; j < len(arr[r1]); j++ {
		if arr[r1][j] != arr[r2][j] {
			//fmt.Printf("%v | %v | %v\n", string(arr[r1]), string(arr[r2]), "no")
			return false
		}
	}
	//fmt.Printf("%v | %v | %v\n", string(arr[r1]), string(arr[r2]), "yes")
	return true
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = make([][][]byte, 0)
	temp := make([]string, 0)
	for _, line := range lines {
		if len(line) == 0 {
			gInput = append(gInput, iutils.ExtractByte2DFromString1D(temp, "", nil, 0))
			temp = make([]string, 0)
		} else {
			temp = append(temp, line)
		}
	}
	gInput = append(gInput, iutils.ExtractByte2DFromString1D(temp, "", nil, 0))

	//for i := 0; i < len(lines); i++ {
	//	fmt.Println(lines[i])
	//}
	//
	//fmt.Println(len(gInput))
	//bytez.PPrint3D(gInput)
}
