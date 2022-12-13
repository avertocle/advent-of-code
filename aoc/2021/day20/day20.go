package day20

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/iutils"
	"log"
	"strconv"
	"strings"
)

var gInpImg [][]byte
var gInpLen int
var gInpAlg []byte

func SolveP1() string {
	ans := applyAlgo(gInpImg, 2)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := applyAlgo(gInpImg, 50)
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func applyAlgo(img [][]byte, n int) int {
	padVal := byte('.') // initially pad with dark pixels
	padSize := 3        // padSize = inspect-section-size = 3 is enough to stablize image edges
	for i := 0; i < n; i++ {
		img = bytez.Pad2D(img, len(img), len(img), padSize, padVal)
		img = enhance(img, padVal)
		padVal = img[0][0] // new padding should be what pixel edges have stablized to
		//fmt.Printf("litCount = %v\n", bytez.Count2D(img, '#'))
	}
	litCount := bytez.Count2D(img, '#')
	return litCount
}

func enhance(img [][]byte, padVal byte) [][]byte {
	var algSec [][]byte
	var algIdx int
	imgNew := bytez.Init2D(len(img), len(img), 0) // this padVal doesn't matter
	for i, row := range img {
		for j, _ := range row {
			algSec = bytez.ExtractSq2D(img, []int{i, j}, 3, padVal)
			algIdx = calcAlgIdx(algSec)
			imgNew[i][j] = gInpAlg[algIdx]
		}
	}
	return imgNew
}

func calcAlgIdx(arr [][]byte) int {
	s := ""
	for _, row := range arr {
		s += string(row)
	}
	s = strings.ReplaceAll(s, ".", "0")
	s = strings.ReplaceAll(s, "#", "1")
	ans, err := strconv.ParseInt(s, 2, 32)
	if err != nil {
		fmt.Printf("error : calcAlgIdx %v | %v\n", s, err)
	}
	return int(ans)
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	gInpAlg = []byte(lines[0])
	gInpImg = iutils.ExtractByte2DFromString1D(lines[2:], "", nil, 0)
	gInpLen = len(lines) - 2
}
