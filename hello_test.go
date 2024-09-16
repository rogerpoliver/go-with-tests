package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello()
	expected := "Hello, World!"

	if result != expected {
		t.Errorf("Result: %q, Expected: %q", result, expected)
	}
}

func TestHelloTo(t *testing.T) {
	assertMessage := func(t testing.TB, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("Result: %q, Expected: %q", result, expected)
		}
	}

	t.Run("Should say Hello to people", func(t *testing.T) {
		result := HelloTo("Alice")
		expected := "Hello, Alice"

		assertMessage(t, result, expected)
	})

	t.Run("Should say 'Hello, World' when a empty string is passed by parameter", func(t *testing.T) {
		result := HelloTo("Alice")
		expected := "Hello, Alice"

		assertMessage(t, result, expected)
	})
}
