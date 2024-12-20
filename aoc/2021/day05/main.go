package main

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/numz"
	"log"
	"strconv"
	"strings"

	"github.com/avertocle/contests/metrics"
)

const inputFilePath = "input.txt"
const maxIntersections = 2

func main() {
	metrics.ProgStart()
	input := getInputOrDie()
	metrics.InputLen(len(input))

	ans := problem1(input)
	fmt.Printf("ans = %v\n", ans)

	// ans := problem2(iutils)
	// fmt.Printf("ans = %v\n", ans)

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func problem1(input []*GeoLine) int {
	dim := getMaxCoord(input) + 1
	grid := bytez.Init2D(dim, dim, 0)
	for _, gl := range input {
		if gl.isHori() {
			markGridHori(grid, gl.x1, gl.x2, gl.y1)
		} else if gl.isVert() {
			markGridVert(grid, gl.y1, gl.y2, gl.x1)
		} else {
		}
	}
	ans := bytez.CountIf2D(grid, func(val byte, x, y int) bool {
		return val >= maxIntersections
	})
	return ans
}

func problem2(input []*GeoLine) int {
	dim := getMaxCoord(input) + 1
	grid := bytez.Init2D(dim, dim, 0)
	for _, gl := range input {
		if gl.isHori() {
			markGridHori(grid, gl.x1, gl.x2, gl.y1)
		} else if gl.isVert() {
			markGridVert(grid, gl.y1, gl.y2, gl.x1)
		} else {
			markGridDiag(grid, gl.x1, gl.y1, gl.x2, gl.y2)
		}
	}
	ans := bytez.CountIf2D(grid, func(val byte, x, y int) bool {
		return val >= maxIntersections
	})
	return ans
}

func markGridHori(grid [][]byte, x1, x2, y int) {
	for i := numz.Min(x1, x2); i <= numz.Max(x1, x2); i++ {
		grid[i][y]++
	}
}

func markGridVert(grid [][]byte, y1, y2, x int) {
	for i := numz.Min(y1, y2); i <= numz.Max(y1, y2); i++ {
		grid[x][i]++
	}
}

func markGridDiag(grid [][]byte, x1, y1, x2, y2 int) {
	// y = mx + c
	m := (y2 - y1) / (x2 - x1)
	c := (y1 - (m * x1))
	var x, y int
	for x = numz.Min(x1, x2); x <= numz.Max(x1, x2); x++ {
		y = (m * x) + c
		grid[x][y]++
	}
}

func getMaxCoord(input []*GeoLine) int {
	maxCoord := 0
	for _, gl := range input {
		maxCoord = numz.Max(maxCoord, numz.Max(numz.Max(gl.x1, gl.x2), numz.Max(gl.x1, gl.x2)))
	}
	return maxCoord
}

func getInputOrDie() []*GeoLine {
	lines, err := iutils.FromFile(inputFilePath, true)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	input := make([]*GeoLine, len(lines))
	var tokens []string
	var gl *GeoLine
	for i, line := range lines {
		tokens = strings.Split(line, "->")
		gl = new(GeoLine)
		gl.x1, gl.y1 = parseCoordinates(strings.TrimSpace(tokens[0]))
		gl.x2, gl.y2 = parseCoordinates(strings.TrimSpace(tokens[1]))
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

func (this *GeoLine) isHori() bool {
	return this.y1 == this.y2
}

func (this *GeoLine) isVert() bool {
	return this.x1 == this.x2
}

func (this *GeoLine) s() string {
	return fmt.Sprintf("%v,%v  to %v,%v", this.x1, this.y1, this.x2, this.y2)
}

func printGrid(input []*GeoLine) {
	for _, this := range input {
		fmt.Printf("%v,%v  -> %v,%v\n", this.x1, this.y1, this.x2, this.y2)
	}

}
