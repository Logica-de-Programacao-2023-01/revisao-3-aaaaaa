package bonus

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		l1     *LinkedList
		l2     *LinkedList
		result *LinkedList
	}{
		{
			l1:     createLinkedList([]int{1, 2, 4}),
			l2:     createLinkedList([]int{1, 3, 4}),
			result: createLinkedList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			l1:     createLinkedList([]int{1, 2, 4}),
			l2:     createLinkedList([]int{1, 3, 4, 5}),
			result: createLinkedList([]int{1, 1, 2, 3, 4, 4, 5}),
		},
		{
			l1:     createLinkedList([]int{1, 2, 4, 5}),
			l2:     createLinkedList([]int{1, 3, 4}),
			result: createLinkedList([]int{1, 1, 2, 3, 4, 4, 5}),
		},
		{
			l1:     createLinkedList([]int{1, 2, 4, 5}),
			l2:     createLinkedList([]int{1, 3, 4, 6}),
			result: createLinkedList([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			l1:     createLinkedList([]int{1, 2, 4, 5}),
			l2:     createLinkedList([]int{1, 3, 4, 6, 7}),
			result: createLinkedList([]int{1, 1, 2, 3, 4, 4, 5, 6, 7}),
		},
	}

	for _, test := range tests {
		result := MergeTwoLists(test.l1, test.l2)
		if !compareLinkedList(result.Head, test.result.Head) {
			t.Errorf("MergeTwoLists(%v, %v) => %v,want %v", test.l1, test.l2, result, test.result)
		}
	}
}

func compareLinkedList(l1 *Node, l2 *Node) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	if l1 == nil || l2 == nil {
		return false
	}
	if l1.Value != l2.Value {
		return false
	}
	return compareLinkedList(l1.Next, l2.Next)
}

func createLinkedList(input []int) *LinkedList {
	head := &LinkedList{
		Head: &Node{Value: input[0]},
	}
	current := head.Head
	for i := 1; i < len(input); i++ {
		current.Next = &Node{
			Value: input[i],
		}
		current = current.Next
	}
	return head

}
