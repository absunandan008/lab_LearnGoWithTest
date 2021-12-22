package main

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d and want %d", got, want)
		}
	}

	t.Run("make sums of some slice", func(t *testing.T) {
		got := SumAllTails([]int{3, 4}, []int{4, 5})
		want := []int{4, 5}
		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices ", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}
