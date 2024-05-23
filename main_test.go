package main

import "testing"

func TestMain(t *testing.T) {
	t.Run("Test Hello World", func(t *testing.T) {
		want := "Hello, World!"
		if got := "Hello, World!"; got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}