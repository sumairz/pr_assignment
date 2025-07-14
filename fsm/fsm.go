// Package fsm provides a generic implementation of a Finite State Machine (FSM).
// It allows for the creation and processing of any FSM defined by the 5-tuple
package fsm

import "fmt"

// Define generic types for states and symbols for maximum flexibility.
type State string
type Symbol string

// TransitionFunction defines the delta mapping'
// It is represented as a map where map[State]map[Symbol]State.
type TransitionFunction map[State]map[Symbol]State

// FSM represents a Finite State Machine, defined by the 5-tuple.
type FSM struct {
	states          map[State]struct{}
	alphabet        map[Symbol]struct{}
	initialState    State
	acceptingStates map[State]struct{}
	transitionFunc  TransitionFunction
}

// New creates and validates a new FSM.
// It ensures that the initial state is in the set of states and that
// all accepting states are a subset of the total states.
func New(
	states []State,
	alphabet []Symbol,
	initialState State,
	acceptingStates []State,
	transitions TransitionFunction,
) (*FSM, error) {

	stateSet := make(map[State]struct{})
	for _, s := range states {
		stateSet[s] = struct{}{}
	}

	if _, ok := stateSet[initialState]; !ok {
		return nil, fmt.Errorf("initial state %q is not in the set of states", initialState)
	}

	acceptingSet := make(map[State]struct{})
	for _, s := range acceptingStates {
		if _, ok := stateSet[s]; !ok {
			return nil, fmt.Errorf("accepting state %q is not in the set of states", s)
		}
		acceptingSet[s] = struct{}{}
	}

	alphabetSet := make(map[Symbol]struct{})
	for _, s := range alphabet {
		alphabetSet[s] = struct{}{}
	}

	return &FSM{
		states:          stateSet,
		alphabet:        alphabetSet,
		initialState:    initialState,
		acceptingStates: acceptingSet,
		transitionFunc:  transitions,
	}, nil
}

// Process runs the FSM against a sequence of input symbols.
// It starts in the initial state and moves through states based on the transition function.
// It returns the final state reached after all symbols are processed.
// An error is returned if a symbol is not in the FSM's alphabet.
func (m *FSM) Process(input []Symbol) (State, error) {
	currentState := m.initialState

	for _, symbol := range input {
		if _, ok := m.alphabet[symbol]; !ok {
			return "", fmt.Errorf("input symbol %q is not in the alphabet", symbol)
		}

		// Look up the next state from the transition function.
		if stateTransitions, ok := m.transitionFunc[currentState]; ok {
			if nextState, ok := stateTransitions[symbol]; ok {
				currentState = nextState
			} else {
				// This case implies an incomplete transition function, which could be an error.
				// For this implementation, we assume the machine halts on undefined transitions.
				return "", fmt.Errorf("no transition defined for state %q with symbol %q", currentState, symbol)
			}
		} else {
			return "", fmt.Errorf("no transitions defined for state %q", currentState)
		}
	}

	return currentState, nil
}
