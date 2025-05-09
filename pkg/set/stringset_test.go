package set

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// nolint: gochecknoglobals // these are constants in a test package
var (
	inputSet            = NewSet(slices.Values([]string{"a1", "a2", "a2"}))
	inputMap            = maps.Keys(map[string]string{"a1": "", "a2": ""})
	inputList           = slices.Values([]string{"a2", "a3", "a2"})
	outputList1         = []string{"a1", "a2"}
	outputList2         = []string{"a2", "a3"}
	intersection        = []string{"a2"}
	union               = []string{"a1", "a2", "a3"}
	difference          = []string{"a1"}
	symmetricDifference = []string{"a1", "a3"}
)

// TestCreationListAddRemove tests simple additon and removal of fields.
func TestCreationListAddRemove(t *testing.T) {
	s := NewSet(inputList)
	assert.ElementsMatch(t, outputList2, s.List())

	// add something that already exists
	assert.Equal(t, true, s.Contains("a3"))
	s.Add("a3")
	assert.Equal(t, true, s.Contains("a3"))

	// add something that doe not exist
	assert.Equal(t, false, s.Contains("a4"))
	s.Add("a4")
	assert.Equal(t, true, s.Contains("a4"))

	// delete something that does not exist
	assert.Equal(t, false, s.Contains("a5"))
	s.Remove("a5")
	assert.Equal(t, false, s.Contains("a5"))

	// delete something that doesexist
	assert.Equal(t, true, s.Contains("a2"))
	s.Remove("a2")
	assert.Equal(t, false, s.Contains("a2"))
}

// TestCreationMap tests creation from map.
func TestCreationMap(t *testing.T) {
	s := NewSet(inputMap)
	assert.ElementsMatch(t, outputList1, s.List())
}

// TestIntersection tests intersection of 2 sets.
func TestIntersection(t *testing.T) {
	s2 := NewSet(inputList)
	assert.ElementsMatch(t, intersection, inputSet.Intersection(s2).List())
}

// TestIntersectionIter tests intersection of set and iterable.
func TestIntersectionIter(t *testing.T) {
	assert.ElementsMatch(t, intersection, inputSet.IntersectionIter(inputList).List())
}

// TestUnion tests union of 2 sets.
func TestUnion(t *testing.T) {
	s2 := NewSet(inputList)
	assert.ElementsMatch(t, union, inputSet.Union(s2).List())
}

// TestUnionIter tests union of a set and a iterable.
func TestUnionIter(t *testing.T) {
	assert.ElementsMatch(t, union, inputSet.UnionIter(inputList).List())
}

// TestDifference tests difference of 2 sets.
func TestDifference(t *testing.T) {
	s2 := NewSet(inputList)
	assert.ElementsMatch(t, difference, inputSet.Difference(s2).List())
}

// TestDifferenceIter tests difference of set and iterable.
func TestDifferenceIter(t *testing.T) {
	assert.ElementsMatch(t, difference, inputSet.DifferenceIter(inputList).List())
}

// TestSymmetricDifference tests symmetric difference of 2 sets.
func TestSymmetricDifference(t *testing.T) {
	s2 := NewSet(inputList)
	assert.ElementsMatch(t, symmetricDifference, inputSet.SymmetricDifference(s2).List())
}

// TestSymmetricDifferenceIter tests symmetric difference of set and iterable.
func TestSymmetricDifferenceIter(t *testing.T) {
	assert.ElementsMatch(t, symmetricDifference, inputSet.SymmetricDifferenceIter(inputList).List())
}
