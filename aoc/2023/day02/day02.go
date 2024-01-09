package day02

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"strconv"
	"strings"
)

var gInput map[int][]*cubeCx

const DirPath = "../2023/day01"

func SolveP1() string {
	ans := 0
	refGame := newCubeCx(12, 13, 14)
	for gameId, game := range gInput {
		if refGame.isGameValid(game) {
			ans += gameId
		}
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	for _, game := range gInput {
		gameLowerBound := calcGameLowerBound(game)
		power := calcPower(gameLowerBound)
		ans += power
	}
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

type cubeCx struct {
	r int
	g int
	b int
}

func newCubeCx(r, g, b int) *cubeCx {
	return &cubeCx{
		r: r,
		g: g,
		b: b,
	}
}

func (o *cubeCx) isMoveValid(m *cubeCx) bool {
	return m.r <= o.r && m.g <= o.g && m.b <= o.b
}

func (o *cubeCx) isGameValid(moves []*cubeCx) bool {
	for _, move := range moves {
		if !o.isMoveValid(move) {
			return false
		}
	}
	return true
}

func (o *cubeCx) str() string {
	return fmt.Sprintf("[%v %v %v]", o.r, o.g, o.b)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

func calcGameLowerBound(moves []*cubeCx) *cubeCx {
	maxR, maxG, maxB := 0, 0, 0
	for _, move := range moves {
		if move.r > maxR {
			maxR = move.r
		}
		if move.g > maxG {
			maxG = move.g
		}
		if move.b > maxB {
			maxB = move.b
		}
	}
	return newCubeCx(maxR, maxG, maxB)
}

func calcPower(ccx *cubeCx) int {
	return ccx.r * ccx.g * ccx.b
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = make(map[int][]*cubeCx)
	for _, line := range lines {
		gameId, moves := parseGameLine(line)
		gInput[gameId] = moves
	}
	//printInput()
}

func parseMoveLine(line string) *cubeCx {
	r, g, b := 0, 0, 0
	draws := strings.Split(line, ", ")
	for _, token := range draws {
		draw := strings.Split(token, " ")
		if strings.Compare(draw[1], "red") == 0 {
			r, _ = strconv.Atoi(draw[0])
		} else if strings.Compare(draw[1], "green") == 0 {
			g, _ = strconv.Atoi(draw[0])
		} else if strings.Compare(draw[1], "blue") == 0 {
			b, _ = strconv.Atoi(draw[0])
		}
	}
	return newCubeCx(r, g, b)
}

func parseGameLine(line string) (int, []*cubeCx) {
	tokens := strings.Split(line, ": ")
	gameId, _ := strconv.Atoi(strings.Split(tokens[0], " ")[1])
	moveLines := strings.Split(tokens[1], "; ")
	moves := make([]*cubeCx, len(moveLines))
	for i, moveLine := range moveLines {
		moves[i] = parseMoveLine(moveLine)
	}
	return gameId, moves
}

func printInput() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("%v -> ", i)
		for _, move := range gInput[i] {
			fmt.Printf("%v | ", move.str())
		}
		fmt.Println()
	}
}
