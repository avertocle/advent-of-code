package day16

import (
	"fmt"
	"github.com/avertocle/contests/io/baze"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"log"
	"strings"
)

var gInput string
var gInpBin []byte
var gInpLen int

func SolveP1() string {
	var pktVer, pktTid int
	var litEnd, litDec int
	var subPktStart, subPktLen, subPktCnt int
	var lenTid byte
	ptr := 0
	versions := make([]int, 0)
	depth := 0
	for {
		if ptr+6 >= gInpLen {
			break
		}
		pktVer, pktTid = parsePktHeader(gInpBin[ptr : ptr+6])
		ptr += 6
		if pktVer == 0 {
			break
		} else {
			versions = append(versions, pktVer)
		}
		if pktTid == 4 {
			litEnd, litDec = parseLitPkt(ptr)
			ptr = litEnd
		} else {
			subPktStart, lenTid, subPktLen, subPktCnt = parseOprPkt(ptr)
			ptr = subPktStart
			depth++
		}
		fmt.Printf("%v - v(%v) t(%v)\n", strings.Repeat(" ", 2*depth), pktVer, pktTid)
	}
	fmt.Printf(">>>>>> ignore %v,%v,%v\n", lenTid, subPktLen, subPktCnt)
	//fmt.Printf("%v\n", string(gInpBin))
	fmt.Printf(">>>>>> ignore %v, %v, %v, %v\n", pktVer, pktTid, litEnd, litDec)
	ans := intz.Sum1D(versions)
	fmt.Printf("%v = sum(%v)\n", ans, versions)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

//110-100-10111-11110-00101-000

/***** Common Functions *****/

func parsePktHeader(vbits []byte) (int, int) {
	pktVer := baze.BTod32Q(vbits[:3], -1)
	pktTid := baze.BTod32Q(vbits[3:], -1)
	return pktVer, pktTid
}

// return end of literal value marker and the literal value
func parseLitPkt(start int) (int, int) {
	end := start
	litBin := make([]byte, 0)
	for i := start; i < gInpLen; i += 5 {
		litBin = append(litBin, gInpBin[i+1:i+5]...)
		if gInpBin[i] == '0' {
			end = i + 5
			break
		}
	}
	litDec := baze.BTod32Q(litBin, -1)
	return end, litDec
}

// return : subPktStart, lenTid, subPktLen, subPktCnt
func parseOprPkt(start int) (int, byte, int, int) {
	var lenTid byte
	var ptr, subPktLen, subPktCnt, subPktStart int
	ptr = start
	lenTid = gInpBin[ptr]
	ptr++
	if lenTid == '0' {
		subPktLen = baze.BTod32Q(gInpBin[ptr:ptr+15+1], -1)
		ptr += 15
	} else {
		subPktCnt = baze.BTod32Q(gInpBin[ptr:ptr+11+1], -1)
		ptr += 11
	}
	subPktStart = ptr
	return subPktStart, lenTid, subPktLen, subPktCnt
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	gInput = lines[0]
	gInpBin = baze.HtobQ(gInput)
	gInpLen = len(gInpBin)
}
