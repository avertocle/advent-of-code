package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
	"strings"
	"time"
)

var gBoard [][]byte
var gFruit *Point
var gSnake []*Point
var gScore int
var gLog []string

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	initialiseGame()
	runForever := true
	for runForever {
		gameLoop()
		if shouldReplay() {
			initialiseGame()
		} else {
			runForever = false
		}
	}
}

func initialiseGame() {
	gBoard = make([][]byte, BoardY)
	makeEmptyBoard()
	gSnake = []*Point{generateRandomPosOnBoard()}
	gFruit = generateRandomPosOnBoard()
	gScore = 0
}

func gameLoop() {
	isGameActive := true
	currDirec := 0
	ctr := 0
	for isGameActive {
		if newDirec := getUserInput(); newDirec != 0 {
			currDirec = newDirec
		}
		updateBoard(currDirec)
		isGameActive = !checkSnakeDead()
		ctr++
		renderGame()
		time.Sleep(100 * time.Millisecond)
	}
}

func getUserInput() int {
	return int(getAutoPilotInput())
}

func updateBoard(direc int) {
	updateGameElements(direc)
	makeEmptyBoard()
	addSnakeToBoard()
	addFruitToBoard()
}

func updateGameElements(direc int) {
	if direc < 1 || direc > 4 {
		fmt.Printf("debug : invalid direction : %v\n", direc)
		return
	}
	newSnakeStart := getUpdatedSnakeStart(direc)
	Logf("%v", newSnakeStart)
	didSnakeEatFruit := checkSnakeEatFruit(newSnakeStart)
	gSnake = makeUpdatedSnake(newSnakeStart, didSnakeEatFruit)
	if didSnakeEatFruit {
		gFruit = generateRandomPosOnBoard()
		gScore += 10
	}
}

func makeEmptyBoard() {
	for i := 0; i < BoardY; i++ {
		gBoard[i] = make([]byte, BoardX)
	}
}

func addSnakeToBoard() {
	if len(gSnake) == 0 {
		return
	}
	addPointToBoard(gSnake[0], CharSnakeHead)
	for i := 1; i < len(gSnake); i++ {
		addPointToBoard(gSnake[i], CharSnakeBody)
	}
}

func addFruitToBoard() {
	if gFruit == nil {
		return
	}
	addPointToBoard(gFruit, CharFruit)
}

func shouldReplay() bool {
	fmt.Printf("\nPress %v to play again, any other key to exit.\n", StrReplay)
	var replay string
	fmt.Scanf("%s", &replay)
	return strings.Compare(replay, StrReplay) == 0
}

/*******************************************************
snake methods
*******************************************************/

func checkSnakeEatFruit(snakeStart *Point) bool {
	return snakeStart.X == gFruit.X && snakeStart.Y == gFruit.Y
}

func checkSnakeDead() bool {
	return false
}

func getUpdatedSnakeStart(currDir int) *Point {
	snakeStart := new(Point)
	if currDir == 1 {
		snakeStart.X = gSnake[0].X
		snakeStart.Y = gSnake[0].Y - 1
	} else if currDir == 2 {
		snakeStart.X = gSnake[0].X
		snakeStart.Y = gSnake[0].Y + 1
	} else if currDir == 3 {
		snakeStart.X = gSnake[0].X - 1
		snakeStart.Y = gSnake[0].Y
	} else if currDir == 4 {
		snakeStart.X = gSnake[0].X + 1
		snakeStart.Y = gSnake[0].Y
	}
	return snakeStart
}

func makeUpdatedSnake(snakeStart *Point, incSize bool) []*Point {
	newSnake := make([]*Point, 1)
	newSnake[0] = snakeStart
	for i := 0; i < len(gSnake)-1; i++ {
		newSnake = append(newSnake, gSnake[i])
	}
	if incSize {
		newSnake = append(newSnake, gSnake[len(gSnake)-1])
	}
	return newSnake
}

/*******************************************************
Utils
*******************************************************/

func addPointToBoard(p *Point, val byte) {
	gBoard[p.Y][p.X] = val
}

func generateRandomPosOnBoard() *Point {
	p := &Point{
		X: rand.Int() % BoardX,
		Y: rand.Int() % BoardY,
	}
	return p
}

type Point struct {
	X, Y int
}

func Logf(format string, a ...any) {
	gLog = append(gLog, fmt.Sprintf(format, a...))
}
