package mapz

import "fmt"

type MMIntInt struct {
	mapp map[string]int
}

func NewMMIntInt() *MMIntInt {
	return &MMIntInt{
		mapp: make(map[string]int),
	}
}

func (m *MMIntInt) Put(k1, k2, v int) {
	m.mapp[makeKey(k1, k2)] = v
}

func (m *MMIntInt) Get(k1, k2 int) (int, bool) {
	v, ok := m.mapp[makeKey(k1, k2)]
	return v, ok
}

func (m *MMIntInt) Del(k1, k2 int) {
	delete(m.mapp, makeKey(k1, k2))
}

func (m *MMIntInt) Size() int {
	return len(m.mapp)
}

func makeKey(k1, k2 int) string {
	return fmt.Sprintf("%v-%v", k1, k2)
}
