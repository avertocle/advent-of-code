package day21

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"math"
)

const DirPath = "../2024/day21"

var gInput [][]byte

const (
	A1     = 10
	Empty1 = 11

	up     = 0
	down   = 1
	left   = 2
	right  = 3
	A2     = 4
	Empty2 = 5
)

/*
+---+---+---+
| 7 | 8 | 9 |
+---+---+---+
| 4 | 5 | 6 |
+---+---+---+
| 1 | 2 | 3 |
+---+---+---+
	| 0 | A |
	+---+---+

    +---+---+
    | ^ | A |
+---+---+---+
| < | v | > |
+---+---+---+
*/

func SolveP1() string {
	ans := 0
	manDisNum, manDisDir := makeManDisMaps()
	//for k, v := range manDisDir {
	//	fmt.Printf("%v: %q\n", k, v)
	//}
	//fmt.Println(manDisNum)
	//fmt.Println(manDisDir)
	for _, c := range gInput {
		r1 := findRobotSeq(c, manDisNum)
		r2 := findRobotSeq(r1, manDisDir)
		r3 := findRobotSeq(r2, manDisDir)
		fmt.Printf("%v\n%v\n%v\n%v\n\n", string(c), string(r1), string(r2), string(r3))
		fmt.Println(len(r3), numPart(c))
		ans += len(r3) * numPart(c)
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	return fmt.Sprintf("%v", ans)
}

func numPart(c []byte) int {
	ans := 0
	c = c[:len(c)-1]      // remove the last 'A'
	for i, b := range c { // skip the last 'A'
		ans += int(b-'0') * int(math.Pow(float64(10), float64(len(c)-i-1)))
	}
	return ans

}

func findRobotSeq(seq []byte, itr int, m map[string]string) []byte {
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
