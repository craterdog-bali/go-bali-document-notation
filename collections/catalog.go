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

// ASSOCIATION IMPLEMENTATION

// This constructor creates a new association with the specified key and value.
func Association[K any, V any](key K, value V) abstractions.AssociationLike[K, V] {
	return &association[K, V]{key, value}
}

// This type defines the structure and methods associated with a key-value
// pair. This type is parameterized as follows:
//   - K is an ordered type of key.
//   - V is any type of value.
//
// This structure is used by the catalog type to maintain its associations.
type association[K any, V any] struct {
	key   K
	value V
}

// This method returns the key for this association.
func (v *association[K, V]) GetKey() K {
	return v.key
}

// This method returns the value for this association.
func (v *association[K, V]) GetValue() V {
	return v.value
}

// This method sets the value of this association to a new value and returns the
// old value.
func (v *association[K, V]) SetValue(value V) V {
	var old = v.value
	v.value = value
	return old
}

// This constructor creates a new empty catalog.
func Catalog[K any, V any]() abstractions.CatalogLike[K, V] {
	var keys = map[any]abstractions.AssociationLike[K, V]{}
	var associations = List[abstractions.AssociationLike[K, V]]()
	return &catalog[K, V]{associations, associations, associations, keys}
}

// CATALOG IMPLEMENTATION

// This type defines the structure and methods associated with a catalog of
// key-value pair associations. This type is parameterized as follows:
//   - K is a primitive Go type of ordered key.
//   - V is any type of value.
type catalog[K any, V any] struct {
	// Note: The delegated methods don't see the real collection type.
	abstractions.Sequential[abstractions.AssociationLike[K, V]]
	abstractions.Indexed[abstractions.AssociationLike[K, V]]
	associations abstractions.ListLike[abstractions.AssociationLike[K, V]]
	keys         map[any]abstractions.AssociationLike[K, V]
}

// ASSOCIATIVE INTERFACE

// This method appends the specified association to the end of this catalog.
func (v *catalog[K, V]) AddAssociation(association abstractions.AssociationLike[K, V]) {
	var key = association.GetKey()
	var value = association.GetValue()
	v.SetValue(key, value) // This copies the association.
}

// This method appends the specified associations to the end of this catalog.
func (v *catalog[K, V]) AddAssociations(associations []abstractions.AssociationLike[K, V]) {
	for _, association := range associations {
		// Add the association to the end of this catalog
		v.AddAssociation(association)
	}
}

// This method returns the keys for this catalog.
func (v *catalog[K, V]) GetKeys() []K {
	var keys = make([]K, v.associations.GetSize())
	var index = 0
	var iterator = agents.Iterator[abstractions.AssociationLike[K, V]](v.associations)
	for iterator.HasNext() {
		// Extract the key.
		var association = iterator.GetNext()
		keys[index] = association.GetKey()
		index++
	}
	return keys
}

// This method returns the value that is associated with the specified key in
// this catalog.
func (v *catalog[K, V]) GetValue(key K) V {
	var value V // Set the return value to its zero value.
	var association, exists = v.keys[key]
	if exists {
		// Extract the value.
		value = association.GetValue()
	}
	return value
}

// This method returns the values associated with the specified keys for this
// catalog. The values are returned in the same order as the keys in the
// catalog.
func (v *catalog[K, V]) GetValues(keys []K) []V {
	var values = make([]V, len(keys))
	var index = 0
	for _, key := range keys {
		// Extract the value.
		values[index] = v.GetValue(key)
		index++
	}
	return values
}

// This method sets the value associated with the specified key to the
// specified value and returns the old value if one already existed in the
// catalog.
func (v *catalog[K, V]) SetValue(key K, value V) V {
	var old V // Set the return value to its zero value.
	var association, exists = v.keys[key]
	if exists {
		// Set the value of an existing association.
		old = association.SetValue(value)
	} else {
		// Add a new association.
		association = Association[K, V](key, value)
		v.associations.AddItem(association)
		v.keys[key] = association
	}
	return old
}

// This method removes the association associated with the specified key from the
// catalog and returns it.
func (v *catalog[K, V]) RemoveValue(key K) V {
	var old V // Set the return value to its zero value.
	var association, exists = v.keys[key]
	if exists {
		var index = v.associations.GetIndex(association)
		v.associations.RemoveItem(index)
		old = association.GetValue()
		delete(v.keys, key)
	}
	return old
}

// This method removes the associations associated with the specified keys from
// the catalog and returns the removed values.
func (v *catalog[K, V]) RemoveValues(keys []K) []V {
	var values = make([]V, len(keys))
	var index = 0
	for _, key := range keys {
		// Remove the association associated with the key.
		values[index] = v.RemoveValue(key)
		index++
	}
	return values
}

// This method removes all associations from this catalog.
func (v *catalog[K, V]) RemoveAll() {
	v.keys = map[any]abstractions.AssociationLike[K, V]{}
	v.associations.RemoveAll()
}

// This method sorts this catalog using the canonical rank function to compare
// the keys.
func (v *catalog[K, V]) SortAssociations() {
	v.associations.SortItems()
}

// This method sorts this catalog using the specified rank function to compare
// the keys.
func (v *catalog[K, V]) SortAssociationsWithRanker(rank abstractions.RankingFunction) {
	v.associations.SortItemsWithRanker(rank)
}

// This method reverses the order of all associations in this catalog.
func (v *catalog[K, V]) ReverseAssociations() {
	v.associations.ReverseItems()
}

// CATALOGS LIBRARY

// This constructor creates a new catalogs library for the specified generic
// key and value types.
func Catalogs[K any, V any]() *catalogs[K, V] {
	return &catalogs[K, V]{}
}

// This type defines the library functions that operate on catalogs. Since
// catalogs have parameterized key and value types this library type is also
// parameterized as follows:
//   - K is a primitive Go type of ordered key.
//   - V is any type of value.
type catalogs[K any, V any] struct{}

// BLENDABLE INTERFACE

// This library function returns a new catalog containing all of the associations
// that are in the specified catalogs in the order that they appear in each
// catalog.
func (l *catalogs[K, V]) Merge(first, second abstractions.CatalogLike[K, V]) abstractions.CatalogLike[K, V] {
	var result = Catalog[K, V]()
	result.AddAssociations(first.AsArray())
	result.AddAssociations(second.AsArray())
	return result
}

// This library function returns a new catalog containing only the associations
// that are in the specified catalog that have the specified keys. The
// associations in the resulting catalog will be in the same order as the
// specified keys.
func (l *catalogs[K, V]) Extract(catalog abstractions.CatalogLike[K, V], keys []K) abstractions.CatalogLike[K, V] {
	var result = Catalog[K, V]()
	for _, key := range keys {
		var value = catalog.GetValue(key)
		result.SetValue(key, value)
	}
	return result
}