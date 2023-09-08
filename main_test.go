package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4
	res := count(b, false)

	if res != exp {
		t.Errorf("Expected %d, got %d insted.\n", exp, res)
	}
}

// To execute the test, use the go test tool. Please check README.md for instructions. 

func TestConutLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 line1 \n line3\n word3")
	exp := 3
	res := count(b, true)
	
	if res != exp {
		t.Errorf("Expected %d, got %d insted. \n",exp, res)
	}
}