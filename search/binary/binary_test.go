package binary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		desc     string
		input    []int
		search   int
		expected int
	}{
		{
			desc:     "element in odd list",
			input:    []int{1, 2, 3, 4, 5},
			search:   3,
			expected: 2,
		},
		{
			desc:     "element first in odd list",
			input:    []int{1, 2, 3, 4, 5},
			search:   1,
			expected: 0,
		},
		{
			desc:     "element last in odd list",
			input:    []int{1, 2, 3, 4, 5},
			search:   5,
			expected: 4,
		},
		{
			desc:     "element not in odd list",
			input:    []int{1, 2, 3, 4, 5},
			search:   7,
			expected: -1,
		},
		{
			desc:     "element in even list",
			input:    []int{1, 2, 3, 4, 5, 6},
			search:   3,
			expected: 2,
		},
		{
			desc:     "element first in even list",
			input:    []int{1, 2, 3, 4, 5, 6},
			search:   1,
			expected: 0,
		},
		{
			desc:     "element last in even list",
			input:    []int{1, 2, 3, 4, 5, 6},
			search:   6,
			expected: 5,
		},
		{
			desc:     "element not in even list",
			input:    []int{1, 2, 3, 4, 5, 6},
			search:   7,
			expected: -1,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, Search(tt.input, tt.search))
	}
}
