package main

import "testing"

func TestUnpack(t *testing.T) {
	testCases := map[string]string{
		"a4bc2d5e": "aaaabccddddde",
		"abcd": "abcd",
		"45": "",
		"": "",
		"qwe\\4\\5": "qwe45",
		"qwe\\45": "qwe44444",
		"qwe\\\\5": "qwe\\\\\\\\\\",
	}

	for str, want := range testCases {
		val := Unpack(str)

		if want != val {
			t.Fatalf(`Unpack("%s") = %q, want %#q`, str, val, want)
		}
	}
}