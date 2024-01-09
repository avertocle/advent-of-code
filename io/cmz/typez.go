package cmz

type Primitive interface {
	int | byte | bool
}

type Number interface {
	int | int64 | float64
}

type PrimitivePlus interface {
	int | byte | bool | string
}
