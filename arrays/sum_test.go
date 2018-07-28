package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("adding a collection of ints", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumTails(t *testing.T) {

	checkSums := func(t *testing.T, want, got []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("basic test for the ability to sum up the tail of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4})
		want := []int{2, 4}
		checkSums(t, want, got)
	})

	t.Run("test adding an empty set", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3, 4})
		want := []int{0, 9}
		checkSums(t, want, got)
	})
}
