package suuid

import (
	"sort"
	"strings"
)

// StringSet is string sets val
type StringSet struct {
	set    map[string]bool
	list   []string
	sorted bool
}

// NewStringSet is new set
func NewStringSet() *StringSet {
	return &StringSet{make(map[string]bool), make([]string, 0), false}
}

// Add is
func (set *StringSet) Add(i string) bool {
	_, found := set.set[i]
	set.set[i] = true
	if !found {
		set.sorted = false
	}
	return !found //False if it existed already
}

// Contains is
func (set *StringSet) Contains(i string) bool {
	_, found := set.set[i]
	return found //true if it existed already
}

// Remove is
func (set *StringSet) Remove(i string) {
	set.sorted = false
	delete(set.set, i)
}

// Len is
func (set *StringSet) Len() int {
	return len(set.set)
}

// ItemByIndex is
func (set *StringSet) ItemByIndex(idx int) string {
	set.Sort()
	return set.list[idx]
}

// Index is
func (set *StringSet) Index(c string) int {
	for i, s := range set.list {
		if c == s {
			return i
		}
	}
	return 0
}

// Sort is
func (set *StringSet) Sort() {
	if set.sorted {
		return
	}
	set.list = make([]string, 0)
	for s := range set.set {
		set.list = append(set.list, s)
	}
	sort.Strings(set.list)
	set.sorted = true
}

// String is
func (set *StringSet) String() string {
	set.Sort()
	return strings.Join(set.list, "")
}
