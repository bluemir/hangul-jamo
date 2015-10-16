package main

import (
	"bytes"
	"testing"
)

func TestDecompose(t *testing.T) {
	cases := []struct {
		input  string
		expect []Hangul
	}{
		{"한글", []Hangul{{0x12, 0x0, 0x4}, {0x0, 0x12, 0x8}}},
	}

	for _, test := range cases {
		var actual HangulString = decompose(test.input)
		if actual.Compare(test.expect) != 0 {
			t.Error("expect:", test.expect, "actual", actual)
		}
	}
}
func TestDecomposeToBytes(t *testing.T) {
	cases := []struct {
		input  string
		expect []byte
	}{
		{"한글", []byte{0x12, 0x0, 0x4, 0x0, 0x12, 0x8}},
	}

	for _, test := range cases {
		str := decompose(test.input)
		actual := HangulString(str).toBytes()
		if bytes.Compare(actual, test.expect) != 0 {
			t.Error("expect:", test.expect, "actual", actual)
		}
	}
}
