package main

func getAutoPilotInput() Direction {
	head := gSnake[0]
	var neck *Point
	if len(gSnake) > 1 {
		neck = gSnake[1]
	}
	var direc Direction
	if gFruit.X > head.X { // fruit ahead of snake
		direc = getFirstPossibleDirection([]Direction{Right, Up, Down, Left}, head, neck)
		return direc
	} else { // fruit behind (or same X) snake
		if gFruit.Y > head.Y { // fruit below snake
			direc = getFirstPossibleDirection([]Direction{Down, Left, Right, Up}, head, neck)
			return direc
		} else if gFruit.Y < head.Y { // fruit above snake
			direc = getFirstPossibleDirection([]Direction{Up, Left, Right, Down}, head, neck)
			return direc
		} else { // fruit Y same as snake
			direc = getFirstPossibleDirection([]Direction{Left, Up, Down, Right}, head, neck)
			return direc
		}
	}
}

func getFirstPossibleDirection(direcList []Direction, head, next *Point) Direction {
	for _, d := range direcList {
		if canMoveInDirection(d, head, next) {
			return d
		}
	}
	return None
}

func canMoveInDirection(direc Direction, head, next *Point) bool {
	var newHead *Point
	switch direc {
	case Up:
		newHead = &Point{X: head.X, Y: head.Y - 1}
		break
	case Down:
		newHead = &Point{X: head.X, Y: head.Y + 1}
		break
	case Left:
		newHead = &Point{X: head.X - 1, Y: head.Y}
		break
	case Right:
		newHead = &Point{X: head.X + 1, Y: head.Y}
		break
	}
	return isPosValid(newHead) && !isSameCoord(newHead, next)
}

func isPosValid(p *Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < BoardX && p.Y < BoardY
}

func isSameCoord(a, b *Point) bool {
	if a == nil || b == nil {
		return false
	}
	return a.X == b.X && a.Y == b.Y
}
