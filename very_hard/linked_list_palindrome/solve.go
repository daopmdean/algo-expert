package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func LinkedListPalindrome(head *LinkedList) bool {
	slowNode := head
	fastNode := head
	for fastNode != nil && fastNode.Next != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
	}

	reversedSecondHalfNode := reverseLinkedList(slowNode)
	firstHalfNode := head

	for reversedSecondHalfNode != nil {
		if reversedSecondHalfNode.Value != firstHalfNode.Value {
			return false
		}
		reversedSecondHalfNode = reversedSecondHalfNode.Next
		firstHalfNode = firstHalfNode.Next
	}

	return true
}

func reverseLinkedList(head *LinkedList) *LinkedList {
	var previousNode *LinkedList = nil
	var currentNode = head
	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
	return previousNode
}
