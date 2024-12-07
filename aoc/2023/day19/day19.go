package day19

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"strings"
)

var gInputWf map[string]*workflow
var gInputParts []*part

const DirPath = "../2023/day19"

func SolveP1() string {
	ans := 0
	accParts := make([]*part, 0)
	for _, p := range gInputParts {
		wf := gInputWf["in"]
		for {
			res := wf.eval(p)
			errz.HardAssert(len(res) > 0, "invalid workflow eval : %v | w=(%v) | p=(%v)", res, wf.str(), p.str())
			if res == "A" {
				accParts = append(accParts, p)
				break
			} else if res == "R" {
				break
			} else {
				wf = gInputWf[res]
			}
		}
	}
	for _, p := range accParts {
		ans += p.m["s"] + p.m["a"] + p.m["m"] + p.m["x"]
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInputWf = make(map[string]*workflow)
	i := 0
	for ; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			break
		}
		wf := newWorkflow(lines[i])
		gInputWf[wf.id] = wf
	}
	gInputParts = make([]*part, 0)
	for i++; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			break
		}
		p := newPart(lines[i])
		gInputParts = append(gInputParts, p)
	}

	for _, wf := range gInputWf {
		fmt.Println(wf.str())
	}
	fmt.Println()
	for _, p := range gInputParts {
		fmt.Println(p.str())
	}
}

/***** Types *****/

type ruleFunc func(int, int) bool

type part struct {
	m map[string]int
}

func newPart(line string) *part {
	t := stringz.SplitMulti(line, []string{",", "=", "{", "}"})
	p := &part{m: make(map[string]int)}
	p.m["x"] = stringz.AtoI(t[2], -1)
	p.m["m"] = stringz.AtoI(t[4], -1)
	p.m["a"] = stringz.AtoI(t[6], -1)
	p.m["s"] = stringz.AtoI(t[8], -1)
	return p
}

func (p *part) str() string {
	return fmt.Sprintf("%v", p.m)
}

type workflow struct {
	id    string
	rules []*rule
}

func (w *workflow) str() string {
	str := fmt.Sprintf("%v : ", w.id)
	for _, r := range w.rules {
		str += fmt.Sprintf("%v | ", r.str())
	}
	return str
}

func newWorkflow(line string) *workflow {
	t := stringz.SplitMulti(line, []string{"{", "}", ","})
	id := t[0]
	rules := make([]*rule, 0)
	for _, s := range t[1:] {
		if len(s) > 0 {
			rules = append(rules, newRule(s))
		}
	}
	return &workflow{id: id, rules: rules}
}

func (w *workflow) eval(p *part) string {
	for _, r := range w.rules {
		if r.eval(p) {
			return r.re
		}
	}
	return ""
}

type rule struct {
	op string
	cm int
	rf ruleFunc
	re string
}

func (r *rule) str() string {
	return fmt.Sprintf("%v %v %v", r.op, r.re, r.rf)
}

func (r *rule) eval(p *part) bool {
	val, _ := p.m[r.op]
	return r.rf(val, r.cm)
}

func newRule(line string) *rule {
	if strings.Contains(line, ":") {
		t := stringz.SplitMulti(line, []string{">", "<", ":"})
		if strings.Contains(line, ">") {
			return &rule{op: t[0], cm: stringz.AtoI(t[1], -1), rf: gt, re: t[2]}
		} else if strings.Contains(line, "<") {
			return &rule{op: t[0], cm: stringz.AtoI(t[1], -1), rf: lt, re: t[2]}
		}
		errz.HardAssert(false, "invalid rule : %v", line)
		return nil
	} else {
		return &rule{op: "", cm: 0, rf: tr, re: line}
	}
}

func gt(i, n int) bool {
	return i > n
}

func lt(i, n int) bool {
	return i < n
}

func tr(i, j int) bool {
	return true
}
