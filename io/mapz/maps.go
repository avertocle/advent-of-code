package mapz

import "github.com/avertocle/contests/io/tpz"

func FromArr1D[K tpz.PrimitivePlus, V tpz.Number](arr []K, defVal V) map[K]V {
	m := make(map[K]V)
	for _, k := range arr {
		m[k] = defVal
	}
	return m
}

func SumValues[K tpz.PrimitivePlus, V tpz.Number](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func IncValue[K tpz.PrimitivePlus, V tpz.Number](m map[K]V, k K, delta V, initVal V) {
	if _, ok := m[k]; !ok {
		m[k] = initVal
	}
	m[k] = m[k] + delta
}

func DecValue[K tpz.PrimitivePlus, V tpz.Number](m map[K]V, k K, delta V, initVal V) {
	if _, ok := m[k]; !ok {
		m[k] = initVal
	}
	m[k] = m[k] - delta
}

func Keys[T tpz.PrimitivePlus, U any](m map[T]U) []T {
	keys := make([]T, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}
