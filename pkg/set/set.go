// .*@mycompany\.com MY COMPANY 2025
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package set

import (
	"iter"
	"maps"
)

type (
	Set[T comparable] map[T]struct{}
)

// NewSet creates a new set with optional input values.
func NewSet[T comparable](values iter.Seq[T]) Set[T] {
	s := make(map[T]struct{})

	if values != nil {
		for value := range values {
			s[value] = struct{}{}
		}
	}

	return s
}

// Iter returns an terator over the set.
func (s Set[T]) Iter() iter.Seq[T] {
	return maps.Keys(s)
}

// List returns the set as the original array.
func (s Set[T]) List() []T {
	result := make([]T, 0, len(s))

	for value := range s {
		result = append(result, value)
	}

	return result
}

// Add adds one or more values to set.
func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

// Remove removes items from a set.
func (s Set[T]) Remove(value T) {
	delete(s, value)
}

// Contains returns true if item is present in the set.
func (s Set[T]) Contains(value T) bool {
	_, c := s[value]

	return c
}

// SymmetricDifference returns set that consists of items that are not in both
// sets.
func (s Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	return s.Difference(other).Union(other.Difference(s))
}

// SymmetricDifferenceIter returns set that consists of items that are not in
// set and iterable.
func (s Set[T]) SymmetricDifferenceIter(values iter.Seq[T]) Set[T] {
	// Hmm - how to avoid the NewSet call?
	return s.SymmetricDifference(NewSet(values))
}

// Difference returns set that consists of items that are in set and not in
// second set.
func (s Set[T]) Difference(other Set[T]) Set[T] {
	result := NewSet[T](nil)

	for k := range s {
		if !other.Contains(k) {
			result[k] = struct{}{}
		}
	}

	return result
}

// DifferenceIter returns set that consists of items that are in set and not
// in iterable.
func (s Set[T]) DifferenceIter(values iter.Seq[T]) Set[T] {
	// Hmm - how to avoid the NewSet call?
	return s.Difference(NewSet[T](values))
}

// Intersection returns set that consists of items that are in both sets.
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	return s.IntersectionIter(other.Iter())
}

// IntersectionIter returns set that consists of items that are in both set
// and iterable.
func (s Set[T]) IntersectionIter(values iter.Seq[T]) Set[T] {
	result := NewSet[T](nil)

	for k := range values {
		if s.Contains(k) {
			result[k] = struct{}{}
		}
	}

	return result
}

// Union returns set that consists of items that are in either of the 2 sets.
func (s Set[T]) Union(other Set[T]) Set[T] {
	return s.UnionIter(maps.Keys(other))
}

// UnionList returns set that consists of items that are in either the set or
// the iterable.
func (s Set[T]) UnionIter(values iter.Seq[T]) Set[T] {
	result := NewSet[T](nil)

	for k := range s {
		result[k] = struct{}{}
	}

	for k := range values {
		result[k] = struct{}{}
	}

	return result
}
