package mapz

type MMIntIntArr struct {
	mapp map[string][]int
}

func NewMMIntIntArr() *MMIntIntArr {
	return &MMIntIntArr{
		mapp: make(map[string][]int),
	}
}

func (m *MMIntIntArr) Put(k1, k2 int, v []int) {
	m.mapp[makeKey(k1, k2)] = v
}

func (m *MMIntIntArr) AddTo(k1, k2 int, v int) {
	k := makeKey(k1, k2)
	if _, ok := m.mapp[k]; !ok {
		m.mapp[k] = []int{}
	}
	m.mapp[k] = append(m.mapp[k], v)
}

func (m *MMIntIntArr) Get(k1, k2 int) ([]int, bool) {
	v, ok := m.mapp[makeKey(k1, k2)]
	return v, ok
}

func (m *MMIntIntArr) Del(k1, k2 int) {
	delete(m.mapp, makeKey(k1, k2))
}

func (m *MMIntIntArr) Size() int {
	return len(m.mapp)
}
