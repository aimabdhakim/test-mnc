package main

import "fmt"

func question3() {
	var parantheses string
	fmt.Scanln(&parantheses)

	fmt.Println(isValid(parantheses))
}

func isValid(s string) bool {
	// check if the string length is even or not
	// return false if not even
	if len(s)%2 != 0 {
		return false
	}

	// set up stack and map
	st := []rune{}

	// map for checking the open and close parantheses
	open := map[rune]rune{
		'<': '>',
		'{': '}',
		'[': ']',
	}

	// loop over string
	for _, r := range s {

		// check if the current character is the open "< { ["
		// put its close parantheses (the pair "< } ]" ) to the stack and continue
		if closer, ok := open[r]; ok {
			st = append(st, closer)
			continue
		}

		// otherwise, dealing with close parantheses
		// check to make sure the stack isn't empty
		// and whether the top of the stack matches the current chars
		l := len(st) - 1
		if l < 0 || r != st[l] {
			return false
		}

		// take the last element off the stack
		st = st[:l]
	}

	// if the stack empty, true, otherwise false
	return len(st) == 0
}
