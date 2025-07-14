package fsm

import "testing"

// TestFSM_Generic tests the core FSM engine with a simple machine
// that checks for an even number of '1's.
func TestFSM_Generic(t *testing.T) {
	states := []State{"S_even", "S_odd"}
	alphabet := []Symbol{"0", "1"}
	initialState := State("S_even")
	acceptingStates := []State{"S_even"}
	transitions := TransitionFunction{
		"S_even": {"0": "S_even", "1": "S_odd"},
		"S_odd":  {"0": "S_odd", "1": "S_even"},
	}

	fsm, err := New(states, alphabet, initialState, acceptingStates, transitions)
	if err != nil {
		t.Fatalf("Failed to create FSM: %v", err)
	}

	// Table-driven tests
	testCases := []struct {
		name        string
		input       []Symbol
		expected    State
		expectError bool
	}{
		{"Empty input", []Symbol{}, "S_even", false},
		{"Even ones", []Symbol{"1", "0", "1"}, "S_even", false},
		{"Odd ones", []Symbol{"1", "0", "1", "1"}, "S_odd", false},
		{"No ones", []Symbol{"0", "0", "0"}, "S_even", false},
		{"Invalid symbol", []Symbol{"1", "0", "2"}, "", true},
		{"Special characters", []Symbol{"%", "&", "@"}, "", true},

		// edge cases
		{"Single zero", []Symbol{"0"}, "S_even", false},
		{"Single one", []Symbol{"1"}, "S_odd", false},
		{"Multiple zeros", []Symbol{"0", "0", "0"}, "S_even", false},
		{"Multiple ones", []Symbol{"1", "1", "1"}, "S_odd", false},
		{"Alternating zeros and ones", []Symbol{"0", "1", "0", "1"}, "S_even", false},
		{"Long sequence of zeros", []Symbol{"0", "0", "0", "0", "0"}, "S_even", false},
		{"Long sequence of ones", []Symbol{"1", "1", "1", "1", "1"}, "S_odd", false},
		{"Mixed sequence", []Symbol{"1", "0", "1", "0", "1"}, "S_odd", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			finalState, err := fsm.Process(tc.input)
			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error for input %v, but got none", tc.input)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for input %v: %v", tc.input, err)
				}
				if finalState != tc.expected {
					t.Errorf("For input %v, expected state %q, but got %q", tc.input, tc.expected, finalState)
				}
			}
		})
	}
}

// TestFSM_Validation tests the error-checking in the FSM constructor.
func TestFSM_Validation(t *testing.T) {
	_, err := New(
		[]State{"S0"},
		[]Symbol{"0"},
		"S1", // Invalid initial state
		[]State{"S0"},
		TransitionFunction{},
	)
	if err == nil {
		t.Error("Expected error for invalid initial state, but got nil")
	}

	_, err = New(
		[]State{"S0"},
		[]Symbol{"0"},
		"S0",
		[]State{"S1"}, // Invalid accepting state
		TransitionFunction{},
	)
	if err == nil {
		t.Error("Expected error for invalid accepting state, but got nil")
	}
}
