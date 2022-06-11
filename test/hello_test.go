package main

import "testing"

func Hello() string {
	return "Hello"
}

func TestHello(t *testing.T) {
	actualString := Hello()
	expectedString := "Hello"
	if actualString != expectedString {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedString, actualString)
	}
}
