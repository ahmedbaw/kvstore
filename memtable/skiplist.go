package memtable

import "time"

const (
	maxLevel    = 10
	probability = 0.5 // Probability for level promotion
)

type Node struct {
	key     string
	value   []byte
	forward []*Node
}

type SkipList struct {
	header    *Node
	level     int
	size      int
	maxSize   int
	maxAge    time.Duration
	lastFlush time.Time
}

func NewSkipList(maxSize int, maxAge time.Duration) *SkipList {
	header := &Node{
		forward: make([]*Node, maxLevel),
	}
	return &SkipList{
		header:    header,
		level:     0,
		size:      0,
		maxSize:   maxSize,
		maxAge:    maxAge,
		lastFlush: time.Now(),
	}
}
