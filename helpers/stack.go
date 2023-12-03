package helpers

import (
	"fmt"
	"math"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		var r int
		return r, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s *Stack) Print() {
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

func (s *Stack) ConvertToNumber() int {
	total := 0
	powerOfTen := 0

	for !s.IsEmpty() {
		v, _ := s.Top()
		s.Pop()
		powerValue := int(math.Pow10(powerOfTen))
		total += v * powerValue
		powerOfTen++
	}

	return total
}
