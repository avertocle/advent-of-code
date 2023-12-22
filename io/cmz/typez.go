package cmz

type Primitive interface {
	int | byte | bool
}

type PrimitivePlus interface {
	int | byte | bool | string
}
