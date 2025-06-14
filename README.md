# Iterable

A Go package providing a fluent interface for working with slices through functional programming operations. Built with Go's generics, it offers a chainable API for transforming collections while maintaining type safety.

## Features

- 🔄 Method chaining for composing operations
- 🎯 Type-safe operations using Go generics
- 🚀 Efficient in-place mutations when needed
- 🔍 Filter elements using predicates
- 🔄 Transform elements with mapping
- 🎭 Mutate elements in place
- ⚡ Remove duplicates while preserving order

## Installation

```bash
go get github.com/yourusername/iterable
```

## Quick Start

```go
package main

import (
    "fmt"
    "strings"
    "github.com/yourusername/iterable"
)

func main() {
    // Create a new Iterable from a slice
    numbers := []int{1, 2, 2, 3, 3, 4, 5}
    
    // Chain multiple operations
    result := iterable.New(numbers).
        Unique().                                      // Remove duplicates
        Filter(func(n int) bool { return n%2 == 0 }). // Keep even numbers
        Mutate(func(n *int) { *n *= 2 }).            // Double each number
        Collect()                                     // Get the final slice
    
    fmt.Println(result) // Output: [4, 8]
    
    // Working with strings
    words := []string{"hello", "world", "hello", "go"}
    upperWords := iterable.New(words).
        Unique().
        Mutate(func(s *string) { *s = strings.ToUpper(*s) }).
        Collect()
    
    fmt.Println(upperWords) // Output: [HELLO WORLD GO]
    
    // Type conversion using Map
    nums := []int{65, 66, 67}
    letters := iterable.Map(iterable.New(nums), func(n int) string {
        return string(rune(n))
    }).Collect()
    
    fmt.Println(letters) // Output: [A B C]
}
```

## API Reference

### Creating an Iterable

- `New[T comparable](collection []T) *Iterable[T]`
  - Creates a new Iterable from a slice of comparable elements

### Methods

- `Filter(predicate func(item T) bool) *Iterable[T]`
  - Removes elements that don't satisfy the predicate
  - Returns the same Iterable for chaining

- `Mutate(mutate func(item *T)) *Iterable[T]`
  - Modifies elements in place using the provided function
  - Returns the same Iterable for chaining

- `Unique() *Iterable[T]`
  - Removes duplicate elements while preserving order
  - Returns the same Iterable for chaining

- `Collect() []T`
  - Returns the final slice after all operations

- `Len() int`
  - Returns the current number of elements

### Transformations

- `Map[T, U comparable](iter *Iterable[T], mapper func(item T) U) *Iterable[U]`
  - Creates a new Iterable by transforming elements from type T to type U

## Examples

### Filtering and Mutating Numbers

```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

result := iterable.New(numbers).
    Filter(func(n int) bool { return n%2 == 0 }). // Keep even numbers
    Mutate(func(n *int) { *n *= 2 }).            // Double each number
    Collect()

fmt.Println(result) // Output: [4, 8, 12, 16, 20]
```

### Working with Strings

```go
input := []string{"hello", "world", "hello", "go", "world"}

result := iterable.New(input).
    Unique().                                           // Remove duplicates
    Filter(func(s string) bool { return len(s) > 2 }). // Keep strings longer than 2 chars
    Mutate(func(s *string) { *s = strings.Title(*s) }). // Capitalize first letter
    Collect()

fmt.Println(result) // Output: [Hello World Go]
```

### Type Conversion

```go
numbers := []int{1, 2, 3, 4}
booleans := iterable.Map(iterable.New(numbers), func(n int) bool {
    return n%2 == 0
}).Collect()

fmt.Println(booleans) // Output: [false true false true]
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.