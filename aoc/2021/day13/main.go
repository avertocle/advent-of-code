package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/avertocle/contests/io"
	"github.com/avertocle/contests/metrics"
)

const inputFilePath = "input.txt"

var in *input

type input struct {
	grid [][]int
	axes []int
	rows int
	cols int
}

func main() {
	metrics.ProgStart()
	in = getInputOrDie()
	metrics.InputLen(in.rows)

	//io.PrettyArray2DInt(in.grid)
	fmt.Printf("%v\n", in.axes)

	ans := problem1()
	//ans := problem2()
	fmt.Printf("ans = %v\n", ans)

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func getInputOrDie() *input {
	lines, err := io.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("input error | %v", err)
	}

	var grid [][]int
	var axes []int
	flag := false
	for _, l := range lines {
		if !flag {
			if len(l) == 0 {
				flag = true
				continue
			} else {
				grid = append(grid, io.StringToInt1D(l, ","))
			}
		} else {
			axes = append(axes, parseFolds(l))
		}
	}
	return &input{
		grid: grid,
		axes: axes,
		rows: len(grid),
		cols: len(grid[0]),
	}
}

func parseFolds(line string) int {
	tokens := strings.Split(strings.Fields(line)[2], "=")
	v, _ := strconv.Atoi(tokens[1])
	if tokens[0] == "x" {
		return v
	} else {
		return (-1 * v)
	}
}

/***** Logic Begins here *****/

func problem1() int {
	for _, axis := range in.axes {
		if axis > 0 {
			foldAlongY(axis)
		} else {
			foldAlongX(-1 * axis)
		}
	}
	showGrid(100, 100)
	return io.Unique1DIntIn2DInt(in.grid)
}

func foldAlongY(axis int) {
	newGrid := make([][]int, 0)
	for _, p := range in.grid {
		if p[0] > axis {
			newGrid = append(newGrid, []int{axis - (p[0] - axis), p[1]})
		} else {
			newGrid = append(newGrid, p)
		}
	}
	in.grid = newGrid
}

func foldAlongX(axis int) {
	newGrid := make([][]int, 0)
	for _, p := range in.grid {
		if p[1] > axis {
			newGrid = append(newGrid, []int{p[0], axis - (p[1] - axis)})
		} else {
			newGrid = append(newGrid, p)
		}
	}
	in.grid = newGrid
}

func showGrid(rows, cols int) {
	g := io.Init2DByte(rows, cols, '.')
	for _, p := range in.grid {
		g[p[1]][p[0]] = '#'
	}
	io.PrettyArray2DByte(g)
}
