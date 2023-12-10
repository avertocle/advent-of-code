package day07

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"strconv"
	"strings"
)

const DirPath = "../2023/day07"

var gInput map[string]int
var gInputHands []string

const (
	FiveOfAKind = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

func SolveP1() string {
	ans := 0
	cardStrengthMap := makeCardMapP1()
	sortHands(gInputHands, cardStrengthMap)
	ans = calcTotalWinnings(gInputHands)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	cardStrengthMap := makeCardMapP2()
	sortHands(gInputHands, cardStrengthMap)
	ans = calcTotalWinnings(gInputHands)
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func calcTotalWinnings(sortedHands []string) int {
	w := 0
	for i, h := range sortedHands {
		v := gInput[h]
		fmt.Println(v, i+1)
		w += v * (i + 1)
	}
	return w
}

func sortHands(hands []string, cardStrengthMap map[byte]int) {
	for i := 0; i < len(hands); i++ {
		for j := i + 1; j < len(hands); j++ {
			if compareHands(hands[i], hands[j], cardStrengthMap) > 0 {
				hands[i], hands[j] = hands[j], hands[i]
			}
		}
	}
}

func compareHands(h1, h2 string, cardStrengthMap map[byte]int) int {
	h1Type := calcHandType(h1)
	h2Type := calcHandType(h2)
	if h1Type > h2Type {
		return -1
	} else if h1Type < h2Type {
		return 1
	}

	// compare hands of same type
	bh1 := []byte(h1)
	bh2 := []byte(h2)
	for i := 0; i < len(bh1); i++ {
		// compare individual cards
		r := cardStrengthMap[bh1[i]] - cardStrengthMap[bh2[i]]
		if r == 0 {
			continue
		}
		return r
	}
	errz.HardAssert(false, "invalid hands | [%v], [%v]", h1, h2)
	return 0
	//return compareHandsSameType(h1, h2, h1Type, cardStrengthMap)
}

func calcHandType(h string) int {
	cMap := makeCharCountMap([]byte(h))
	if len(cMap) == 1 {
		return FiveOfAKind
	} else if len(cMap) == 2 {
		for _, v := range cMap {
			if v == 4 {
				return FourOfAKind
			} else if v == 3 {
				return FullHouse
			}
		}
	} else if len(cMap) == 3 {
		for _, v := range cMap {
			if v == 3 {
				return ThreeOfAKind
			} else if v == 2 {
				return TwoPair
			}
		}
	} else if len(cMap) == 4 {
		return OnePair
	} else {
		return HighCard
	}
	errz.HardAssert(false, "invalid hand | %v, cmap(%v)", h, len(cMap))
	return -1
}

func makeCharCountMap(bh []byte) map[byte]int {
	cMap := make(map[byte]int)
	for _, c := range bh {
		if _, ok := cMap[c]; !ok {
			cMap[c] = 0
		}
		cMap[c]++
	}
	return cMap
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = make(map[string]int)
	gInputHands = make([]string, len(lines))
	for i, line := range lines {
		tokens := strings.Fields(line)
		gInputHands[i] = tokens[0]
		gInput[tokens[0]], _ = strconv.Atoi(tokens[1])
	}
	//fmt.Println(gInputHands)
	//fmt.Println(gInput)
}

func makeCardMapP1() map[byte]int {
	cMap := make(map[byte]int)
	for i := 2; i < 10; i++ {
		cMap['0'+byte(i)] = i
	}
	cMap['T'] = 10
	cMap['J'] = 11
	cMap['Q'] = 12
	cMap['K'] = 13
	cMap['A'] = 14
	return cMap
}

func makeCardMapP2() map[byte]int {
	cMap := makeCardMapP1()
	cMap['J'] = 1
	return cMap
}
