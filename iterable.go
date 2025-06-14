// Package iterable provides a fluent interface for working with slices of comparable types.
// It enables common functional programming operations like filtering, mapping, and mutation
// while maintaining method chaining capabilities.
package iterable

import (
	"slices"
)

// New creates a new Iterable instance from a slice of comparable elements.
// It serves as the entry point for creating chainable slice operations.
func New[T comparable](collection []T) *Iterable[T] {
	return &Iterable[T]{collection: collection}
}

// Iterable represents a wrapper around a slice that provides chainable operations.
// The type parameter T must satisfy the comparable constraint to ensure elements
// can be compared for equality.
type Iterable[T comparable] struct {
	collection []T
}

// Filter removes elements from the collection that don't satisfy the predicate function.
// It returns the same Iterable instance to enable method chaining.
// The predicate function should return true for elements that should be kept.
func (i *Iterable[T]) Filter(predicate func(item T) bool) *Iterable[T] {
	i.collection = slices.DeleteFunc(i.collection, func(e T) bool {
		return !predicate(e)
	})

	return i
}

// Mutate applies a mutation function to each element in the collection.
// The mutation function receives a pointer to each element, allowing it to modify
// the element in place. Returns the same Iterable instance to enable method chaining.
func (i *Iterable[T]) Mutate(mutate func(item *T)) *Iterable[T] {
	for idx := range i.collection {
		mutate(&i.collection[idx])
	}

	return i
}

// Unique removes duplicate elements from the collection, keeping only the first
// occurrence of each unique element. The order of remaining elements is preserved.
// Returns the same Iterable instance to enable method chaining.
func (i *Iterable[T]) Unique() *Iterable[T] {
	seen := make(map[T]bool)
	result := make([]T, 0, len(i.collection))

	for _, item := range i.collection {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	i.collection = result
	return i
}

// Collect returns the underlying slice containing all elements in the collection.
// This method is typically used at the end of a chain of operations to obtain
// the final result as a standard slice.
func (i *Iterable[T]) Collect() []T {
	return i.collection
}

// Len returns the current number of elements in the collection.
// This method is useful for getting the size of the collection after
// filtering or other operations that may modify its length.
func (i *Iterable[T]) Len() int {
	return len(i.collection)
}

// Map creates a new Iterable by transforming each element in the source Iterable
// using the provided mapper function. The mapper function converts elements of type T
// to elements of type U, where both types must satisfy the comparable constraint.
func Map[T comparable, U comparable](iter *Iterable[T], mapper func(item T) U) *Iterable[U] {
	mapped := make([]U, 0, iter.Len())
	for _, item := range iter.Collect() {
		mapped = append(mapped, mapper(item))
	}

	return New(mapped)
}
