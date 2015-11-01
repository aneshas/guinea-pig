package main

import (
	"testing"
)

func TestFoo(t *testing.T) {
	result := foo("bar")
	if result != "baz" {
		t.Fatal("Test failed")
	}
	t.Log("Test passed")
}
