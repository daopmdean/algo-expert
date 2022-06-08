package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func FindLoop(head *LinkedList) *LinkedList {
	first := head.Next
	second := first.Next
	for first != second {
		first, second = first.Next, second.Next.Next
	}
	first = head
	for first != second {
		first = first.Next
		second = second.Next
	}
	return first
}
