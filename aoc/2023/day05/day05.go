package day05

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"math"
	"strings"
)

var gInputSeedsP1 []int
var gInputSeedsP2 []int
var gInputMaps []*resMap

const DirPath = "../2023/day05"

func SolveP1() string {
	ans := 0
	minLocation := math.MaxInt
	for i := 0; i < len(gInputSeedsP1); i++ {
		location := gInputSeedsP1[i]
		for _, rm := range gInputMaps {
			location = rm.lookup(location)
		}
		if location < minLocation {
			minLocation = location
		}
	}
	ans = minLocation
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	minLocation := math.MaxInt
	for i := 0; i < len(gInputSeedsP1); i += 2 {
		seedStart := gInputSeedsP1[i]
		seedLen := gInputSeedsP1[i+1]
		location := 0
		for s := seedStart; s < seedStart+seedLen; s++ {
			location = s
			for _, rm := range gInputMaps {
				location = rm.lookup(location)
			}
			if location < minLocation {
				minLocation = location
			}
		}
		fmt.Println(minLocation)
	}
	ans = minLocation
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

type resMap struct {
	ranges []*resMapRange
}

func newResMap() *resMap {
	return &resMap{
		ranges: make([]*resMapRange, 0),
	}
}

func (rm *resMap) addRange(arr []int) {
	rm.ranges = append(rm.ranges, newResMapRange(arr))
}

func (rm *resMap) lookup(a int) int {
	for _, r := range rm.ranges {
		if x := r.lookup(a); x != -1 {
			return x
		}
	}
	return a
}

func (rm *resMap) String() string {
	ans := ""
	for _, r := range rm.ranges {
		ans += fmt.Sprintf("%v-%v : %v-%v | ", r.srStart, r.srStart+r.rlen-1, r.drStart, r.drStart+r.rlen-1)
	}
	return ans
}

type resMapRange struct {
	srStart int
	drStart int
	rlen    int
}

func newResMapRange(arr []int) *resMapRange {
	return &resMapRange{
		srStart: arr[1],
		drStart: arr[0],
		rlen:    arr[2],
	}
}

func (r *resMapRange) lookup(a int) int {
	if a >= r.srStart && a < r.srStart+r.rlen {
		return r.drStart + (a - r.srStart)
	}
	return -1
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	seedStr := strings.Split(lines[0], ": ")[1]
	gInputSeedsP1 = iutils.ExtractInt1DFromString0D(seedStr, " ", -1)
	gInputSeedsP2 = parseInputSeedsP2(seedStr)
	gInputMaps = make([]*resMap, 0)
	for i := 3; i < len(lines); i++ {
		resMap := newResMap()
		for ; i < len(lines) && len(lines[i]) != 0; i++ {
			resMap.addRange(iutils.ExtractInt1DFromString0D(lines[i], " ", -1))
		}
		gInputMaps = append(gInputMaps, resMap)
		i++
	}
	//fmt.Printf("seeds: %v\n", gInputSeedsP1)
	//for _, m := range gInputMaps {
	//	fmt.Printf("resmap: %+v\n", m)
	//}
}

func parseInputSeedsP2(line string) []int {
	seedTokens := iutils.ExtractInt1DFromString0D(line, " ", -1)
	seeds := make([]int, 0)
	for i := 0; i < len(seedTokens); i += 2 {
		seeds = append(seeds, seedTokens[i])
	}
	return seeds
}
