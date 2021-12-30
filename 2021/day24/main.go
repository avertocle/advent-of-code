package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/avertocle/adventofcode/io"
	"github.com/avertocle/adventofcode/metrics"
)

const inputFilePath = "input.txt"

var in *input

type input struct {
	cmds   [][]string
	rows   int
	params [][]int
}

func main() {
	metrics.ProgStart()
	in = getInputOrDie()
	metrics.InputLen(in.rows)

	//io.PrettyArray2DString(in.cmds)
	io.PrettyArray2DInt(in.params)

	ans := problem1()
	//ans := problem2()
	fmt.Printf("ans = %v\n", ans)

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func getInputOrDie() *input {
	lines, err := io.FromFile(inputFilePath, true)
	if err != nil {
		log.Fatalf("input error | %v", err)
	}
	fmt.Printf("input-len = %v\n", len(lines))
	rows := len(lines)
	cmds := make([][]string, rows)
	params := make([][]int, 14)
	var tmp int
	var tokens []string
	for i := 0; i < rows; i++ {
		tokens = strings.Fields(lines[i])
		cmds[i] = tokens
		if i%18 == 4 {
			tmp, _ = strconv.Atoi(tokens[2])
			params[i/18] = []int{tmp, 0, 0}
		} else if i%18 == 5 {
			tmp, _ = strconv.Atoi(tokens[2])
			params[i/18][1] = tmp
		} else if i%18 == 15 {
			tmp, _ = strconv.Atoi(tokens[2])
			params[i/18][2] = tmp
		}
	}
	return &input{
		cmds:   cmds,
		rows:   rows,
		params: params,
	}
}

/***** Logic Begins here *****/

const simCount = 40

func problem1() int64 {
	var max, i, w int64
	var z int
	max = 99999999999999
	for i = max; i > 0; i-- {
		if i%100000000 == 0 {
			fmt.Printf("%v-", i)
		}
		z = 0
		for j := 0; j < 14; j++ {
			w = max % (int64(math.Pow10(14 - j)))
			if w == 0 {
				continue
			}
			z = iterate1(int(w), z, j)
		}
		if z == 0 {
			return i
		}
	}
	return 0
}

func iterate1(w, z, j int) int {
	k := in.params[j]
	if w == (z%26)+k[1] {
		return (z / k[0])
	} else {
		return ((z * 26) / k[0]) + w + k[2]
	}
}

func problem2() int {
	return 0
}
