package strset

import (
	"fmt"
	"math"
	"strings"
)

var (
	// helpful to not write struct{}{} everywhere
	keyExists   = struct{}{}
	nonExistent string
)

type Set struct {
	m map[string]struct{}
}

// New creates and initializes a new set.
func New(ts ...string) *Set {
	s := Set{make(map[string]struct{})}
	for _, item := range ts {
		s.m[item] = keyExists
	}
	return &s
}

// NewWithSize creates a new set and gives make map a size hint.
func NewWithSize(size int) *Set {
	return &Set{make(map[string]struct{}, size)}
}

// Add includes he specified items (one or more) to the set.
// The underlying Set s is modified. If passed nothing it silently returns.
func (s *Set) Add(items ...string) {
	for _, item := range items {
		s.m[item] = keyExists
	}
}

// Clear removes all items from the set.
func (s *Set) Clear() {
	s.m = make(map[string]struct{})
}

// Size returns the number of items in a Set.
func (s *Set) Size() int {
	return len(s.m)
}

// Copy returns a new Set with a copy of s.
func (s *Set) Copy() *Set {
	newSet := NewWithSize(s.Size())
	for item := range s.m {
		newSet.m[item] = keyExists
	}
	return newSet
}

// Has looks for the existence of items passed. It returns false if nothing
// is passed. For multiple items it returns true only if all of the items exist.
func (s *Set) Has(items ...string) bool {
	for _, item := range items {
		_, has := s.m[item]
		if !has {
			return false
		}
	}
	return true
}

// HasAny looks for the existence of any of the items passed. It returns false if
// nothing is passed. For multiple items it returns true if any of the items exist.
func (s *Set) HasAny(items ...string) bool {
	for _, item := range items {
		if _, has := s.m[item]; has {
			return true
		}
	}
	return false
}

// IsEmpty reports whether the Set is empty.
func (s *Set) IsEmpty() bool {
	return len(s.m) == 0
}

// List returns a slice of all items. There is also StringSlice() and
// IntSlice() methods for returning slices of type sting or int.
func (s *Set) List() []string {
	ss := make([]string, 0, s.Size())
	for item := range s.m {
		ss = append(ss, item)
	}
	return ss
}

// Pop deletes and returns an item form the Set. The underlying
// Set s is modified. If Set is empty, the zero value is returned.
func (s *Set) Pop() string {
	for pop := range s.m {
		s.Remove(pop)
		return pop
	}
	return nonExistent
}

// Pop2 tries to delete and return an item from the Set.
// The underlying Set s is modified. The second value is a bool that is
// true if the item existed in the set, and false if not.
// If Set is empty, the zero value and false are returned.
func (s *Set) Pop2() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	for pop := range s.m {
		s.Remove(pop)
		return pop, true
	}
	return nonExistent, true
}

// Remove deletes the specified items from the Set.
// The underlying Set is modified. If passed nothing it silently returns.
func (s *Set) Remove(items ...string) {
	for _, item := range items {
		delete(s.m, item)
	}
}

// Each traverses the items in the Set, calling the provided function
// for each Set member. Traversal will continue until all items in the
// Set have been visited, or if the closure returns false.
func (s *Set) Each(f func(item string) bool) {
	for item := range s.m {
		if !f(item) {
			break
		}
	}
}

// IsEqual test whether s and t are the same in size and have the same items.
func (s *Set) IsEqual(t *Set) bool {
	if s.Size() != t.Size() {
		return false
	}

	for item := range s.m {
		if !t.Has(item) {
			return false
		}
	}
	return true
}

// IsSubset tests whether t is a subset of s.
func (s *Set) IsSubset(t *Set) bool {
	if s.Size() < t.Size() {
		return false
	}

	for item := range t.m {
		if !s.Has(item) {
			return false
		}
	}

	return true
}

// IsSuperset tests whether t is a superset of s.
func (s *Set) IsSuperset(t *Set) bool {
	return t.IsSubset(s)
}

// Merge is like Union, however it modifies the current Set
// it's applied on with the given t Set.
func (s *Set) Merge(t *Set) {
	for item := range t.m {
		s.m[item] = keyExists
	}
}

// Separate removes the Set items containing in t from Set s.
// Please aware that it's not the opposite of Merge.
func (s *Set) Separate(t *Set) {
	for item := range t.m {
		s.Remove(item)
	}
}

// String returns a string representation of s.
func (s *Set) String() string {
	v := make([]string, 0, s.Size())
	for item := range s.m {
		v = append(v, item)
	}
	return fmt.Sprintf("[%v]", strings.Join(v, ", "))
}

// Difference returns a new set which contains items
// which are in th first set but not in the others.
func Difference(set1 *Set, sets ...*Set) *Set {
	v := set1.Copy()
	for _, set := range sets {
		v.Separate(set)
	}
	return v
}

// Intersection returns a new set which contains items that only exist in all given sets.
func Intersection(sets ...*Set) *Set {
	minPos := -1
	minSize := math.MaxInt64
	for i, set := range sets {
		if l := set.Size(); l < minSize {
			minPos = i
			minSize = l
		}
	}
	if minPos == -1 || minSize == 0 {
		return New()
	}

	t := sets[minPos].Copy()
	for i, set := range sets {
		if i == minPos {
			continue
		}
		for item := range t.m {
			if !set.Has(item) {
				t.Remove(item)
			}
		}
	}

	return t
}

// Union is the merger of multiple sets. It returns a new set with all the
// elements present in all the sets that are passed.
func Union(sets ...*Set) *Set {
	maxPos := -1
	maxSize := 0
	for i, set := range sets {
		if l := set.Size(); l > maxSize {
			maxPos = i
			maxSize = l
		}
	}
	if maxSize == 0 {
		return New()
	}

	u := sets[maxPos].Copy()
	for i, set := range sets {
		if i == maxPos {
			continue
		}
		for k := range set.m {
			u.Add(k)
		}
	}

	return u
}

// SymmetricDifference returns a new set which is the difference of items which
// are in one of either, but not in both.
func SymmetricDifference(s *Set, t *Set) *Set {
	u := Union(s, t)
	v := Intersection(s, t)
	return Difference(u, v)
	// // 另一种写法：
	// u := Difference(s, t)
	// v := Difference(t, s)
	// return Union(u, v)
}
