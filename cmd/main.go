package main

import (
	"fmt"

	"bracket-sequence-validator/internal/validator"
)

func main() {
	tests := []struct {
		payload  []byte
		expected bool
	}{
		{[]byte("([{}])"), true},
		{[]byte("(){}(())[]"), true},
		{[]byte("({]}[()"), false},
		{[]byte("()({]}[()"), false},
		{[]byte("}{]}[()"), false},
		{[]byte("[({]}()"), false},
		{[]byte("(((("), false},
		{[]byte("([)]"), false},
		{[]byte("("), false},
		{[]byte("]"), false},
		{[]byte("(())"), true},
		{[]byte("(()())"), true},
		{[]byte("()()"), true},
		{[]byte("(()"), false},
		{[]byte("()()("), false},
		{[]byte(")()()("), false},
		{[]byte("))(("), false},
		{[]byte("(()(())"), false},
		{[]byte("(()(()"), false},
		{[]byte("())()()("), false},
	}
	for i, test := range tests {
		if validator.Validate(test.payload) != test.expected {
			panic(fmt.Sprintf("bad test case %d", i))
		}
	}
}
