package day25

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/tpz"
	"strings"
)

const DirPath = "../2024/day25"

var gInput [][][]byte

const SizeI = 7
const SizeJ = 5
const Space = '.'
const Block = '#'

var SpaceChunk = strings.Repeat(string(Space), SizeJ)
var BlockChunk = strings.Repeat(string(Block), SizeJ)

func SolveP1() string {
	keys, locks := parseAll()
	combo := make(tpz.Set[string])
	for _, key := range keys {
		for _, lock := range locks {
			if lockFitsKey(lock, key) {
				combo[fmt.Sprintf("%v-%v", lock, key)] = true
			}
		}
	}
	ans := len(combo)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func lockFitsKey(lock, key []int) bool {
	for i := 0; i < SizeJ; i++ {
		if lock[i]+key[i] > SizeI {
			return false
		}
	}
	return true
}

/***** P2 Functions *****/

/***** Common Functions *****/

func parseAll() ([][]int, [][]int) {
	keys, locks := make([][]int, 0), make([][]int, 0)
	for _, arr := range gInput {
		if isKey(arr) {
			keys = append(keys, parseKey(arr))
		} else if isLock(arr) {
			locks = append(locks, parseLock(arr))
		}
	}
	return keys, locks
}

func isLock(arr [][]byte) bool {
	return string(arr[0]) == BlockChunk && string(arr[len(arr)-1]) == SpaceChunk
}

func isKey(arr [][]byte) bool {
	return string(arr[0]) == SpaceChunk && string(arr[len(arr)-1]) == BlockChunk
}

func parseKey(arr [][]byte) []int {
	ans := make([]int, SizeJ)
	for j := 0; j < SizeJ; j++ {
		for i := 0; i < SizeI; i++ {
			if arr[i][j] == Block {
				ans[j] = SizeI - i
				break
			}
		}
	}
	return ans
}

func parseLock(arr [][]byte) []int {
	ans := make([]int, SizeJ)
	for j := 0; j < SizeJ; j++ {
		for i := 0; i < SizeI; i++ {
			if arr[i][j] == Space {
				ans[j] = i
				break
			}
		}
	}
	return ans
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	sections := iutils.BreakByEmptyLineString1D(lines)
	gInput = make([][][]byte, 0)
	for _, section := range sections {
		gInput = append(gInput, iutils.ExtractByte2DFromString1D(section, "", nil, 0))
	}
}
