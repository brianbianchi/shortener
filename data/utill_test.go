package data

import (
	"testing"
)

func TestRandSeq(t *testing.T) {
	n := 6
	result := RandSeq(n)
	if len(result) != n {
		t.Fatalf(`RandSeq(%d) returned %s, which is not %d characters long.`, n, result, n)
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
