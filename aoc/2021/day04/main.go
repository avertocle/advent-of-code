package main

import (
	"fmt"
	"github.com/avertocle/contests/io/input"
	"log"
	"strconv"
	"strings"

	"github.com/avertocle/contests/metrics"
)

const inputFilePath = "input.txt"
const boardSize = 5

func main() {
	metrics.ProgStart()
	boards, draws := getInputOrDie()
	metrics.InputLen(len(boards))

	fmt.Printf("draws = %v\n\n", draws)
	fmt.Printf("boards(%v), board-size(%vx%v)\n\n", len(boards), len(boards[0]), len(boards[0][0]))
	//input.PrettyArray3D(boards)

	// sum, draw := problem1(boards, draws)
	// fmt.Printf("score(%v) sum(%v) draw(%v) \n", sum*draw, sum, draw)

	sum, draw := problem2(boards, draws)
	fmt.Printf("score(%v) sum(%v) draw(%v) \n", sum*draw, sum, draw)

	//input.PrettyArray3D(boards)

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func problem1(boards [][][]int, draws []int) (int, int) {
	for _, draw := range draws {
		for _, board := range boards {
			if findAndMarkInBoardAndCheckBingo(board, draw) {
				return sumUnmarkedNumbers(board), draw
			}
		}
	}
	return 0, 0
}

func problem2(boards [][][]int, draws []int) (int, int) {
	wonBoards := make(map[int]bool)
	lastWonBoard := -1
	lastDraw := -1
	for _, draw := range draws {
		for i, board := range boards {
			if _, ok := wonBoards[i]; !ok {
				if findAndMarkInBoardAndCheckBingo(board, draw) {
					wonBoards[i] = true
					lastWonBoard = i
					lastDraw = draw
					fmt.Printf("board won : %v\n", i)
					if len(wonBoards) == len(boards) {
						return sumUnmarkedNumbers(boards[lastWonBoard]), lastDraw
					}
				}
			}
		}
	}
	return 0, 0
}

func findAndMarkInBoardAndCheckBingo(board [][]int, draw int) bool {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == draw {
				board[i][j] = -1
				if checkBoardComplete(board, i, j) {
					return true
				}
			}
		}
	}
	return false
}

func checkBoardComplete(board [][]int, recentX, recentY int) bool {
	rowComplete := true
	for i := 0; i < boardSize; i++ {
		if board[recentX][i] != -1 {
			rowComplete = false
			break
		}
	}
	colComplete := true
	for i := 0; i < boardSize; i++ {
		if board[i][recentY] != -1 {
			colComplete = false
			break
		}
	}
	return rowComplete || colComplete
}

func sumUnmarkedNumbers(board [][]int) int {
	sum := 0
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] != -1 {
				sum += board[i][j]
			}
		}
	}
	return sum
}

func getInputOrDie() ([][][]int, []int) {
	lines, err := input.FromFile(inputFilePath, true)
	if err != nil {
		log.Fatalf("input error | %v", err)
	}
	inputLen := len(lines)
	boards := make([][][]int, (inputLen-1)/boardSize)
	draws := SplitToIntArray(lines[0], ",")
	boardNo := 0
	boardRowNo := 0
	for i := 1; i < inputLen; i++ {
		x := SplitToIntArray(lines[i], " ")
		boardNo = (i - 1) / boardSize
		boardRowNo = (i - 1) % boardSize
		if len(boards[boardNo]) == 0 {
			boards[boardNo] = make([][]int, 5)
		}
		boards[boardNo][boardRowNo] = x
	}
	return boards, draws
}

func SplitToIntArray(line string, sep string) []int {
	tokens := strings.Split(line, sep)
	if sep == " " {
		tokens = strings.Fields(line)
	}
	ans := make([]int, len(tokens))
	var err error
	for i, t := range tokens {
		ans[i], err = strconv.Atoi(t)
		if err != nil {
			fmt.Printf("strconv.Atoi failed for (%v) (%v) | %v", i, t, err)
			fmt.Println(strings.Join(tokens, "|"))
		}
	}
	return ans
}
