package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func ZipLinkedList(linkedList *LinkedList) *LinkedList {
	if linkedList.Next == nil || linkedList.Next.Next == nil {
		return linkedList
	}

	firstHalfHead := linkedList
	secondHalfHead := splitLinkedList(linkedList)

	reversedSecondHalfHead := reverseLinkedList(secondHalfHead)

	return connect(firstHalfHead, reversedSecondHalfHead)
}

func splitLinkedList(linkedList *LinkedList) *LinkedList {
	slowIterator := linkedList
	fastIterator := linkedList
	for fastIterator != nil && fastIterator.Next != nil {
		slowIterator = slowIterator.Next
		fastIterator = fastIterator.Next.Next
	}

	secondHalfHead := slowIterator.Next
	slowIterator.Next = nil
	return secondHalfHead
}

func connect(list1 *LinkedList, list2 *LinkedList) *LinkedList {
	iterateList1 := list1
	iterateList2 := list2
	for iterateList1 != nil && iterateList2 != nil {
		iterateList1Next := iterateList1.Next
		iterateList2Next := iterateList2.Next

		iterateList1.Next = iterateList2
		iterateList2.Next = iterateList1Next

		iterateList1 = iterateList1Next
		iterateList2 = iterateList2Next
	}

	return list1
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
