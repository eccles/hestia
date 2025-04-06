package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// nolint: gochecknoglobals // these are constants in a test package
var (
	inputSet            = NewStringSet([]string{"a1", "a2", "a2"}...)
	inputMap1           = map[string]string{"a1": "", "a2": ""}
	inputList2          = []string{"a2", "a3", "a2"}
	inputMap2           = map[string]string{"a2": "", "a3": ""}
	outputList1         = []string{"a1", "a2"}
	outputList2         = []string{"a2", "a3"}
	intersection        = []string{"a2"}
	union               = []string{"a1", "a2", "a3"}
	difference          = []string{"a1"}
	symmetricDifference = []string{"a1", "a3"}
)

// TestCreationListAddRemove tests simple additon and removal of fields.
func TestCreationListAddRemove(t *testing.T) {
	s := NewStringSet(inputList2...)
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
	s := NewStringSetFromMap(inputMap1)
	assert.ElementsMatch(t, outputList1, s.List())
}

// TestIntersection tests intersection of 2 sets.
func TestIntersection(t *testing.T) {
	s2 := NewStringSet(inputList2...)
	assert.ElementsMatch(t, intersection, inputSet.Intersection(s2).List())
}

// TestIntersectionList tests intersection of set and list.
func TestIntersectionList(t *testing.T) {
	assert.ElementsMatch(t, intersection, inputSet.IntersectionList(inputList2...).List())
}

// TestIntersectionMap tests intersection of set and map.
func TestIntersectionMap(t *testing.T) {
	assert.ElementsMatch(t, intersection, inputSet.IntersectionMap(inputMap2).List())
}

// TestUnion tests union of 2 sets.
func TestUnion(t *testing.T) {
	s2 := NewStringSet(inputList2...)
	assert.ElementsMatch(t, union, inputSet.Union(s2).List())
}

// TestUnionList tests union of a set and a list.
func TestUnionList(t *testing.T) {
	assert.ElementsMatch(t, union, inputSet.UnionList(inputList2...).List())
}

// TestUnionMap tests union of a set and a map.
func TestUnionMap(t *testing.T) {
	assert.ElementsMatch(t, union, inputSet.UnionMap(inputMap2).List())
}

// TestDifference tests difference of 2 sets.
func TestDifference(t *testing.T) {
	s2 := NewStringSet(inputList2...)
	assert.ElementsMatch(t, difference, inputSet.Difference(s2).List())
}

// TestDifferenceList tests difference of set and list.
func TestDifferenceList(t *testing.T) {
	assert.ElementsMatch(t, difference, inputSet.DifferenceList(inputList2...).List())
}

// TestDifferenceMap tests difference of set and map.
func TestDifferenceMap(t *testing.T) {
	assert.ElementsMatch(t, difference, inputSet.DifferenceMap(inputMap2).List())
}

// TestSymmetricDifference tests symmetric difference of 2 sets.
func TestSymmetricDifference(t *testing.T) {
	s2 := NewStringSet(inputList2...)
	assert.ElementsMatch(t, symmetricDifference, inputSet.SymmetricDifference(s2).List())
}

// TestSymmetricDifferenceList tests symmetric difference of set and list.
func TestSymmetricDifferenceList(t *testing.T) {
	assert.ElementsMatch(t, symmetricDifference, inputSet.SymmetricDifferenceList(inputList2...).List())
}

// TestSymmetricDifferenceMap tests symmetric difference of set and map.
func TestSymmetricDifferenceMap(t *testing.T) {
	assert.ElementsMatch(t, symmetricDifference, inputSet.SymmetricDifferenceMap(inputMap2).List())
}
