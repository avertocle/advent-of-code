package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/avertocle/adventofcode/io"
	"github.com/avertocle/adventofcode/metrics"
)

const inputFilePath = "input.txt"

func main() {
	metrics.ProgStart()
	input := getInputOrDie()
	metrics.InputLen(len(input))

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func problem1(input []*GeoLine) {

}

func problem2(input []*GeoLine) {

}

func getInputOrDie() []*GeoLine {
	lines, err := io.FromFile(inputFilePath, true)
	if err != nil {
		log.Fatalf("input error | %v", err)
	}
	input := make([]*GeoLine, len(lines))
	var tokens []string
	var gl *GeoLine
	for i, line := range lines {
		tokens = strings.Split(line, "->")
		gl = new(GeoLine)
		gl.x1, gl.y1 = parseCoordinates(tokens[0])
		gl.x2, gl.y2 = parseCoordinates(tokens[1])
		input[i] = gl
	}
	return input
}

func parseCoordinates(str string) (int, int) {
	tokens := strings.Split(str, ",")
	x, _ := strconv.Atoi(tokens[0])
	y, _ := strconv.Atoi(tokens[1])
	return x, y
}

type GeoLine struct {
	x1 int
	y1 int
	x2 int
	y2 int
}
