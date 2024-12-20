package day25

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/numz"
	"math"
	"strconv"
	"strings"
)

// u5 => unbalanced 5
// u10 => unbalanced 10
// snafu = snafu

var gInput []string

func SolveP1() string {
	sumDec := int64(0)
	for _, fu := range gInput {
		sumDec += b5ToU10(snafuToB5(fu))
	}
	ans := b5ToSnafu(u10ToB5(sumDec))
	return fmt.Sprintf("%s", ans)
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
	gInput = lines
}

func b5ToU10(u5 []int) int64 {
	dec := int64(0)
	for i, fu := range u5 {
		dec += int64(fu) * numz.Pow(5, len(u5)-i-1)
	}
	return dec
}

func u10ToB5(u10 int64) []int {
	u5 := strings.Repeat("0", 10) + strconv.FormatInt(u10, 5)
	b5 := iutils.ExtractInt1DFromString0D(u5, "", math.MaxInt)
	for i := len(b5) - 1; i > 0; i-- {
		if b5[i] >= 3 {
			b5[i] -= 5
			b5[i-1] += 1
		}
	}
	return b5
}

func snafuToB5(snafu string) []int {
	b5 := make([]int, len(snafu))
	for i, fu := range snafu {
		if fu == '-' {
			b5[i] = -1
		} else if fu == '=' {
			b5[i] = -2
		} else if fu == '0' || fu == '1' || fu == '2' {
			b5[i] = int(fu - '0')
		} else {
			errz.HardAssert(false, "invalid snafu string (%v) idx(%v) byte(%v)", snafu, i, string(fu))
		}
	}
	return b5
}

func b5ToSnafu(b5 []int) string {
	snafu := make([]byte, 0)
	isPadding := true
	for i, x := range b5 {
		if x == 0 && isPadding == true {
			continue
		} else {
			isPadding = false
		}
		if x == -1 {
			snafu = append(snafu, '-')
		} else if x == -2 {
			snafu = append(snafu, '=')
		} else if x == 0 || x == 1 || x == 2 {
			snafu = append(snafu, byte('0'+x))
		} else {
			errz.HardAssert(false, "invalid snafu array (%v) idx(%v) byte(%v)", b5, i, b5[i])
		}
	}
	return string(snafu)
}
