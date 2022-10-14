package main

import "github.com/nsf/termbox-go"

const BoardX = 60
const BoardY = 30

const CharBoundary = ' '
const CharSpace = 0
const CharFruit = '*'
const CharSnakeHead = 'O'
const CharSnakeBody = ' '
const StrReplay = "r"

const ClrBoundaryBg = termbox.ColorRed
const ClrBoundaryFg = termbox.ColorWhite
const ClrBoardBg = termbox.ColorBlack
const ClrBoardFg = termbox.ColorBlack
const ClrSnakeBg = termbox.ColorCyan
const ClrSnakeFg = termbox.ColorWhite
const ClrFruitBg = termbox.ColorGreen
const ClrFruitFg = termbox.ColorWhite
const ClrDefaultBg = termbox.ColorCyan
const ClrDefaultFg = termbox.ColorCyan

type Direction int

const (
	None Direction = iota
	Up
	Down
	Left
	Right
)
