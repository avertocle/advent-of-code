package baze

import (
	"fmt"
	"strconv"
)

var mapHexToBin map[byte][]byte

func init() {
	m := make(map[byte][]byte)
	m['0'] = []byte("0000")
	m['1'] = []byte("0001")
	m['2'] = []byte("0010")
	m['3'] = []byte("0011")
	m['4'] = []byte("0100")
	m['5'] = []byte("0101")
	m['6'] = []byte("0110")
	m['7'] = []byte("0111")
	m['8'] = []byte("1000")
	m['9'] = []byte("1001")
	m['A'] = []byte("1010")
	m['B'] = []byte("1011")
	m['C'] = []byte("1100")
	m['D'] = []byte("1101")
	m['E'] = []byte("1110")
	m['F'] = []byte("1111")
	mapHexToBin = m
}

// no defVal here, ok if it breaks
func HtobQ(h string) []byte {
	hexBytes := []byte(h)
	ans := make([]byte, 0)
	for _, b := range hexBytes {
		ans = append(ans, mapHexToBin[b]...)
	}
	return ans
}

// no defVal here, ok if it breaks
func BTod32Q(bits []byte, defVal int) int {
	ans, err := strconv.ParseInt(string(bits), 2, 32)
	if err != nil {
		fmt.Printf("BTodQ error for (%v) | %v\n", string(bits), err)
		return defVal
	}
	return int(ans)
}
