package day09

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

const DirPath = "../2024/day09"

var gFile []byte
var gFree []byte

func SolveP1() string {
	var ans int64
	fileSystem := makeFileSystem()
	defragmentP1(fileSystem)
	ans = calcChecksum(fileSystem)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	var ans int64
	fileSystem := makeFileSystem()
	defragmentP2(fileSystem)
	ans = calcChecksum(fileSystem)
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func defragmentP1(expFs []int) {
	for s, e := 0, len(expFs)-1; s < e; {
		for s < e && expFs[s] > -1 {
			s++
		}
		for e > s && expFs[e] == -1 {
			e--
		}
		expFs[s], expFs[e] = expFs[e], expFs[s]
	}
}

/***** P2 Functions *****/

func defragmentP2(expFs []int) {
	maxFileId := len(gFile) - 1
	for id := maxFileId; id >= 0; id-- {
		fileChunks := arrz.FindRepeatedByVal1D(expFs, id, nil, 0, 1)
		errz.HardAssert(len(fileChunks) == 1, "exactly one chunk must be found")
		fileChunk, fileChunkLen := fileChunks[0], (fileChunks[0][1]-fileChunks[0][0])+1
		biggerFreeChunks := arrz.FindRepeatedByVal1D(expFs, -1, []int{0, fileChunk[0]}, fileChunkLen, 1)
		if len(biggerFreeChunks) > 0 {
			arrz.SwapRangesInPlace1D(expFs, fileChunk, biggerFreeChunks[0])
		}
	}
}

/***** Common Functions *****/

func calcChecksum(expFs []int) int64 {
	var cs int64
	for i := 0; i < len(expFs); i++ {
		if expFs[i] > -1 {
			cs += int64(i * expFs[i])
		}
	}
	return cs
}

func makeFileSystem() []int {
	fs := make([]int, 0)
	for i := 0; i < len(gFile); i++ {
		file, free := gFile[i], gFree[i]
		for b := 0; b < int(file-'0'); b++ {
			fs = append(fs, i)
		}
		for f := 0; f < int(free-'0'); f++ {
			fs = append(fs, -1)
		}
	}
	return fs
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	line := []byte(lines[0])
	gFile = make([]byte, len(line)/2+1)
	gFree = make([]byte, len(line)/2+1)
	for i := 0; i < len(line)-1; i += 2 {
		gFile[i/2] = lines[0][i]
		gFree[i/2] = lines[0][i+1]
	}
	if len(line)%2 == 1 {
		gFile[len(line)/2] = lines[0][len(line)-1]
		gFree[len(line)/2] = '0'
	}
}
