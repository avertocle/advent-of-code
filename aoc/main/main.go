/*
Generic main.go for all aoc packages
Just change the import and dirPath
*/
package main

import (
	"fmt"
	prob "github.com/avertocle/contests/aoc/2021/day17"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/outils"
	"log"
	"path"
	"strings"
)

const dirPath = "../2021/day17"

func main() {
	displayPrettyHeader()
	inputFileNames, err := iutils.GetInputFileList(dirPath)
	if err != nil {
		log.Fatalf("error fetching input file : dir(%v) | %v", dirPath, err)
	}

	//inputFileNames = []string{"input_small.txt"}

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
	ansP2 := prob.SolveP2()
	return ansP1, ansP2
}

func displayPrettyHeader() {
	fmt.Println()
	fmt.Println()
	displayPrettyHorLine()
	fmt.Println(outils.Clr(fmt.Sprintf("Solving %v", dirPath), outils.Green))
	fmt.Println()
	fmt.Println(outils.Clr("<<< prog logs start >>>", outils.Yellow))
	fmt.Println()
}

func displayPrettyResult(ifName, ansP1, ansP2 string) {
	fmt.Println()
	fmt.Println(outils.Clr("<<< prog logs end >>>", outils.Yellow))
	fmt.Println()
	displayPrettyHorLine()
	fmt.Printf("%v : ansP1 = %v : ans-P2 = %v\n",
		outils.Clr(ifName, outils.Yellow),
		outils.Clr(ansP1, outils.Green),
		outils.Clr(ansP2, outils.Green))
	displayPrettyHorLine()
	fmt.Println()
}

func displayPrettyHorLine() {
	fmt.Println(outils.Clr(strings.Repeat("~-", 30)+"~", outils.Yellow))
}
