package modthree

import (
	"fmt"
	"fsm/fsm"
)

// Define the specific states and symbols for this problem.
const (
	S0 fsm.State = "S0"
	S1 fsm.State = "S1"
	S2 fsm.State = "S2"

	ZERO fsm.Symbol = "0"
	ONE  fsm.Symbol = "1"
)

var (
	modThreeFSM *fsm.FSM
)

// init runs once when the package is first used. It configures and
// initializes the mod-three FSM based on the document's definition.
func init() {
	// Q = {S0, S1, S2}
	states := []fsm.State{S0, S1, S2}

	// Σ = {0, 1}
	alphabet := []fsm.Symbol{ZERO, ONE}

	// q0 = S0
	initialState := S0

	// F = {S0, S1, S2}
	acceptingStates := []fsm.State{S0, S1, S2}

	// Define the transition function
	transitions := fsm.TransitionFunction{
		S0: {ZERO: S0, ONE: S1},
		S1: {ZERO: S2, ONE: S0},
		S2: {ZERO: S1, ONE: S2},
	}

	// Use the FSM library's constructor to create our specific FSM.
	var err error
	modThreeFSM, err = fsm.New(states, alphabet, initialState, acceptingStates, transitions)
	if err != nil {
		// not sure how we can restart the program if this fails, so panic.
		panic(fmt.Sprintf("failed to initialize ModThree FSM: %v", err))
	}
}

// ModThree calculates n mod 3 for a binary string.
// It uses the pre-configured FSM to process the input and maps the
// final state to an integer result.
func ModThree(binaryString string) (int, error) {
	// Convert the input string into a slice of Symbols for the FSM library.
	symbols := make([]fsm.Symbol, len(binaryString))
	for i, char := range binaryString {
		symbols[i] = fsm.Symbol(char)
	}

	finalState, err := modThreeFSM.Process(symbols)
	if err != nil {
		return 0, err
	}

	// Convert the final state to the integer output
	switch finalState {
	case S0:
		return 0, nil
	case S1:
		return 1, nil
	case S2:
		return 2, nil
	default:
		// This should be unreachable if the FSM is configured correctly.
		return 0, fmt.Errorf("unknown final state: %q", finalState)
	}
}
