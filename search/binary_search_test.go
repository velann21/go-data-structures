package main

import "testing"

func TestSearch_BinarySearch(t *testing.T) {
	var sr = Search{}
	sr.BinarySearch(nil, 10)
}
