package cmz

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

type MapIB map[int]bool

type MapVisited map[string]bool
