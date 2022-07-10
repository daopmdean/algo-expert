package main

func NodeSwap(head *LinkedList) *LinkedList {
	tempNode := &LinkedList{Value: 0}
	tempNode.Next = head

	preNode := tempNode
	for preNode.Next != nil && preNode.Next.Next != nil {
		firstNode := preNode.Next
		secondNode := preNode.Next.Next

		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode
		preNode.Next = secondNode

		preNode = firstNode
	}
	return tempNode.Next
}
