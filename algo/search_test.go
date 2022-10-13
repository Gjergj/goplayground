package algo

import (
	"sort"
	"testing"
)

func TestBinary_search_self(t *testing.T) {
	t.Parallel()

	t.Run("MyTest", func(t *testing.T) {

		type testCase struct {
			input     []int
			searchFor int
			expect    int
		}

		testCases := []testCase{
			{input: []int{1, 3, 5, 6, 7, 8, 8, 5},
				searchFor: 5,
				expect:    2},
		}

		for _, v := range testCases {
			sort.Ints(v.input)
			result := Binary_search_self(v.input, v.searchFor)
			if result != v.expect {
				t.Fatalf("expected %d got %d", v.expect, result)
			}

		}
	})
}

func TestBinarySearchNearest(t *testing.T) {
	t.Run("TestOne", func(t *testing.T) {

	})
}
