package iterable

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type IterableSuite struct {
	suite.Suite
}

func TestIterableSuite(t *testing.T) {
	suite.Run(t, new(IterableSuite))
}

func (s *IterableSuite) TestNew() {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "multiple elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			iter := New(tt.input)
			s.Equal(tt.expected, iter.Collect())
		})
	}
}

func (s *IterableSuite) TestFilter() {
	s.Run("integer filtering", func() {
		tests := []struct {
			name      string
			input     []int
			predicate func(item int) bool
			expected  []int
		}{
			{
				name:      "empty slice",
				input:     []int{},
				predicate: func(int) bool { return true },
				expected:  []int{},
			},
			{
				name:      "filter none",
				input:     []int{1, 2, 3},
				predicate: func(int) bool { return true },
				expected:  []int{1, 2, 3},
			},
			{
				name:      "filter all",
				input:     []int{1, 2, 3},
				predicate: func(int) bool { return false },
				expected:  []int{},
			},
			{
				name:      "filter even numbers",
				input:     []int{1, 2, 3, 4, 5, 6},
				predicate: func(item int) bool { return item%2 == 0 },
				expected:  []int{2, 4, 6},
			},
		}

		for _, tt := range tests {
			s.Run(tt.name, func() {
				result := New(tt.input).Filter(tt.predicate).Collect()
				s.Equal(tt.expected, result)
			})
		}
	})

	s.Run("string filtering", func() {
		input := []string{"hello", "world", "test", "go"}
		result := New(input).
			Filter(func(s string) bool { return len(s) > 3 }).
			Collect()
		s.Equal([]string{"hello", "world", "test"}, result)
	})
}

func (s *IterableSuite) TestMutate() {
	s.Run("integer mutation", func() {
		tests := []struct {
			name     string
			input    []int
			mutator  func(item *int)
			expected []int
		}{
			{
				name:     "empty slice",
				input:    []int{},
				mutator:  func(item *int) { *item *= 2 },
				expected: []int{},
			},
			{
				name:     "double values",
				input:    []int{1, 2, 3},
				mutator:  func(item *int) { *item *= 2 },
				expected: []int{2, 4, 6},
			},
			{
				name:     "set to zero",
				input:    []int{1, 2, 3},
				mutator:  func(item *int) { *item = 0 },
				expected: []int{0, 0, 0},
			},
		}

		for _, tt := range tests {
			s.Run(tt.name, func() {
				result := New(tt.input).Mutate(tt.mutator).Collect()
				s.Equal(tt.expected, result)
			})
		}
	})

	s.Run("string mutation", func() {
		input := []string{"hello", "world"}
		result := New(input).
			Mutate(func(s *string) { *s = strings.ToUpper(*s) }).
			Collect()
		s.Equal([]string{"HELLO", "WORLD"}, result)
	})
}

func (s *IterableSuite) TestMap() {
	s.Run("basic type mapping", func() {
		tests := []struct {
			name     string
			input    []int
			mapper   func(item int) string
			expected []string
		}{
			{
				name:     "empty slice",
				input:    []int{},
				mapper:   func(i int) string { return string(rune(i)) },
				expected: []string{},
			},
			{
				name:  "int to string",
				input: []int{65, 66, 67},
				mapper: func(i int) string {
					return string(rune(i))
				},
				expected: []string{"A", "B", "C"},
			},
		}

		for _, tt := range tests {
			s.Run(tt.name, func() {
				result := Map(New(tt.input), tt.mapper).Collect()
				s.Equal(tt.expected, result)
			})
		}
	})

	s.Run("complex mapping", func() {
		input := []int{1, 2, 3, 4}
		result := Map(New(input), func(i int) bool {
			return i%2 == 0
		}).Collect()
		s.Equal([]bool{false, true, false, true}, result)
	})
}

func (s *IterableSuite) TestChaining() {
	s.Run("multiple operations", func() {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		result := New(input).
			Filter(func(i int) bool { return i%2 == 0 }). // Keep even numbers
			Mutate(func(i *int) { *i *= 2 }).             // Double each number
			Collect()

		s.Equal([]int{4, 8, 12, 16, 20}, result)

		// Chain with Map
		stringResult := Map(New(result), func(i int) string {
			return string(rune(i + 64)) // Convert numbers to uppercase letters (4->D, 8->H, etc)
		}).Collect()

		s.Equal([]string{"D", "H", "L", "P", "T"}, stringResult)
	})
}

func (s *IterableSuite) TestUnique() {
	s.Run("integer deduplication", func() {
		tests := []struct {
			name     string
			input    []int
			expected []int
		}{
			{
				name:     "empty slice",
				input:    []int{},
				expected: []int{},
			},
			{
				name:     "no duplicates",
				input:    []int{1, 2, 3, 4, 5},
				expected: []int{1, 2, 3, 4, 5},
			},
			{
				name:     "with duplicates - order preserved",
				input:    []int{3, 1, 2, 2, 1, 3, 4, 5, 5},
				expected: []int{3, 1, 2, 4, 5},
			},
			{
				name:     "all duplicates",
				input:    []int{1, 1, 1, 1, 1},
				expected: []int{1},
			},
			{
				name:     "zero values",
				input:    []int{0, 0, 1, 0, 2, 0},
				expected: []int{0, 1, 2},
			},
		}

		for _, tt := range tests {
			s.Run(tt.name, func() {
				result := New(tt.input).Unique().Collect()
				s.Equal(tt.expected, result)
			})
		}
	})

	s.Run("string deduplication", func() {
		input := []string{"hello", "world", "hello", "go", "world", "unique"}
		result := New(input).Unique().Collect()
		s.Equal([]string{"hello", "world", "go", "unique"}, result)
	})

	s.Run("chaining with other operations", func() {
		input := []int{4, 2, 2, 3, 4, 3, 6, 6, 5}
		result := New(input).
			Filter(func(i int) bool { return i%2 == 0 }). // Keep even numbers: [4,2,2,4,6,6]
			Unique().                                     // Remove duplicates: [4,2,6]
			Mutate(func(i *int) { *i *= 2 }).             // Double each number: [8,4,12]
			Collect()

		s.Equal([]int{8, 4, 12}, result)
	})
}

func (s *IterableSuite) TestEdgeCases() {
	s.Run("nil handlers", func() {
		s.Panics(func() {
			New([]int{1, 2, 3}).Filter(nil)
		}, "Filter with nil predicate should panic")

		s.Panics(func() {
			New([]int{1, 2, 3}).Mutate(nil)
		}, "Mutate with nil mutator should panic")
	})

	s.Run("zero values", func() {
		input := []int{0, 0, 0}
		result := New(input).
			Filter(func(i int) bool { return i == 0 }).
			Mutate(func(i *int) { *i++ }).
			Collect()
		s.Equal([]int{1, 1, 1}, result)
	})

	s.Run("large slice", func() {
		input := make([]int, 1000)
		for i := range input {
			input[i] = i
		}

		result := New(input).
			Filter(func(i int) bool { return i%2 == 0 }).
			Mutate(func(i *int) { *i *= 2 }).
			Collect()

		s.Len(result, 500)
		s.Equal(0, result[0])
		s.Equal(1996, result[len(result)-1])
	})
}
