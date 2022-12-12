/*
Generic main.go for all aoc packages
Just change the import and dirPath
*/
package main

import (
	"fmt"
	prob "github.com/avertocle/contests/aoc/2021/day16"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/outils"
	"log"
	"path"
)

const dirPath = "../2021/day16"

func main() {

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
		displayPretty(ifPath, ansP1, ansP2)
	}
}

func runForOneInputFile(inputFilePath string) (string, string) {
	prob.ParseInput(inputFilePath)
	ansP1 := prob.SolveP1()
	ansP2 := prob.SolveP2()
	return ansP1, ansP2
}

func displayPretty(ifPath, ansP1, ansP2 string) {
	fmt.Printf("=> %v : ansP1 = %v : ans-P2 = %v\n",
		outils.Clr(ifPath, outils.Gray),
		outils.Clr(ansP1, outils.Green),
		outils.Clr(ansP2, outils.Green))

}
