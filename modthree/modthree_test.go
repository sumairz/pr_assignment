package modthree

import "testing"

func TestModThree(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    int
		expectError bool
	}{
		{"Binary for 6", "110", 0, false},
		{"Binary for 13", "1101", 1, false},
		{"Binary for 14", "1110", 2, false},
		{"Binary for 15", "1111", 0, false},
		{"Binary for 10", "1010", 1, false},
		// Edge cases
		{"Empty string", "", 0, false},
		{"Single zero", "0", 0, false},
		{"Single one", "1", 1, false},
		{"Special character", "%*(#(%))", 0, true},
		// Error cases
		{"Invalid character", "102", 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ModThree(tc.input)

			if tc.expectError {
				if err == nil {
					t.Errorf("ModThree(%q) expected an error, but got none", tc.input)
				}
			} else {
				if err != nil {
					t.Errorf("ModThree(%q) returned an unexpected error: %v", tc.input, err)
				}
				if result != tc.expected {
					t.Errorf("ModThree(%q) = %d; want %d", tc.input, result, tc.expected)
				}
			}
		})
	}
}
