package main

import "testing"

func TestCounter(t *testing.T) {
	// t.Fatal("not implemented")
	t.Run("increamenting the counter 3 times and leaves at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		/*
			if counter.Value() != 3 {
				t.Errorf("got %d but want %d", counter.Value(), 3)
			}
		*/
		assertCounter(t, counter, 3)
	})
}

func assertCounter(t testing.TB, got Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d but want %d", got.Value(), 3)
	}
}
