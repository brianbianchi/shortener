package data

import (
	"testing"
)

func TestRandSeq(t *testing.T) {
	seqLen := 6
	result := RandSeq(6)
	if len(result) != seqLen {
		t.Fatalf(`RandSeq(%d) returned %s, which is not %d characters long.`, seqLen, result, seqLen)
	}
}

func TestIsURL(t *testing.T) {
	url := "https://www.google.com/"
	result := IsURL(url)
	expected := true
	if result != expected {
		t.Fatalf(`IsURL(%v) returned %t, expected %t`, url, result, expected)
	}

	url = "https://google.com/"
	result = IsURL(url)
	expected = true
	if result != expected {
		t.Fatalf(`IsURL(%v) returned %t, expected %t`, url, result, expected)
	}

	url = "http://google.com/"
	result = IsURL(url)
	expected = true
	if result != expected {
		t.Fatalf(`IsURL(%v) returned %t, expected %t`, url, result, expected)
	}

	url = "google.com"
	result = IsURL(url)
	expected = false
	if result != expected {
		t.Fatalf(`IsURL(%v) returned %t, expected %t`, url, result, expected)
	}

	url = "www.google.com"
	result = IsURL(url)
	expected = false
	if result != expected {
		t.Fatalf(`IsURL(%v) returned %t, expected %t`, url, result, expected)
	}
}
