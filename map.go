package cid

import (
	"time"
)

type Value struct {
	Expir uint64
}

type Map struct {
	CidMap map[Cid]Value
}

// NewMap initializes and returns a new Map.
func NewMap() *Map {
	return &Map{CidMap: make(map[Cid]Value)}
}

// Add puts a Cid in the Map.
func (m *Map) Add(c Cid, expir uint64) {
	m.CidMap[c] = Value{Expir: expir}
}

// Has returns if the Map contains a given Cid.
func (m *Map) Has(c Cid) bool {
	_, ok := m.CidMap[c]
	return ok
}

// Get returns Value from the Map for the given key `c`.
func (m *Map) Get(c Cid) (Value, bool) {
	v, found := m.CidMap[c]
	if !found {
		return Value{}, false
	}
	return v, true
}

// Remove deletes a Cid from the Map.
func (m *Map) Remove(c Cid) {
	delete(m.CidMap, c)
}

// IsExpired returns whether the entry of the given `c' is expired.
// This method assumes the given `c` is an existing key of "m".
func (m *Map) IsExpired(c Cid) bool {
	v, _ := m.Get(c)
	return v.Expir != 0 && v.Expir <= uint64(time.Now().Unix())
}

// HasExpiration returns true if the Value of the
// given key has expir.
func (m *Map) HasExpiration(c Cid) bool {
	v, found := m.Get(c)
	if !found {
		return false
	}
	return v.Expir > 0
}

// Len returns how many elements the Map has.
func (m *Map) Len() int {
	return len(m.CidMap)
}

// Keys returns the Cids in the CidMap.
func (m *Map) Keys() []Cid {
	out := make([]Cid, 0, len(m.CidMap))
	for k := range m.CidMap {
		out = append(out, k)
	}
	return out
}

// Visit adds a Cid to the CidMap only if it is
// not in it already.
func (m *Map) Visit(c Cid) bool {
	if !m.Has(c) {
		m.Add(c, 0)
		return true
	}

	return false
}

// ForEach allows to run a custom function on each
// Cid in the CidMap.
func (m *Map) ForEach(f func(c Cid) error) error {
	for c := range m.CidMap {
		err := f(c)
		if err != nil {
			return err
		}
	}
	return nil
}
