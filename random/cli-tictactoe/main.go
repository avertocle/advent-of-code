package main

import "fmt"

const boardlen = 3

type Board [boardlen][boardlen]byte
type History [boardlen * boardlen]bool

func main() {

	board := initBoard()
	moveHistory := new(History)

	var currPlayer byte
	var currMove, moveCtr int
	isGameInProgress := true

	for moveCtr = 0; moveCtr < 9 && isGameInProgress; moveCtr++ {
		currPlayer = getCurrPlayer(moveCtr)
		displayBoard(board)
		currMove = getInput(currPlayer, moveHistory)
		moveHistory[currMove] = true
		moveX, moveY := convertMoveToIndex(currMove)
		board[moveX][moveY] = currPlayer
		isGameInProgress = !isGameOver(board, currMove, currPlayer)
	}
	processGameOver(board, currPlayer, moveCtr)
}

func processGameOver(board *Board, currPlayer byte, moveCtr int) {
	displayBoard(board)
	if moveCtr == 9 {
		fmt.Printf("Game Draw")
	} else {
		fmt.Printf("Game Over : Player-%c has won\n", currPlayer)
	}
}

func getCurrPlayer(moveCtr int) byte {
	if moveCtr%2 == 0 {
		return 'O'
	} else {
		return 'X'
	}
}

func getInput(currPlayer byte, moveHistory *History) int {
	var err error
	input := 0
	for true {
		if input, err = getInputFromConsole(currPlayer, moveHistory); err != nil {
			fmt.Printf("Invalid move. Please try again. | %v\n", err)
		} else {
			fmt.Printf("\n\n")
			break
		}
	}
	return input
}

func getInputFromConsole(playerId byte, moveHistory *History) (int, error) {
	fmt.Printf("Player-%c's turn. Enter any number from 1-9 to indicate your move : ", playerId)
	var input int
	if _, err := fmt.Scan(&input); err != nil {
		return 0, fmt.Errorf("error reading move | %v", err)
	} else if input < 1 || input > 9 {
		return 0, fmt.Errorf("move must be a number between 0 to 9")
	} else if moveHistory[input] {
		return 0, fmt.Errorf("cell is occupied")
	}
	return input, nil
}

func displayBoard(board *Board) {
	for i := 0; i < boardlen; i++ {
		for j := 0; j < boardlen; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}

func isGameOver(board *Board, recentMove int, recentChar byte) bool {
	movex, movey := convertMoveToIndex(recentMove)

	var allSame bool

	allSame = true
	for i := 0; i < boardlen; i++ {
		if board[movex][i] != recentChar {
			allSame = false
			break
		}
	}
	if allSame {
		return allSame
	}

	allSame = true
	for i := 0; i < boardlen; i++ {
		if board[i][movey] != recentChar {
			allSame = false
			break
		}
	}
	if allSame {
		return allSame
	}

	allSame = true
	for i := 0; i < boardlen; i++ {
		if board[i][i] != recentChar {
			allSame = false
			break
		}
	}
	if allSame {
		return allSame
	}

	allSame = true
	for i := 0; i < boardlen; i++ {
		if board[i][boardlen-i-1] != recentChar {
			allSame = false
			break
		}
	}
	if allSame {
		return allSame
	}

	return allSame
}

func convertMoveToIndex(move int) (x, y int) {
	x = (move - 1) / boardlen
	y = (move - 1) % boardlen
	return x, y
}

func convertMoveToIndexTest(move int) {
	for i := 1; i <= 9; i++ {
		x, y := convertMoveToIndex(i)
		fmt.Printf("test for %v = [%v, %v]\n", i, x, y)
	}
}

func initBoard() *Board {
	board := new(Board)
	for i := 0; i < boardlen; i++ {
		for j := 0; j < boardlen; j++ {
			board[i][j] = '-'
		}
	}
	return board
}
