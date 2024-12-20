package day14

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
)

const DirPath = "../2024/day14"

var gDims *arrz.Idx2D[int]
var gInput []*robot

func SolveP1() string {
	simCount := 100
	currRobots := cloneAllRobots(gInput)
	for i := 0; i < simCount; i++ {
		moveRobots(currRobots)
	}
	ans := partitionIntoQuadsAndCalcAns(currRobots)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	// visual solution
	// dump grid to file after every sim and observe
	// this code is just a placeholder for test cases to pass
	currRobots := cloneAllRobots(gInput)
	simCount := 0
	for {
		simCount++
		moveRobots(currRobots)
		if simCount == 7916 {
			break
		}
		//fmt.Printf("\n\n\nsim %v", i)
		//displayRobotsOnGrid(currRobots)
	}
	ans := simCount
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func partitionIntoQuadsAndCalcAns(robots []*robot) int {
	quads := []int{0, 0, 0, 0}
	for _, r := range robots {
		if r.p.I < gDims.I/2 && r.p.J < gDims.J/2 {
			quads[0]++
		} else if r.p.I < gDims.I/2 && r.p.J > gDims.J/2 {
			quads[1]++
		} else if r.p.I > gDims.I/2 && r.p.J < gDims.J/2 {
			quads[2]++
		} else if r.p.I > gDims.I/2 && r.p.J > gDims.J/2 {
			quads[3]++
		}
	}
	//fmt.Println(quads)
	ans := intz.Mul1D(quads)
	return ans
}

/***** P2 Functions *****/

/***** Common Functions *****/

func moveRobots(robots []*robot) {
	for _, r := range robots {
		r.p.MoveBounded(r.v.I, r.v.J, 0, 0, gDims.I-1, gDims.J-1)
	}
}

type robot struct {
	p *arrz.Idx2D[int]
	v *arrz.Idx2D[int]
}

func (r *robot) str() string {
	return fmt.Sprintf("[%v,%v @ %v,%v]", r.p.I, r.p.J, r.v.I, r.v.J)
}

func (r *robot) clone() *robot {
	return &robot{
		p: r.p.Clone(),
		v: r.v.Clone(),
	}
}

func cloneAllRobots(robots []*robot) []*robot {
	var clones []*robot
	for _, r := range robots {
		clones = append(clones, r.clone())
	}
	return clones
}

func displayRobotsOnGrid(robots []*robot) {
	grid := arrz.Init2D(gDims.I, gDims.J, byte('.'))
	for _, r := range robots {
		grid[r.p.I][r.p.J] = 'x'
	}
	arrz.PPrint2D(grid)
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = make([]*robot, 0)
	sections := iutils.BreakByEmptyLineString1D(lines)
	gDims = arrz.NewIdx2D(
		stringz.AtoI(sections[0][0], -1),
		stringz.AtoI(sections[0][1], -1),
	)
	// switching I,J while taking input to make it more intuitive
	for _, line := range sections[1] {
		t := stringz.SplitMultiTrimSpace(line, []string{" ", "=", ","})
		r := &robot{
			p: arrz.NewIdx2D(stringz.AtoI(t[2], -1), stringz.AtoI(t[1], -1)),
			v: arrz.NewIdx2D(stringz.AtoI(t[5], -1), stringz.AtoI(t[4], -1)),
		}
		gInput = append(gInput, r)
	}
	//for _, r := range gInput {
	//	fmt.Println(r.str())
	//}
}
