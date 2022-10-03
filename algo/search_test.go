package algo

import "testing"

func TestBinary_search_self(t *testing.T) {
	t.Parallel()

	t.Run("MyTest", func(t *testing.T) {
		if Binary_search_self([]int{1, 3, 5, 6, 7, 8, 8, 5}, 5) != 0 {
			t.Errorf("expected zero")
		}
	})
}
