package day11

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"sort"
	"strings"
)

var input []string

func SolveP1() string {
	monkeys := makeMonkeys()
	//printMonkeys(monkeys)
	runSim(monkeys, 20, applyReliefP0, 3)
	//printAllMonkeysItems(monkeys)
	ans := calcMonkeyBusi(monkeys)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	monkeys := makeMonkeys()
	//printMonkeys(monkeys)
	runSim(monkeys, 10000, applyReliefP1, getReductionKey(monkeys))
	//printAllMonkeysItems(monkeys)
	ans := calcMonkeyBusi(monkeys)
	return fmt.Sprintf("%v", ans)
}

func applyReliefP0(val, key int64) int64 {
	return val / key
}

func applyReliefP1(val, key int64) int64 {
	return val % key
}

func runSim(monkeys []*monkey, iterations int, reliefFunc func(int64, int64) int64, reliefKey int64) {
	newMkid := -1
	var item *item
	for r := 0; r < iterations; r++ {
		for _, m := range monkeys {
			for len(m.sitems) > 0 {
				newMkid = m.inspectOneItem(reliefFunc, reliefKey)
				item = m.trwItem()
				monkeys[newMkid].rcvItem(item)
			}
		}
	}
}

func makeMonkeys() []*monkey {
	monkeys := make([]*monkey, (len(input)+1)/7)
	for i := 0; i < len(input); i += 7 {
		monkeys[i/7] = newMonkey(input[i : i+7])
	}
	return monkeys
}

func calcMonkeyBusi(monkeys []*monkey) int {
	arr := make([]int, len(monkeys))
	for i, m := range monkeys {
		arr[i] = m.insCtr
	}
	sort.Ints(arr)
	return arr[len(arr)-1] * arr[len(arr)-2]
}

/***** Common Functions *****/

func getReductionKey(monkeys []*monkey) int64 {
	key := 1
	for _, m := range monkeys {
		key *= m.tst.x
	}
	return int64(key)
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	input = lines
}

/*************** Structs ***************/

type monkey struct {
	id     int
	sitems []*item
	op     *op
	tst    *tst
	insCtr int
}

func (this *monkey) inspectOneItem(reliefFunc func(int64, int64) int64, reliefKey int64) int {
	if len(this.sitems) == 0 {
		return -1
	}
	item := this.sitems[0]
	item.wlvl = this.op.eval(item.wlvl)
	item.wlvl = reliefFunc(item.wlvl, reliefKey)
	newMid := this.tst.eval(item.wlvl)
	this.insCtr++
	return newMid
}

func (this *monkey) rcvItem(item *item) {
	this.sitems = append(this.sitems, item)
}

func (this *monkey) trwItem() *item {
	item := this.sitems[0]
	this.sitems = this.sitems[1:]
	return item
}

func (this *monkey) s() string {
	itemsStr := ""
	for _, it := range this.sitems {
		itemsStr += it.s() + " "
	}
	return fmt.Sprintf("%v : tst(%v) : sitems(%v) : ", this.id, this.tst.s(), itemsStr)
}

func newMonkey(lines []string) *monkey {
	id := stringz.AtoI(strings.Split(strings.Fields(lines[0])[1], ":")[0], -1)

	items := make([]*item, 0)
	itemStr := strings.TrimSpace(strings.Split(lines[1], ":")[1])
	tokens := strings.Split(itemStr, ",")
	for _, t := range tokens {
		items = append(items, newItem(stringz.AtoI(t, -1)))
	}

	op := newOp(strings.TrimSpace(strings.Split(lines[2], "=")[1]))
	tst := newTest(lines[3:6])
	m := &monkey{
		id:     id,
		sitems: items,
		op:     op,
		tst:    tst,
	}
	return m
}

/************************/

type op struct {
	x int
	f func(int64, int) int64
}

func (this *op) eval(old int64) int64 {
	return this.f(old, this.x)
}

func (this *op) s() string {
	return fmt.Sprintf("%v", this.x)
}

func newOp(s string) *op {
	tokens := strings.Split(s, " ")
	var x int
	var f func(int64, int) int64
	if strings.Compare(tokens[1], "+") == 0 {
		f = opAddToOld
		x = stringz.AtoI(tokens[2], -1)
	} else if strings.Compare(tokens[1], "*") == 0 {
		if strings.Compare(tokens[2], "old") == 0 {
			x = 0
			f = opSqrOld
		} else {
			x = stringz.AtoI(tokens[2], -1)
			f = opMulToOld
		}
	}
	return &op{
		x: x,
		f: f,
	}
}

/************************/

type tst struct {
	x     int
	mkTru int
	mkFal int
	f     func(int64, int) bool
}

func newTest(s []string) *tst {
	f := isDivBy
	x := stringz.AtoI(strings.Fields(s[0])[3], -1)
	mkTru := stringz.AtoI(strings.Fields(s[1])[5], -1)
	mkFal := stringz.AtoI(strings.Fields(s[2])[5], -1)
	return &tst{
		x:     x,
		f:     f,
		mkTru: mkTru,
		mkFal: mkFal,
	}
}

func (this *tst) s() string {
	return fmt.Sprintf("%v = %v or %v", this.x, this.mkTru, this.mkFal)
}

func (this *tst) eval(val int64) int {
	if this.f(val, this.x) {
		return this.mkTru
	} else {
		return this.mkFal
	}
}

/************************/

type item struct {
	id   int
	wlvl int64
}

var itemCtr int

func newItem(wlvl int) *item {
	itemCtr++
	return &item{
		id:   itemCtr,
		wlvl: int64(wlvl),
	}
}

func (this *item) s() string {
	return fmt.Sprintf("%v@%v", this.id, this.wlvl)
}

/***** Util Functions *****/

func opAddToOld(old int64, x int) int64 {
	return old + int64(x)
}

func opMulToOld(old int64, x int) int64 {
	return old * int64(x)
}

func opSqrOld(old int64, x int) int64 {
	return old * old
}

func isDivBy(val int64, d int) bool {
	return val%int64(d) == 0
}

func printMonkeys(monkeys []*monkey) {
	for _, m := range monkeys {
		fmt.Printf("%v\n", m.s())
	}
}

func printAllMonkeysItems(monkeys []*monkey) {
	var itemsStr string
	for _, m := range monkeys {
		itemsStr = ""
		for _, it := range m.sitems {
			itemsStr += (it.s() + " ")
		}
		fmt.Printf("m-%v : insCtr %v : %v \n", m.id, m.insCtr, itemsStr)
	}
}
