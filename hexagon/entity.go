package hexagon

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/faiface/pixel"
)

// Entity ..
type Entity struct {
	ID     string
	Sprite *pixel.Sprite
	index  HexIndex
	Type   string
	Ally   bool
	mutex  *sync.RWMutex
}

// setIndex ..
func (e *Entity) setIndex(idx HexIndex) {
	e.index = idx
}

// SetIndex ..
func (e *Entity) SetIndex(idx HexIndex) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.setIndex(idx)
}

// getIndex ..
func (e *Entity) getIndex() HexIndex {
	return e.index
}

// GetIndex ..
func (e *Entity) GetIndex() HexIndex {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.getIndex()
}

// NewEntity returns a new entity
func NewEntity(s *pixel.Sprite, idx HexIndex, t string, ally bool) *Entity {
	id := fmt.Sprintf("%s_%d", t, rand.Intn(100000))

	return &Entity{
		ID:     id,
		Sprite: s,
		index:  idx,
		Type:   t,
		Ally:   ally,
		mutex:  &sync.RWMutex{},
	}
}
