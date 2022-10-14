package main

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

var vWidth, vHeight int
var vTop, vBottom, vLeft, vRight, vMidX, vMidY int

func renderGame() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	initViewPort()
	drawBoardBoundary()
	drawBoard()
	termbox.Flush()
	drawHeader()
	drawBoard()
	drawFooter()
}

func initViewPort() {
	vWidth, vHeight = termbox.Size()
	vMidY = vHeight / 2
	vLeft = (vWidth - BoardX) / 2
	vRight = (vWidth + BoardX) / 2
	vTop = vMidY - (BoardY / 2)
	vBottom = vMidY + (BoardY / 2) + 1
}

func coordToView(x, y int) (int, int) {
	return x + vLeft, y + vTop
}

func drawBoard() {
	var vX, vY int
	var clrFg, clrBg termbox.Attribute
	for i := 0; i < BoardX; i++ {
		for j := 0; j < BoardY; j++ {
			vX, vY = coordToView(i, j)
			clrFg, clrBg = getColorsByChar(gBoard[j][i])
			termbox.SetCell(vX, vY, rune(gBoard[j][i]), clrFg, clrBg)
		}
	}
}

func drawBoardBoundary() {
	for i := vLeft - 1; i <= vRight+1; i++ {
		termbox.SetCell(i, vTop-1, CharBoundary, ClrBoundaryFg, ClrBoundaryBg)
		termbox.SetCell(i, vTop-2, CharBoundary, ClrBoundaryFg, ClrBoundaryBg)
		termbox.SetCell(i, vBottom+1, CharBoundary, ClrBoundaryFg, ClrBoundaryBg)
		termbox.SetCell(i, vBottom+2, CharBoundary, ClrBoundaryFg, ClrBoundaryBg)
	}
	for i := vTop - 2; i <= vBottom+2; i++ {
		termbox.SetCell(vLeft-1, i, CharBoundary, ClrBoundaryFg, ClrBoundaryBg)
		termbox.SetCell(vLeft-2, i, CharBoundary, ClrBoundaryFg, ClrBoundaryBg)
		termbox.SetCell(vLeft-3, i, CharBoundary, ClrBoundaryFg, ClrBoundaryBg)
		termbox.SetCell(vLeft-4, i, CharBoundary, ClrBoundaryFg, ClrBoundaryBg)
		termbox.SetCell(vRight+1, i, CharBoundary, ClrBoundaryFg, ClrBoundaryBg)
		termbox.SetCell(vRight+2, i, CharBoundary, ClrBoundaryFg, ClrBoundaryBg)
		termbox.SetCell(vRight+3, i, CharBoundary, ClrBoundaryFg, ClrBoundaryBg)
		termbox.SetCell(vRight+4, i, CharBoundary, ClrBoundaryFg, ClrBoundaryBg)
	}
}

func getColorsByChar(b byte) (termbox.Attribute, termbox.Attribute) {
	switch b {
	case CharSpace:
		return ClrBoardFg, ClrBoardBg
	case CharSnakeHead:
		return ClrSnakeFg, ClrSnakeBg
	case CharSnakeBody:
		return ClrSnakeFg, ClrSnakeBg
	case CharFruit:
		return ClrFruitFg, ClrFruitBg
	}
	return ClrDefaultFg, ClrDefaultBg
}

func drawLogger() {
	currLog := gLog[:10]
	for i, l := range currLog {
		tbprint(0, i, termbox.ColorGreen, termbox.ColorBlack, l)
	}
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func drawHeader() {
	//fmt.Printf("\n")
	//fmt.Printf("Game Starts\n")
	//fmt.Printf("How To Play\n")
	//fmt.Printf(" - Press W/A/S/D to move the Snake\n")
	//fmt.Printf(" - Every Second adds 1 point to the Score\n")
	//fmt.Printf(" - Eat fruit to earn 10x more points\n")
	//fmt.Printf("\n")
}

func drawFooter() {
	//	fmt.Printf("\nScore = %v", gScore)
}
