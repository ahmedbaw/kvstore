package memtable

type Node struct {
	key   []byte
	value []byte
	next  []*Node
}

type SkipList struct {
	head  *Node
	level int
}

func NewMemetable() *Memtable {
	return &Memtable{}
}

func (m *Memtable) Put(key string, value []byte) {
	//TODO
}

func (m *Memtable) Get(key string) ([]byte, bool) {
	//TODO
	return nil, false
}
