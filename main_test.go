package main

import "testing"

func TestEncode(t *testing.T) {
	text := "<script>"
	expected := "%3Cscript%3E"

	received := encode(text, 1)

	if received != expected {
		t.Fatalf("Excpected: %v, Received: %v", expected, received)
	}
}

func TestEncodeWithDepth(t *testing.T) {
	text := "<script>"
	expected := "%253Cscript%253E"

	received := encode(text, 2)

	if received != expected {
		t.Fatalf("Excpected: %v, Received: %v", expected, received)
	}
}

func TestDecode(t *testing.T) {
	text := "%3Cscript%3E"
	expected := "<script>"

	received := decode(text, 1)

	if received != expected {
		t.Fatalf("Excpected: %v, Received: %v", expected, received)
	}
}

func TestDecodeWithDepth(t *testing.T) {
	text := "%253Cscript%253E"
	expected := "<script>"

	received := decode(text, 2)

	if received != expected {
		t.Fatalf("Excpected: %v, Received: %v", expected, received)
	}
}
