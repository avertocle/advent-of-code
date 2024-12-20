package day07

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"math"
	"path"
	"strings"
)

// input file was wrong : removed cd .. from 436, 495

var input []string

func SolveP1() string {
	root := makeTreeFromRawInput()
	calcDirSizes(root)
	ans := new(int)
	traverseP1(root, ans)
	return fmt.Sprintf("%v", *ans)
}

func SolveP2() string {
	root := makeTreeFromRawInput()
	calcDirSizes(root)
	ans := new(int)
	*ans = math.MaxInt
	toDelMin := root.size - (70000000 - 30000000) // min space to be deleted
	traverseP2(root, toDelMin, ans)
	return fmt.Sprintf("%v", *ans)
}
func sumFSizes(m map[string]int) int {
	s := 0
	for _, v := range m {
		s += v
	}
	return s
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

func traverseP1(r *tnode, ans *int) {
	if r.size <= 100000 {
		*ans += r.size
	}
	for _, c := range r.children {
		traverseP1(c, ans)
	}
}

// find smallest dir of size > minSize
func traverseP2(r *tnode, toDelMin int, ans *int) {
	if r.size >= toDelMin && r.size < *ans {
		*ans = r.size
	}
	for _, c := range r.children {
		traverseP2(c, toDelMin, ans)
	}
}

func calcDirSizes(r *tnode) {
	r.size = sumFSizes(r.files)
	for _, c := range r.children {
		calcDirSizes(c)
		r.size += c.size
	}
}

func makeTreeFromRawInput() *tnode {
	root := NewTNode("/", "/", nil)
	curr := root
	for i := 0; i < len(input); {
		l := input[i]
		//fmt.Printf("parsing : %v at line %v\n", l, i+1)
		if isCmd(l) {
			if isLS(l) {
				i++
				for ; i < len(input) && !isCmd(input[i]); i++ {
					l = input[i]
					if isDir(l) {
						addSubDir(curr, extractDname(l))
					} else { // is file output
						fname, size := parseFile(l)
						addFile(curr, fname, size)
					}
				}
			} else { // cmd = cd
				dName := parseCD(l)
				if isRootDirCmd(dName) {
					curr = root
				} else if isPrevDirCmd(dName) {
					curr = curr.parent
				} else {
					curr = curr.children[dName] // assuming cd is always into valid dir and after ls
				}
				i++
			}
		} else {
			errz.HardAssert(false, "should not reach here : cmd (%v) at line %v", l, i)
		}
	}
	return root
}

func addFile(tn *tnode, fname string, size int) {
	tn.files[fname] = size
}

func addSubDir(tn *tnode, dname string) {
	if _, ok := tn.children[dname]; !ok {
		temp := NewTNode(dname, path.Join(tn.absPath, dname), tn)
		tn.children[dname] = temp
	}
}

func isCmd(line string) bool {
	return strings.HasPrefix(line, "$")
}

func isLS(line string) bool {
	return strings.Compare(line, "$ ls") == 0
}

func parseCD(line string) string {
	return strings.Split(line, " ")[2]
}

func isDir(line string) bool {
	return strings.HasPrefix(line, "dir")
}

func isRootDirCmd(dname string) bool {
	return strings.Compare(dname, "/") == 0
}

func isPrevDirCmd(dname string) bool {
	return strings.Compare(dname, "..") == 0
}

func extractDname(line string) string {
	return strings.Split(line, " ")[1]
}

func parseFile(line string) (string, int) {
	t := strings.Split(line, " ")
	return t[1], stringz.AtoI(t[0], -1)
}

func printFileSystem(r *tnode, depth int) {
	printWithDepth(fmt.Sprintf("%v (dir, size=%v", r.name, r.size), depth)
	for _, tn := range r.children {
		printFileSystem(tn, depth+1)
	}
	for fname, size := range r.files {
		printWithDepth(fmt.Sprintf("%v (file, size=%v)", fname, size), depth+1)
	}
}

func printWithDepth(s string, d int) {
	fmt.Printf("%v- %v\n", strings.Repeat(" ", 2*d), s)
}

/***** Structs *****/

type tnode struct {
	name     string
	absPath  string
	parent   *tnode
	children map[string]*tnode
	files    map[string]int
	size     int
}

func NewTNode(name string, absPath string, parent *tnode) *tnode {
	return &tnode{
		name:     name,
		absPath:  absPath,
		parent:   parent,
		children: make(map[string]*tnode),
		files:    make(map[string]int),
		size:     0,
	}
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		fmt.Printf("iutils error | %v", err)
	}
	input = lines
}
