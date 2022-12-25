/*
Generic main.go for all aoc packages
Just change the import and dirPath
*/
package main

import (
	"fmt"
	prob "github.com/avertocle/contests/aoc/2022/day21"
	"github.com/avertocle/contests/io/clr"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"path"
	"strings"
)

const dirPath = "../2022/day21"

func main() {
	displayPrettyHeader()
	inputFileNames, err := iutils.GetInputFileList(dirPath)
	errz.HardAssert(err == nil, "error fetching input file : dir(%v) | %v", dirPath, err)

	//inputFileNames = []string{"input_small.txt"}
	//inputFileNames = []string{"input_final.txt"}

	var ansP1, ansP2 string
	var ifPath string
	for _, ifName := range inputFileNames {
		ifPath = path.Join(dirPath, ifName)
		ansP1, ansP2 = runForOneInputFile(ifPath)
		displayPrettyResult(ifName, ansP1, ansP2)
	}
}

/***** Display Functions *****/

func runForOneInputFile(inputFilePath string) (string, string) {
	prob.ParseInput(inputFilePath)
	ansP1 := prob.SolveP1()
	//fmt.Println("P1 Done")
	ansP2 := prob.SolveP2()
	//fmt.Println("P2 Done")
	return ansP1, ansP2
}

func displayPrettyHeader() {
	fmt.Println()
	fmt.Println()
	displayPrettyHorLine()
	fmt.Println(clr.Str(fmt.Sprintf("Solving %v", dirPath), clr.Green))
	fmt.Println()
	fmt.Println(clr.Str("<<< prog logs start >>>", clr.Yellow))
	fmt.Println()
}

func displayPrettyResult(ifName, ansP1, ansP2 string) {
	fmt.Println()
	fmt.Println(clr.Str("<<< prog logs end >>>", clr.Yellow))
	fmt.Println()
	displayPrettyHorLine()
	fmt.Printf("%v : ansP1 = %v : ans-P2 = %v\n",
		clr.Str(ifName, clr.Yellow),
		clr.Str(ansP1, clr.Green),
		clr.Str(ansP2, clr.Green))
	displayPrettyHorLine()
	fmt.Println()
}

func displayPrettyHorLine() {
	fmt.Println(clr.Str(strings.Repeat("~-", 30)+"~", clr.Yellow))
}
