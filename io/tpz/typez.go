package tpz

type Primitive interface {
	int | byte | bool
}

type Number interface {
	int | int64 | float32 | float64
}

type PrimitivePlus interface {
	int | int64 | byte | bool | string
}

type MapIIB map[int]map[int]bool

type StringSet = map[string]bool

type Set[T PrimitivePlus] map[T]bool

type MapCost map[string]int
