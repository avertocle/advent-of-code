package day07

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"sort"
	"strconv"
	"strings"
)

var gInput map[string]int
var gInputHands []string

const DirPath = "../2023/day07"

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
	cardStrengthMap := makeCardStrengthMapP1()
	calcHandType := calcHandTypeP1
	sortHands(gInputHands, cardStrengthMap, calcHandType)
	ans = calcTotalWinnings(gInputHands)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	cardStrengthMap := makeCardStrengthMapP2()
	calcHandType := calcHandTypeP2
	sortHands(gInputHands, cardStrengthMap, calcHandType)
	ans = calcTotalWinnings(gInputHands)
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func makeCardStrengthMapP1() map[byte]int {
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
func calcHandTypeP1(hand string) int {
	handKey, _ := getHandKeyAndJCount(hand)
	switch handKey {
	case "5":
		return FiveOfAKind
	case "14":
		return FourOfAKind
	case "23":
		return FullHouse
	case "113":
		return ThreeOfAKind
	case "122":
		return TwoPair
	case "1112":
		return OnePair
	case "11111":
		return HighCard
	}
	errz.HardAssert(false, "invalid hand(%v) handKey(%v)", hand, handKey)
	return -1
}

/***** P2 Functions *****/

func calcHandTypeP2(hand string) int {
	handKey, jcount := getHandKeyAndJCount(hand)
	handKey = fmt.Sprintf("%v-%v", handKey, jcount)
	switch handKey {
	case "5-0", "5-5", "14-1", "14-4":
		return FiveOfAKind
	case "14-0":
		return FourOfAKind
	case "23-0":
		return FullHouse
	case "23-2", "23-3":
		return FiveOfAKind
	case "113-0":
		return ThreeOfAKind
	case "113-1", "113-3":
		return FourOfAKind
	case "122-0":
		return TwoPair
	case "122-1":
		return FullHouse
	case "122-2":
		return FourOfAKind
	case "1112-0":
		return OnePair
	case "1112-1", "1112-2":
		return ThreeOfAKind
	case "11111-0":
		return HighCard
	case "11111-1":
		return OnePair
	}
	errz.HardAssert(false, "invalid hand(%v) cardKey(%v) handKey(%v)", hand, handKey)
	return -2
}

func makeCardStrengthMapP2() map[byte]int {
	cMap := makeCardStrengthMapP1()
	cMap['J'] = 1
	return cMap
}

func getHandKeyAndJCount(h string) (string, int) {
	cMap := makeCharCountMap([]byte(h))
	vals := make([]int, 0)
	for _, v := range cMap {
		vals = append(vals, v)
	}

	// handKey = string card-counts in ascending order
	sort.Ints(vals)
	key := ""
	for _, v := range vals {
		key += fmt.Sprintf("%v", v)
	}

	jcount := 0
	if v, ok := cMap['J']; ok {
		jcount = v
	}
	return key, jcount
}

/***** Common Functions *****/

func calcTotalWinnings(sortedHands []string) int {
	w := 0
	for i, h := range sortedHands {
		v := gInput[h]
		//fmt.Println(v, i+1, sortedHands[i])
		w += v * (i + 1)
	}
	return w
}

func sortHands(hands []string, cardStrengthMap map[byte]int, calcHandType func(string) int) {
	for i := 0; i < len(hands); i++ {
		for j := i + 1; j < len(hands); j++ {
			if compareHands(hands[i], hands[j], cardStrengthMap, calcHandType) > 0 {
				hands[i], hands[j] = hands[j], hands[i]
			}
		}
	}
}

func compareHands(h1, h2 string, cardStrengthMap map[byte]int, calcHandType func(string) int) int {
	h1Type := calcHandType(h1)
	h2Type := calcHandType(h2)
	//fmt.Println(h1, h2, h1Type, h2Type)
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
