package main

import "testing"

func TestNormalize(t *testing.T) {
	testCases := []struct{
		input  string
		expect string
	}{
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"(123) 456-7893", "1234567893"},
		{"123-456-7894", "1234567894"},
		{"123456789", "123456789"},
		{"(123)456-7895", "1234567895"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := normalize(tc.input)
			if actual != tc.expect {
				t.Errorf("Got %s, Expect %s", actual, tc.expect)
			}
		})
	}
}
