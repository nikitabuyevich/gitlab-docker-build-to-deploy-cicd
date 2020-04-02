package main

import "testing"

func TestMain(t *testing.T) {
	// Use this to test CI/CD failing
	// If tests fail
	if false {
		t.Errorf("Something went wrong!")
	}
}
