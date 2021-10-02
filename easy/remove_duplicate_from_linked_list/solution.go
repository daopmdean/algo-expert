package easy

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	currentNode := linkedList

	for currentNode.Next != nil {
		nextNode := currentNode.Next

		for nextNode != nil && currentNode.Value == nextNode.Value {
			if nextNode.Next != nil {
				nextNode = nextNode.Next
			} else {
				nextNode = nil
			}
		}
		currentNode.Next = nextNode

		if currentNode.Next != nil {
			currentNode = currentNode.Next
		}
	}

	return linkedList
}
