package utils

import "testing"

func TestContains(t *testing.T) {
	var target string
	options := []string{"apple", "pear", "orange"}

	t.Run("Returns true if options include the target", func(t *testing.T) {
		target = options[0]
		got := Contains(options, target)
		want := true
		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("Returns false if options do not include the target", func(t *testing.T) {
		target = "banana"
		got := Contains(options, target)
		want := false
		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})
}
