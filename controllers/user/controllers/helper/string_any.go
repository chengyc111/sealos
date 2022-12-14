/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by set-gen. DO NOT EDIT.

package helper

import (
	"sort"
)

// sets.String is a set of strings, implemented via map[string]struct{} for minimal memory consumption.
type Any map[string]any

// NewAny creates a Any from a list of values.
func NewAny(items map[string]any) Any {
	ss := Any{}
	ss.Insert(items)
	return ss
}

// Insert adds items to the set.
func (s Any) Insert(items map[string]any) Any {
	if items == nil {
		s = make(map[string]any, 0)
		return s
	}
	for k, item := range items {
		s[k] = item
	}
	return s
}

func (s Any) InsertValue(key string, value any) Any {
	s[key] = value
	return s
}

// Delete removes all items from the set.
func (s Any) Delete(items ...string) Any {
	for _, item := range items {
		delete(s, item)
	}
	return s
}

// Has returns true if and only if item is contained in the set.
func (s Any) Has(item string) bool {
	_, contained := s[item]
	return contained
}

// HasAll returns true if and only if all items are contained in the set.
func (s Any) HasAll(items ...string) bool {
	for _, item := range items {
		if !s.Has(item) {
			return false
		}
	}
	return true
}

// HasAny returns true if any items are contained in the set.
func (s Any) HasAny(items ...string) bool {
	for _, item := range items {
		if s.Has(item) {
			return true
		}
	}
	return false
}

// Difference returns a set of objects that are not in s2
// For example:
// s1 = {a1, a2, a3}
// s2 = {a1, a2, a4, a5}
// s1.Difference(s2) = {a3}
// s2.Difference(s1) = {a4, a5}
func (s Any) Difference(s2 Any) Any {
	result := NewAny(nil)
	for key, v := range s {
		if !s2.Has(key) {
			result.InsertValue(key, v)
		}
	}
	return result
}

// Union returns a new set which includes items in either s1 or s2.
// For example:
// s1 = {a1, a2}
// s2 = {a3, a4}
// s1.Union(s2) = {a1, a2, a3, a4}
// s2.Union(s1) = {a1, a2, a3, a4}
func (s1 Any) Union(s2 Any) Any {
	result := NewAny(nil)
	for key, v := range s1 {
		result.InsertValue(key, v)
	}
	for key, v := range s2 {
		result.InsertValue(key, v)
	}
	return result
}

// Intersection returns a new set which includes the item in BOTH s1 and s2
// For example:
// s1 = {a1, a2}
// s2 = {a2, a3}
// s1.Intersection(s2) = {a2}
func (s1 Any) Intersection(s2 Any) Any {
	var walk, other Any
	result := NewAny(nil)
	if s1.Len() < s2.Len() {
		walk = s1
		other = s2
	} else {
		walk = s2
		other = s1
	}
	for key, v := range walk {
		if other.Has(key) {
			result.InsertValue(key, v)
		}
	}
	return result
}

// IsSuperset returns true if and only if s1 is a superset of s2.
func (s1 Any) IsSuperset(s2 Any) bool {
	for item := range s2 {
		if !s1.Has(item) {
			return false
		}
	}
	return true
}

// Equal returns true if and only if s1 is equal (as a set) to s2.
// Two sets are equal if their membership is identical.
// (In practice, this means same elements, order doesn't matter)
func (s1 Any) Equal(s2 Any) bool {
	return len(s1) == len(s2) && s1.IsSuperset(s2)
}

type sortableSliceOfString []string

func (s sortableSliceOfString) Len() int           { return len(s) }
func (s sortableSliceOfString) Less(i, j int) bool { return lessString(s[i], s[j]) }
func (s sortableSliceOfString) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// List returns the contents as a sorted string slice.
func (s Any) ListKey() []string {
	res := make(sortableSliceOfString, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	sort.Sort(res)
	return []string(res)
}

// List returns the contents as a sorted string slice.
func (s Any) ListValue() []any {
	res := make([]any, 0, len(s))
	for _, v := range s {
		res = append(res, v)
	}
	return []any(res)
}

// UnsortedList returns the slice with contents in random order.
func (s Any) UnsortedList() []string {
	res := make([]string, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	return res
}

// Returns a single element from the set.
func (s Any) PopAny() (string, bool) {
	for key := range s {
		s.Delete(key)
		return key, true
	}
	var zeroValue string
	return zeroValue, false
}

// Len returns the size of the set.
func (s Any) Len() int {
	return len(s)
}

func lessString(lhs, rhs string) bool {
	return lhs < rhs
}
