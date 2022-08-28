package main

func SortStack(stack []int) []int {
	if len(stack) == 0 {
		return stack
	}

	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	SortStack(stack)

	insert(&stack, top)
	return stack
}

func insert(stack *[]int, value int) {
	if len(*stack) == 0 || (*stack)[len(*stack)-1] <= value {
		*stack = append(*stack, value)
		return
	}

	top := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	insert(stack, value)
	*stack = append(*stack, top)
}
