package main

import "testing"

// TestAverage Testing average from package mymath
func TestAverage(t *testing.T) {
	ret := average(2, 3)
	if ret != 2.5 {
		t.Error("Expected 2.5, got", ret)
	}
}
