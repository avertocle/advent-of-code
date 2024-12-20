package mapz

import "github.com/avertocle/contests/io/cmz"

func FromArr1D[K cmz.PrimitivePlus, V cmz.Number](arr []K, defVal V) map[K]V {
	m := make(map[K]V)
	for _, k := range arr {
		m[k] = defVal
	}
	return m
}

func SumValues[K cmz.PrimitivePlus, V cmz.Number](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func IncValue[K cmz.PrimitivePlus, V cmz.Number](m map[K]V, k K, delta V, initVal V) {
	if _, ok := m[k]; !ok {
		m[k] = initVal
	}
	m[k] = m[k] + delta
}

func DecValue[K cmz.PrimitivePlus, V cmz.Number](m map[K]V, k K, delta V, initVal V) {
	if _, ok := m[k]; !ok {
		m[k] = initVal
	}
	m[k] = m[k] - delta
}
