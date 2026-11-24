package arraysandslices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 6})
	want := []int{3, 6}
	assertSumAllMsg(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum all tails of correct input", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 6})
		want := []int{2, 6}
		assertSumAllMsg(t, got, want)
	})

	t.Run("sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 6})
		want := []int{0, 6}
		assertSumAllMsg(t, got, want)
	})
}

func assertSumAllMsg(t testing.TB, got, want []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}
