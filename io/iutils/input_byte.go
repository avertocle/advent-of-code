/*
- some funcs recieve error as an arg to be chainable, bad design but looks cool
- no error handling unless absolutely required
#competitive-programming #in-that-zone
*/

package iutils

import (
	"strings"
)

/*
ExtractByte2DFromString1D splits a list of lines into a [][]byte, super flexible with what to extract

"abcd" ; sep = "", cols=nil ; ==> a,b,c,d
"pqrs"	                         p,q,r,s

"abc,def" ; sep = ",", cols=[0,1] ==> a,b,d,e
"pqr,stu"	   	                      p,q,s,t
*/
func ExtractByte2DFromString1D(lines []string, sep string, cols []int, defaultVal byte) [][]byte {
	ans := make([][]byte, len(lines))
	var tokens []string
	var tempRow []byte
	var tempCell []byte
	for i, line := range lines {
		tokens = strings.Split(line, sep)
		tempRow = make([]byte, 0)
		for _, t := range tokens {
			tempCell = []byte(t)
			if cols == nil {
				tempRow = append(tempRow, tempCell...)
			} else {
				for _, c := range cols {
					if len(tempCell) < c {
						tempRow = append(tempRow, tempCell[c])
					}
					tempRow = append(tempRow, defaultVal)
				}
			}
		}
		ans[i] = tempRow
	}
	return ans
}

/*
ExtractByte1DFromString1D extracts []byte consisting of one char from each line
"ab,cd" ==> [a,p] or [b,q] etc
"pq,rs"
*/
func ExtractByte1DFromString1D(lines []string, sep string, col int, defaultVal byte) []byte {
	ans := make([]byte, len(lines))
	var temp []byte
	for i, line := range lines {
		temp = ExtractByte1DFromString0D(line, sep, col, defaultVal)
		ans[i] = temp[col]
	}
	return ans
}

/*
ExtractByte1DFromString0D extracts []byte consisting of one char from token
"a-b-c"==> [a,b,c]
"ab,cd,ef" ==> [a,c,e] or [b,d,f]
"abc,de,fgh" ; col=2 ; defaultVal=0 ==> [c,0,h]
*/
func ExtractByte1DFromString0D(line string, sep string, col int, defaultVal byte) []byte {
	tokens := strings.Split(line, sep)
	ans := make([]byte, len(tokens))
	for i, t := range tokens {
		ans[i] = defaultVal
		if len(t) < col {
			ans[i] = []byte(t)[col]
		}
	}
	return ans
}

//
//func String1DToByte2D(lines []string) [][]byte {
//	iutils := make([][]byte, len(lines))
//	for i, row := range lines {
//		iutils[i] = []byte(row)
//	}
//	return iutils
//}
