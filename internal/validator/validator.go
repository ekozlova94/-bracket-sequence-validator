package validator

import "bracket-sequence-validator/internal/stack"

func Validate(payload []byte) bool {
	length := len(payload)
	if length == 0 || length%2 != 0 {
		return false
	}
	if payload[0] == ')' || payload[0] == '}' || payload[0] == ']' {
		return false
	}
	var s = new(stack.Stack)
	for i := 0; i < length; i++ {
		if payload[i] == '(' || payload[i] == '{' || payload[i] == '[' {
			s.Push(payload[i])
			continue
		}
		top := s.Top()
		if top != nil && getOpposite(top.Value) == payload[i] {
			s.Pop()
			continue
		}
		return false
	}
	if s.Size() == 0 {
		return true
	}
	return false
}

func getOpposite(value byte) byte {
	switch value {
	case '(':
		return ')'
	case '[':
		return ']'
	default:
		return '}'
	}
}
