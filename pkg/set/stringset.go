package set

type StringSet map[string]struct{}

// NewStringSet creates a new string set with optional input values.
func NewStringSet(values ...string) StringSet {
	s := make(map[string]struct{}, len(values))
	for _, value := range values {
		s[value] = struct{}{}
	}

	return s
}

// NewStringSetFromMap creates a new string set from a map.
func NewStringSetFromMap(values map[string]string) StringSet {
	s := make(map[string]struct{}, len(values))
	for value := range values {
		s[value] = struct{}{}
	}

	return s
}

// List returns the string set as the original array.
func (s StringSet) List() []string {
	result := make([]string, 0, len(s))
	for value := range s {
		result = append(result, value)
	}

	return result
}

// Add adds one or more values to string set.
func (s StringSet) Add(values ...string) {
	for _, value := range values {
		s[value] = struct{}{}
	}
}

// Remove removes items from a strng set.
func (s StringSet) Remove(values ...string) {
	for _, value := range values {
		delete(s, value)
	}
}

// Contains returns true if item is present in the string set.
func (s StringSet) Contains(value string) bool {
	_, c := s[value]

	return c
}

// SymmetricDifference returns string set that consists of items that are not in both sets.
func (s StringSet) SymmetricDifference(other StringSet) StringSet {
	return s.Difference(other).Union(other.Difference(s))
}

// SymmetricDifferenceList returns string set that consists of items that are not in set and list.
func (s StringSet) SymmetricDifferenceList(values ...string) StringSet {
	return s.SymmetricDifference(NewStringSet(values...))
}

// SymmetricDifferenceMap returns string set that consists of items that are not in set and map.
func (s StringSet) SymmetricDifferenceMap(values map[string]string) StringSet {
	return s.SymmetricDifference(NewStringSetFromMap(values))
}

// Difference returns string set that consists of items that are in set and not in second set.
func (s StringSet) Difference(other StringSet) StringSet {
	result := NewStringSet()

	for k := range s {
		if !other.Contains(k) {
			result.Add(k)
		}
	}

	return result
}

// DifferenceList returns string set that consists of items that are in set and not in list.
func (s StringSet) DifferenceList(values ...string) StringSet {
	return s.Difference(NewStringSet(values...))
}

// DifferenceMap returns string set that consists of items that are in set and not in map.
func (s StringSet) DifferenceMap(values map[string]string) StringSet {
	return s.Difference(NewStringSetFromMap(values))
}

// Intersection returns string set that consists of items that are in both sets.
func (s StringSet) Intersection(other StringSet) StringSet {
	result := NewStringSet()

	for k := range s {
		if other.Contains(k) {
			result.Add(k)
		}
	}

	return result
}

// IntersectonList returns string set that consists of items that are in both set and list.
func (s StringSet) IntersectionList(values ...string) StringSet {
	return s.Intersection(NewStringSet(values...))
}

// IntersectonMap returns string set that consists of items that are in both set and map.
func (s StringSet) IntersectionMap(values map[string]string) StringSet {
	return s.Intersection(NewStringSetFromMap(values))
}

// Union returns string set that consists of items that are in either of the 2 sets.
func (s StringSet) Union(other StringSet) StringSet {
	result := NewStringSet()

	for k := range s {
		result.Add(k)
	}

	for k := range other {
		result.Add(k)
	}

	return result
}

// UnionList returns string set that consists of items that are in either the set or the list.
func (s StringSet) UnionList(values ...string) StringSet {
	result := NewStringSet()

	for k := range s {
		result.Add(k)
	}

	for _, k := range values {
		result.Add(k)
	}

	return result
}

// UnionMap returns string set that consists of items that are in either the set or the map.
func (s StringSet) UnionMap(values map[string]string) StringSet {
	result := NewStringSet()

	for k := range s {
		result.Add(k)
	}

	for k := range values {
		result.Add(k)
	}

	return result
}
