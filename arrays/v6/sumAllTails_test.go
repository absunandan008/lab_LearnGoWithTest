package main

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{5, 8})
	want := []int{2, 8}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d and want %d", got, want)
	}
}
