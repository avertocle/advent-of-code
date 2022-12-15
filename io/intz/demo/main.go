package main

import (
	"fmt"
	"github.com/avertocle/contests/io/intz"
	"math"
	"time"
)

func main() {
	benchmark01(14)
}

func benchmark01(maxDigits int64) {
	var max, digits, t int64
	for digits = 1; digits <= maxDigits; digits++ {
		t = time.Now().Unix()
		max = int64(math.Pow(10, float64(digits)))
		for i := int64(0); i < max; i++ {
		}
		fmt.Printf("time taken : digits(%02d) (%04d) (%v)\n", digits, time.Now().Unix()-t, max-1)
	}
}

func demoIncBounded() {
	x := 0
	for i := 0; i < 100; i++ {
		x = intz.IncBounded(x, 21, 10)
		fmt.Printf("%v-", x)
	}
}
