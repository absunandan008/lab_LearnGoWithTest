package main

import "testing"

func TestSearch(t *testing.T) {
	//dictionary := map[string]string{"test": "this is just a test"}
	dictionary := Dictionary{"test": "this is just a test"}

	//got := Search(dictionary, "test")
	got := dictionary.Search("test")
	want := "this is just a test"
	assertString(t, got, want)
	/*
		if got != want {
			t.Errorf("got %q and want %q, given %q", got, want, "test")
		}
	*/
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}

}
