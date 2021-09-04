package kubeconf

import (
	"strconv"
	"sync"
)

type Deduplicator interface {
	GetUniqueName(str string) string
}

func NewDeduplicator() Deduplicator {
	inst := &deduplicator{
		data: make(map[string]int),
	}

	return inst
}

type deduplicator struct {
	data map[string]int
	mu   sync.RWMutex
}

func (d *deduplicator) GetUniqueName(str string) string {
	d.mu.Lock()
	defer d.mu.Unlock()

	newStr := str

	for {
		index := d.data[str]
		d.data[str] = index + 1
		if index > 0 {
			newStr = str + strconv.Itoa(index)
		} else {
			return newStr
		}

		if d.data[newStr] > 0 {
			newStr = str
			continue
		}
		break
	}

	return newStr
}
