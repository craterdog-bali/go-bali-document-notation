/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package collections

import (
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"github.com/craterdog-bali/go-bali-document-notation/agents"
)

// SET IMPLEMENTATION

// This constructor creates a new empty set that uses the canonical order
// function to determine the order of the items in the set.
func Set[T any]() abstractions.SetLike[T] {
	return SetWithRanker[T](nil)
}

// This constructor creates a new empty set that uses the specified order
// function to determine the order of the items in the set.
func SetWithRanker[T any](rank abstractions.RankingFunction) abstractions.SetLike[T] {
	// Groom the arguments.
	if rank == nil {
		rank = agents.RankValues
	}

	// Return a new empty set.
	var items = List[T]()
	return &set[T]{items, items, items, rank}
}

// This type defines the structure and methods associated with a set of items.
// The set uses ORDINAL based indexing rather than ZERO based indexing (see
// the description of what this means in the Sequential interface definition).
// This type is parameterized as follows:
//   - T is any type of item.
type set[T any] struct {
	// Note: The delegated methods don't see the real collection type.
	abstractions.Sequential[T]
	abstractions.Indexed[T]
	items abstractions.ListLike[T]
	rank  abstractions.RankingFunction
}

// SEARCHABLE INTERFACE

// This method determines whether or not this set contains the specified item.
func (v *set[T]) ContainsItem(item T) bool {
	var _, found = v.search(item)
	return found
}

// This method determines whether or not this set contains ANY of the
// specified items.
func (v *set[T]) ContainsAny(items []T) bool {
	for _, item := range items {
		if v.ContainsItem(item) {
			// This set contains at least one of the items.
			return true
		}
	}
	// This set does not contain any of the items.
	return false
}

// This method determines whether or not this set contains ALL of the
// specified items.
func (v *set[T]) ContainsAll(items []T) bool {
	for _, item := range items {
		if !v.ContainsItem(item) {
			// This set is missing at least one of the items.
			return false
		}
	}
	// This set does contains all of the items.
	return true
}

// FLEXIBLE INTERFACE

// This method adds the specified item to this set if it is not already a
// member of the set.
func (v *set[T]) AddItem(item T) {
	var slot, found = v.search(item)
	if !found {
		// The item is not already a member, so add it.
		v.items.InsertItem(slot, item)
	}
}

// This method adds the specified items to this set if they are not already
// members of the set.
func (v *set[T]) AddItems(items []T) {
	for _, item := range items {
		v.AddItem(item)
	}
}

// This method removes the specified item from this set. It returns true if the
// item was in the set and false otherwise.
func (v *set[T]) RemoveItem(item T) {
	var index, found = v.search(item)
	if found {
		// The item is a member, so remove it.
		v.items.RemoveItem(index)
	}
}

// This method removes the specified items from this set. It returns the number
// of items that were removed.
func (v *set[T]) RemoveItems(items []T) {
	for _, item := range items {
		v.RemoveItem(item)
	}
}

// This method removes all items from this set.
func (v *set[T]) RemoveAll() {
	v.items.RemoveAll()
}

// PRIVATE INTERFACE

// This private method performs a binary search of the set for the specified
// item. It returns two results:
//   - index: The index of the item, or if not found, the index of the item
//     before which it could be inserted in the underlying list.
//   - found: A boolean stating whether or not the item was found.
//
// The algoritm performs a true O[log(n)] worst case search.
func (v *set[T]) search(item T) (index int, found bool) {
	// We use iteration instead of recursion for better performance.
	//    start        first      middle       last          end
	//    |-------------||----------||----------||-------------|
	//                  |<-- size -------------->|
	//
	var first = 1          // Start at the beginning.
	var last = v.GetSize() // End at the end.
	var size = last        // Initially all items are candidates.
	for size > 0 {
		var middle = first + size/2 // Rounds down to the nearest integer.
		var candidate = v.GetItem(middle)
		switch v.rank(item, candidate) {
		case -1:
			// The index of the item is less than the middle
			// index so the first index stays the same.
			last = middle - 1 // We already tried the middle index.
			size = middle - first
		case 0:
			// The index of the item is the middle index.
			return middle, true
		case 1:
			// The index of the item is greater than the middle
			// index so the last index stays the same.
			first = middle + 1 // We already tried the middle index.
			size = last - middle
		}
	}
	// The item was not found, the last index represents the SLOT where it
	// would be inserted. Note that since the item was not found, the
	// indexes are inverted: last < first (i.e. last = first - 1).
	return last, false
}

// SETS LIBRARY

// This constructor creates a new sets library for the specified generic
// item type.
func Sets[T any]() *sets[T] {
	return &sets[T]{}
}

// This type defines the library functions that operate on sets. Since
// sets have a parameterized item type this library type is also
// parameterized as follows:
//   - T is any type of item.
type sets[T any] struct{}

// LOGICAL INTERFACE

// This library function returns the logical inverse of the specified set.
func (l *sets[T]) Not(set abstractions.SetLike[T]) abstractions.SetLike[T] {
	panic("Not(set) is meaningless, use Sans(fullSet, set) instead.")
}

// This library function returns the logical conjunction of the specified sets.
func (l *sets[T]) And(first, second abstractions.SetLike[T]) abstractions.SetLike[T] {
	var result = Set[T]()
	var iterator = agents.Iterator[T](first)
	for iterator.HasNext() {
		var item = iterator.GetNext()
		if second.ContainsItem(item) {
			result.AddItem(item)
		}
	}
	return result
}

// This library function returns the logical material non-implication of the
// specified sets.
func (l *sets[T]) Sans(first, second abstractions.SetLike[T]) abstractions.SetLike[T] {
	var result = Set[T]()
	result.AddItems(first.AsArray())
	result.RemoveItems(second.AsArray())
	return result
}

// This library function returns the logical disjunction of the specified sets.
func (l *sets[T]) Or(first, second abstractions.SetLike[T]) abstractions.SetLike[T] {
	var result = Set[T]()
	result.AddItems(first.AsArray())
	result.AddItems(second.AsArray())
	return result
}

// This library function returns the logical exclusive disjunction of the
// specified sets.
func (l *sets[T]) Xor(first, second abstractions.SetLike[T]) abstractions.SetLike[T] {
	var result = l.Or(l.Sans(first, second), l.Sans(second, first))
	return result
}