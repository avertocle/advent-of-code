package testz

import (
	"fmt"
	"strings"
	"testing"
)

func Execute(t *testing.T, testCases [][]string, parser func(string2 string), solvers []func() string) {
	for _, tc := range testCases {

		parser(tc[0])

		if strings.Compare(tc[1], "0") == 0 {
			fmt.Printf("%v : %v : unimplemented", tc[0], "P1")
		}
		ansP1 := solvers[0]()
		if strings.Compare(tc[1], "0") == 0 {
			fmt.Printf("%v : %v : unsolved", tc[0], "P1")
		}

		if strings.Compare(ansP1, tc[1]) != 0 {
			t.Errorf("failed : %v %v : expected (%v) got (%v)",
				tc[0], "P1", tc[1], ansP1)
		}

		ansP2 := solvers[1]()

		if strings.Compare(ansP2, tc[2]) != 0 {
			t.Errorf("failed : %v %v : expected (%v) got (%v)",
				tc[0], "P2", tc[2], ansP2)
		}
	}
}
