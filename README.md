# pr_assignment

## Overview

This project demonstrates how to use a Finite State Machine (FSM) to compute the remainder when a binary number is divided by 3 (i.e., `n mod 3`). The FSM approach is efficient and illustrates how state machines can process binary input without converting it to decimal.

## File Structure

- **main.go**  
  The entry point of the application. It reads a binary string from the command line, calls the FSM logic, and prints the result.

- **fsm/**  
  Contains the FSM library and definitions used to build and process finite state machines.

- **fsm/modthree/modthree.go**  
  Implements the FSM specifically for the modulo-three problem. It defines states, symbols, transitions, and provides the `ModThree` function.

## Why FSM?

FSMs are ideal for problems where the next state depends only on the current state and the current input symbol. For modulo operations on binary numbers, an FSM can track the remainder as each bit is processed, making the solution both elegant and efficient.

## How It Works

1. **FSM Definition**  
   - States represent possible remainders (`S0`, `S1`, `S2` for 0, 1, 2).
   - Symbols are binary digits (`0`, `1`).
   - Transitions define how the remainder changes as each bit is read.

2. **Processing Input**  
   - The binary string is converted into a slice of FSM symbols.
   - The FSM processes each symbol, updating its state.
   - The final state after processing all symbols gives the remainder.

3. **Result Mapping**  
   - The final FSM state is mapped to the integer result (0, 1, or 2).

## Usage

```sh
go run main.go <binary_string>
```

**Example:**
```sh
go run main.go 1101
```
Output:
```
The result of 1101 mod 3 is: 0
```

**Run tests:**
```sh
go test -cover ./...
```
This command runs all tests in the project and shows code coverage.


## Steps

1. **Initialize the FSM**  
   - States, symbols, initial state, accepting states, and transitions are set up in `modthree.go`.

2. **Read Input**  
   - `main.go` reads the binary string from the command line.

3. **Convert Input**  
   - Each character is converted to an FSM symbol.

4. **Process FSM**  
   - The FSM processes the symbol slice and determines the final state.

5. **Output Result**  
   - The result is printed to the user.

## Extending

We can extend this project by:
- Implementing FSMs for other modulo operations.
- Improving error handling and input validation.