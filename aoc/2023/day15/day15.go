package day15

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"strings"
)

var gInput []string

const DirPath = "../2023/day15"

const (
	DASH = iota
	EQUAL
)

func SolveP1() string {
	ans := 0
	for _, s := range gInput {
		ans += hashStr([]byte(s))
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	focLenMap := make(map[string]int)
	boxes := make([][]string, 256)
	for _, s := range gInput {
		label, op, focLen := deconstructLensLabel(s)
		boxIdx := hashStr([]byte(label))
		if op == DASH {
			boxes[boxIdx] = executeDash(boxes[boxIdx], label)
		} else if op == EQUAL {
			boxes[boxIdx] = executeEqual(boxes[boxIdx], label, focLen, focLenMap)
		}
	}
	ans = calcTotalFocLen(boxes, focLenMap)
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

func deconstructLensLabel(step string) (string, int, int) {
	tokens := stringz.SplitMulti(step, []string{"-", "="})
	label := tokens[0]
	if strings.HasSuffix(step, "-") {
		return label, DASH, -1
	} else if strings.Contains(step, "=") {
		focLen := stringz.AtoI(tokens[1], -1)
		return label, EQUAL, focLen
	}
	errz.HardAssert(false, "invalid step | %v", step)
	return "", -1, -1
}

func calcTotalFocLen(boxes [][]string, focLenMap map[string]int) int {
	ans := 0
	for i, box := range boxes {
		if len(box) > 0 {
			t := calcFocLenOneBox(i, box, focLenMap)
			ans += t
		}
	}
	return ans
}

func calcFocLenOneBox(boxIdx int, box []string, lfMap map[string]int) int {
	fl := 0
	for i, lens := range box {
		fl = fl + ((i + 1) * lfMap[lens])
	}
	v := fl * (boxIdx + 1)
	//fmt.Printf("%v : [%v, %v], boxLen(%v) | lenses(%v) | %v\n", boxIdx, fl, v, len(box), box, lfMap)
	return v
}

func executeDash(box []string, label string) []string {
	indexes := stringz.Find1D(box, label)
	errz.HardAssert(len(indexes) <= 1, "executeDash : multiple indexes | [%v in %v]", box, label)
	if len(indexes) == 0 {
		return box
	}
	errz.HardAssert(len(box) > indexes[0], "executeDash : invalid index | [%v in %v]", box, label)
	if len(box) == 0 {
		return []string{label}
	}
	newBox := append(box[:indexes[0]], box[indexes[0]+1:]...)
	return newBox
}

func executeEqual(box []string, label string, focLen int, focLenMap map[string]int) []string {
	indexes := stringz.Find1D(box, label)
	focLenMap[label] = focLen
	if len(indexes) == 0 {
		newBox := append(box, label)
		return newBox
	} else if len(indexes) == 1 {
		return box
	}
	errz.HardAssert(len(indexes) > 1, "invalid box | box(%v), lens(%v)", box, label)
	return []string{}
}

/***** Common Functions *****/

func hashStr(seq []byte) int {
	hash := 0
	for _, b := range seq {
		hash = hashOneChar(b, hash)
	}
	return hash
}
func hashOneChar(b byte, currVal int) int {
	currVal += int(b)
	currVal *= 17
	currVal %= 256
	return currVal
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = strings.Split(lines[0], ",")
}
