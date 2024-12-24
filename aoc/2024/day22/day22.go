package day22

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/numz"
	"github.com/avertocle/contests/io/tpz"
)

const DirPath = "../2024/day22"

var gInput []int64

func SolveP1() string {
	ans := int64(0)
	simTime := 2001
	for _, s := range gInput {
		sAll, _ := evolveNTimes(s, simTime)
		ans += sAll[simTime-1]
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	simTime := 2001
	diffMapsAll := make([]map[string]int, len(gInput))
	uniqueDiffs := make(map[string]bool)
	for i, s := range gInput {
		dm := makeDiffMap(s, simTime, uniqueDiffs)
		diffMapsAll[i] = dm
	}
	bCount, maxBCount := 0, 0
	for k, _ := range uniqueDiffs {
		bCount = 0
		for _, dm := range diffMapsAll {
			if v, ok := dm[k]; ok {
				bCount += v
			}
		}
		maxBCount = numz.Max(maxBCount, bCount)
	}
	ans := maxBCount
	return fmt.Sprintf("%v", ans)
}

// also populates uniqueDiffs for minor optimization
func makeDiffMap(s int64, simTime int, uniqueDiffs tpz.Set[string]) map[string]int {
	_, pAll := evolveNTimes(s, simTime+1)
	diffMap := make(map[string]int)
	for i := 1; i < simTime-2; i++ {
		k := fmt.Sprintf("%v,%v,%v,%v",
			pAll[i]-pAll[i-1], pAll[i+1]-pAll[i],
			pAll[i+2]-pAll[i+1], pAll[i+3]-pAll[i+2])
		// monkey sells at the first matching sequence
		if _, ok := diffMap[k]; !ok {
			diffMap[k] = pAll[i+3]
		}
		uniqueDiffs[k] = true
	}
	return diffMap
}

func evolveNTimes(s int64, N int) ([]int64, []int) {
	ansS, ansP := make([]int64, N), make([]int, N)
	ansS[0], ansP[0] = s, int(s%10)
	for i := 1; i < N; i++ {
		ansS[i] = evolveNumber(ansS[i-1])
		ansP[i] = int(ansS[i] % 10)
	}
	return ansS, ansP
}

func evolveNumber(s int64) int64 {
	ans := prune(mix(s, 64*s))
	ans = prune(mix(ans, ans/32))
	ans = prune(mix(ans, ans*2048))
	return ans
}

func mix(s, m int64) int64 {
	return s ^ m
}

func prune(s int64) int64 {
	return s % 16777216
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractInt641DFromString1D(lines, "", -1, 0)
}
