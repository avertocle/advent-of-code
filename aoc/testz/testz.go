package testz

import (
	"fmt"
	"github.com/avertocle/contests/io/clr"
	"strings"
	"testing"
)

func Execute(t *testing.T, testCases [][]string, parser func(string2 string), solvers []func() string) {
	var res1, res2 string
	var err1, err2 error
	for _, tc := range testCases {
		res1, err1 = runOneTest(tc[0], "P1", tc[1], parser, solvers[0])
		res2, err2 = runOneTest(tc[0], "P2", tc[2], parser, solvers[1])
		t.Logf(res1)
		t.Logf(res2)
		if err1 != nil || err2 != nil {
			t.Fail()
		}
	}

}

func runOneTest(name, part, expAns string, parser func(testName string), solver func() string) (string, error) {
	if strings.Compare(expAns, "0") == 0 {
		return clr.Str(fmt.Sprintf("%v : %v : test-not-found", name, part), clr.Yellow), nil
	}
	parser(name)
	gotAns := solver()
	if strings.Compare(gotAns, "0") == 0 {
		return clr.Str(fmt.Sprintf("%v : %v : unsolved", name, part), clr.Blue), nil
	}

	testPassed := strings.Compare(expAns, gotAns) == 0
	if testPassed {
		return clr.Str(fmt.Sprintf("%v : %v : passed", name, part), clr.Green), nil
	} else {
		return clr.Str(fmt.Sprintf("%v : %v : failed : expected (%v) got (%v)",
			name, part, expAns, gotAns), clr.Red), fmt.Errorf("Z")
	}
}
