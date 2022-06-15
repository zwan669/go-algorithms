package main

import "fmt"

func main() {
	minStack := Constructor()
	minStack.Push(2)
	minStack.Push(0)
	minStack.Push(3)
	minStack.Push(0)
	fmt.Println("before pop:")
	minStack.Print()
	for i := 0; i < 3; i++ {
		minStack.Pop()
		fmt.Println("after pop", i+1, ":")
		minStack.Print()
	}
	fmt.Println(minStack.Min())
}

type MinStack struct {
	mainStack []int
	subStack  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		make([]int, 0),
		make([]int, 0),
	}
}

func (m *MinStack) Push(x int) {
	m.mainStack = append(m.mainStack, x)
	if len(m.subStack) == 0 || x <= m.Min() {
		m.subStack = append(m.subStack, x)
	}
}

func (m *MinStack) Pop() {
	if len(m.mainStack) == 0 {
		return
	}
	if m.Top() == m.Min() {
		m.subStack = m.subStack[:len(m.subStack)-1]
	}
	m.mainStack = m.mainStack[:len(m.mainStack)-1]

}

func (m *MinStack) Top() int {
	return m.mainStack[len(m.mainStack)-1]
}

func (m *MinStack) Min() int {
	return m.subStack[len(m.subStack)-1]
}

func (m *MinStack) Print() {
	fmt.Println(m.mainStack)
	fmt.Println(m.subStack)
}
