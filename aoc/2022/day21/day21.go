package day21

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"math"
	"strings"
)

var gInput []string

func SolveP1() string {
	jobs := parseInDS()
	mname := "root"
	resolveP1(jobs, mname)
	ans := jobs[mname].val
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	jobs := parseInDS()
	mname, res := getChildWithHumanDep(jobs)
	//fmt.Printf("%v : %v\n", mname, res)
	jobs[mname].res = res
	resolveOppP2(jobs, mname)
	ans := jobs["humn"].val
	return fmt.Sprintf("%v", ans)
}

/***** P2 Functions *****/

func resolveOppP2(jobs map[string]*job, mname string) {
	job := jobs[mname]
	//fmt.Printf("resolve-opp-p1 : %v %v : %v\n", mname, job.res, job.str())
	if strings.Compare(mname, "humn") == 0 {
		job.val = job.res
		job.isVal = true
		job.isDone = true
		return
	}
	if job.isVal || job.isDone {
		return
	}
	depL, depR := job.deps[0], job.deps[1]
	var err error
	if err = resolveP2(jobs, depR); err == nil {
		jobs[depL].res = job.ops[2](jobs[depR].val, job.res)
		resolveOppP2(jobs, depL)
	} else if err = resolveP2(jobs, depL); err == nil {
		jobs[depR].res = job.ops[1](jobs[depL].val, job.res)
		resolveOppP2(jobs, depR)
	}
}

func getChildWithHumanDep(jobs map[string]*job) (string, int64) {
	mname1, mname2 := jobs["root"].deps[0], jobs["root"].deps[1]
	mname := ""
	res := int64(0)
	var err error
	if err = resolveP2(jobs, mname1); err == nil {
		mname = mname2
		res = jobs[mname1].val
	} else if err = resolveP2(jobs, mname2); err == nil {
		mname = mname1
		res = jobs[mname2].val
	}
	return mname, res
}

func resolveP2(jobs map[string]*job, mname string) error {
	if strings.Compare(mname, "humn") == 0 {
		return fmt.Errorf("humn detected")
	}
	job := jobs[mname]
	if job.isVal || job.isDone {
		return nil
	}
	dep0, dep1 := job.deps[0], job.deps[1]
	if err := resolveP2(jobs, dep0); err != nil {
		return err
	}
	if err := resolveP2(jobs, dep1); err != nil {
		return err
	}
	job.val = job.ops[0](jobs[dep0].val, jobs[dep1].val)
	job.isDone = true
	return nil
}

/***** P1 Functions *****/

func resolveP1(jobs map[string]*job, mname string) {
	job := jobs[mname]
	if job.isVal || job.isDone {
		return
	}
	dep0, dep1 := job.deps[0], job.deps[1]
	resolveP1(jobs, dep0)
	resolveP1(jobs, dep1)
	job.val = job.ops[0](jobs[dep0].val, jobs[dep1].val)
	job.isDone = true
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = lines
}

func parseInDS() map[string]*job {
	jobs := make(map[string]*job)
	var t []string
	for _, l := range gInput {
		t = stringz.SplitMulti(l, []string{":", " "})
		if len(t) == 3 {
			jobs[t[0]] = newJob(t[2], []string{})
		} else {
			jobs[t[0]] = newJob(t[2], []string{t[2], t[3], t[4]})
		}
	}
	return jobs
}

// frjs: hplr + ldrt
type job struct {
	val    int64
	res    int64
	deps   []string
	ops    []func(int64, int64) int64
	isDone bool
	isVal  bool // can be removed, used for debugging
}

func (j *job) str() string {
	if j.isVal {
		return fmt.Sprintf("%v", j.val)
	} else if j.isDone {
		return fmt.Sprintf("%v, %v, %v, %v, %v", j.val, j.deps, j.res, j.isDone, j.ops)
	} else {
		return fmt.Sprintf("%v, %v, %v, %v", j.deps, j.res, j.isDone, j.ops)
	}
}

func newJob(val string, tokens []string) *job {
	if len(tokens) == 0 {
		return &job{
			val:    int64(stringz.AtoI(val, math.MaxInt)),
			res:    0,
			deps:   nil,
			ops:    nil,
			isDone: true,
			isVal:  true,
		}
	} else {
		return &job{
			val:    math.MaxInt,
			res:    0,
			deps:   []string{tokens[0], tokens[2]},
			ops:    getOps(tokens[1]),
			isDone: false,
			isVal:  false,
		}
	}
}

func getOps(o string) []func(int64, int64) int64 {
	ops := make([]func(int64, int64) int64, 3)
	switch o {
	case "+":
		ops[0] = func(a, b int64) int64 { // op
			return a + b
		}
		ops[1] = func(a, res int64) int64 { // reverse op for operand-1
			return res - a
		}
		ops[2] = func(b, res int64) int64 { // reverse op for operand-2
			return res - b
		}
		return ops
	case "-":
		ops[0] = func(a, b int64) int64 { // op
			return a - b
		}
		ops[1] = func(a, res int64) int64 { // reverse op for operand-1
			return a - res
		}
		ops[2] = func(b, res int64) int64 { // reverse op for operand-2
			return res + b
		}
		return ops
	case "*":
		ops[0] = func(a, b int64) int64 { // op
			return a * b
		}
		ops[1] = func(a, res int64) int64 { // reverse op for operand-1
			return res / a
		}
		ops[2] = func(b, res int64) int64 { // reverse op for operand-2
			return res / b
		}
		return ops
	case "/":
		ops[0] = func(a, b int64) int64 { // op
			return a / b
		}
		ops[1] = func(a, res int64) int64 { // reverse op for operand-1
			return a / res
		}
		ops[2] = func(b, res int64) int64 { // reverse op for operand-2
			return b * res
		}
		return ops
	}
	errz.HardAssert(false, "invalid op %v", o)
	return nil
}
