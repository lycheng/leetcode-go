package str

// Stack use for FILO
type Stack struct {
	Data []string
}

func newStack() *Stack {
	s := Stack{}
	s.Data = make([]string, 0)
	return &s
}

func (s *Stack) push(item string) {
	s.Data = append(s.Data, item)
}

func (s *Stack) pop() bool {
	l := len(s.Data)
	if l <= 0 {
		return false
	}
	s.Data = s.Data[:l-1]
	return true
}

func (s *Stack) top() string {
	l := len(s.Data)
	if l <= 0 {
		return ""
	}
	return s.Data[l-1]
}

func (s *Stack) isEmpty() bool {
	return len(s.Data) == 0
}

// https://leetcode.com/problems/valid-parentheses/description/
func isValid(s string) bool {
	pair := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	stack := newStack()
	for _, c := range s {
		symbol := string(c)
		pairSymbol := pair[symbol]

		if pairSymbol == "" {
			stack.push(symbol)
		} else if stack.isEmpty() {
			return false
		} else {
			if stack.top() != pairSymbol {
				return false
			}
			stack.pop()
		}
	}
	return stack.isEmpty()
}
