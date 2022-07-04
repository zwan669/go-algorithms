package dataStructure

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack[T any] struct {
	data []T
}

func MakeStack[T any](l int, args ...int) Stack[T] {
	if len(args) >= 1 {
		panic("Too many arguments...")
	}
	c := 0
	if len(args) == 1 {
		c = args[0]
	}
	stack := Stack[T]{}
	stack.makeStack(l, c)
	return stack
}

func (s *Stack[T]) makeStack(l, c int) {
	s.data = make([]T, l, c)
}

func (s *Stack[T]) Top() T {
	return s.data[len(s.data)-1]
}

func (s *Stack[T]) Push(element T) {
	s.data = append(s.data, element)
}

func (s *Stack[T]) Pop(element T) T {
	temp := s.Top()
	s.data = s.data[:len(s.data)-1]
	return temp
}

func (s *Stack[T]) Size(element T) int {
	return len(s.data)
}

func (s *Stack[T]) IsEmpty(element T) bool {
	return len(s.data) == 0
}
