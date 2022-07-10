package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func NodeSwapRecursive(head *LinkedList) *LinkedList {
	recursiveList(head)
	return head
}

func recursiveList(node *LinkedList) {
	if node == nil || node.Next == nil {
		return
	}

	node.Value, node.Next.Value = node.Next.Value, node.Value

	if node.Next.Next != nil {
		recursiveList(node.Next.Next)
	}
}
