package day21

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"math"
)

const DirPath = "../2024/day21"

var gInput [][]byte

func SolveP1() string {
	ans := 0
	//perms := generateAllPerms("aacd")
	//fmt.Println(len(perms), perms)
	manDisNum, manDisDir := makeManDisMaps()
	//fmt.Println(manDisNum)
	//fmt.Println(manDisDir)
	dirRobotCount := 2
	for _, c := range gInput {
		r0 := []string{string(c)}
		fmt.Println("r0", len(r0))
		r1 := findAllRobotSeq(r0, manDisNum)
		fmt.Println("rn", len(r1), r1)
		for i := 0; i < dirRobotCount; i++ {
			r1 = findAllRobotSeq(r1, manDisDir)
			fmt.Println("rd-", i, len(r1))
		}
		//r2 := findAllRobotSeq(r1, manDisDir)
		//fmt.Println("r2", len(r2))
		//r3 := findAllRobotSeq(r2, manDisDir)
		//fmt.Println("r3", len(r3))
		//fmt.Printf("\n\n%v\n%v\n%v\n%v\n\n", r0, r1, r2, r3)
		minLen := findMinLength(r1)
		ans += minLen * numPart(c)
		fmt.Println("==>", r0, minLen, numPart(c), ans)
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	return fmt.Sprintf("%v", ans)
}

func findMinLength(seq []string) int {
	minLen := math.MaxInt32
	for _, s := range seq {
		if len(s) < minLen {
			minLen = len(s)
		}
	}
	return minLen
}

func numPart(c []byte) int {
	ans := 0
	c = c[:len(c)-1]      // remove the last 'A'
	for i, b := range c { // skip the last 'A'
		ans += int(b-'0') * int(math.Pow(float64(10), float64(len(c)-i-1)))
	}
	return ans

}

func findAllRobotSeq(seq []string, m map[string][]string) []string {
	all := make([]string, 0)
	for _, s := range seq {
		all = append(all, findRobotSeq(s, m)...)
	}
	return all
}

func findRobotSeq(seq string, m map[string][]string) []string {
	//fmt.Printf("seq : %v | ", seq)
	seq = "A" + seq
	allNewSeq := make([]string, 0)
	allNewSeq = append(allNewSeq, "")
	for i := 1; i < len(seq); i++ {
		xs := m[toKey(seq[i-1], seq[i])]
		if len(xs) == 0 {
			xs = []string{""}
		}
		//fmt.Printf("xs : %v | ", xs)
		errz.HardAssert(len(xs) > 0, "no seq found for %v", toKey(seq[i-1], seq[i]))
		newAllNewSeq := make([]string, 0)
		for _, x := range xs {
			zz := x + "A"
			newAllNewSeq = append(newAllNewSeq, zz)
			//fmt.Printf("newAllNewSeq : %v | ", newAllNewSeq)
		}
		allNewSeq = crossProduct(allNewSeq, newAllNewSeq)
		//fmt.Printf("allNewSeq : %v | ", allNewSeq)
	}
	//fmt.Printf("\n")
	return allNewSeq
}

func crossProduct(a, b []string) []string {
	ans := make([]string, 0)
	for _, x := range a {
		for _, y := range b {
			ans = append(ans, x+y)
		}
	}
	return ans
}

func findRobotSeqV1(seq []byte, m map[string]string) []byte {
	newSeq := ""
	seq = append([]byte{'A'}, seq...)
	for i := 1; i < len(seq); i++ {
		x := m[toKey(seq[i-1], seq[i])] + "A"
		newSeq += x
		fmt.Printf("%v [%v] | ", toKey(seq[i-1], seq[i]), x)
	}
	fmt.Println()
	return []byte(newSeq)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

/***** Input *****/

func toKey(i, j byte) string {
	return string([]byte{i, j})
}

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
}
