package day21

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"math"
)

var gInput map[string]*job

func SolveP1() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = make(map[string]*job)
	var t []string
	for _, l := range lines {
		t = stringz.SplitMulti(l, []string{":", " "})
		if len(t) == 3 {
			gInput[t[0]] = newJob(t[2], []string{})
		} else {
			gInput[t[0]] = newJob(t[2], []string{t[2], t[3], t[4]})
		}
		fmt.Println(t[0], " : ", gInput[t[0]].str())
	}
}

//frjs: hplr + ldrt
type job struct {
	val    int64
	deps   []string
	op     func(int64, int64) int64
	isDone bool
	isVal  bool
}

func (j *job) str() string {
	if j.isVal {
		return fmt.Sprintf("%v", j.val)
	} else if j.isDone {
		return fmt.Sprintf("%v %v, %v, %v", j.val, j.deps, j.isDone, j.op)
	} else {
		return fmt.Sprintf("%v, %v, %v", j.deps, j.isDone, j.op)
	}
}

func newJob(val string, tokens []string) *job {
	if len(tokens) == 0 {
		return &job{
			val:    int64(stringz.AtoiQ(val, math.MaxInt)),
			deps:   nil,
			op:     nil,
			isDone: true,
			isVal:  true,
		}
	} else {
		return &job{
			val:    math.MaxInt,
			deps:   []string{tokens[0], tokens[2]},
			op:     getOp(tokens[1]),
			isDone: false,
			isVal:  false,
		}
	}
}

func getOp(o string) func(int64, int64) int64 {
	switch o {
	case "+":
		return func(a, b int64) int64 {
			return a * b
		}
	case "-":
		return func(a, b int64) int64 {
			return a - b
		}
	case "*":
		return func(a, b int64) int64 {
			return a * b
		}
	case "/":
		return func(a, b int64) int64 {
			return a / b
		}
	}
	errz.HardAssert(false, "invalid op %v", o)
	return nil
}

//fqwq: 10
//jwjf: ssjq - vtlh
