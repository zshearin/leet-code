package main

import (
	"fmt"
)


type ListNode struct {
	Val int
	Next *ListNode
}


func addNodeToEnd(newNode, headNode *ListNode) *ListNode {
	
	if headNode == nil {
		return headNode
	}
	
	for curNode := headNode; curNode != nil; curNode = curNode.Next {
		if curNode.Next == nil {
			curNode.Next = newNode
			return headNode
		}
		
	}
	return headNode
	
}

func printList(headNode *ListNode) {
	for curNode := headNode; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode)
	}
}

func main() {
	fmt.Println("Hello, playground")

	n11 := &ListNode{1, nil}
	n12 := &ListNode{2, nil}
	n13 := &ListNode{3, nil}
	n14 := &ListNode{4, nil}
	n15 := &ListNode{5, nil}
	n16 := &ListNode{6, nil}
	
	n21 := &ListNode{1, nil}
	n22 := &ListNode{1, nil}
	n23 := &ListNode{2, nil}
	n24 := &ListNode{3, nil}
	n25 := &ListNode{5, nil}
	n26 := &ListNode{8, nil}
	
	head1 := n11
	head2 := n21
	
	printList(head1)
	
	head1 = addNodeToEnd(n12, head1)
	head1 = addNodeToEnd(n13, head1)
	head1 = addNodeToEnd(n14, head1)
	head1 = addNodeToEnd(n15, head1)
	head1 = addNodeToEnd(n16, head1)	
	
	head2 = addNodeToEnd(n22, head2)
	head2 = addNodeToEnd(n23, head2)
	head2 = addNodeToEnd(n24, head2)
	head2 = addNodeToEnd(n25, head2)
	head2 = addNodeToEnd(n26, head2)
	
	
	fmt.Println("List1: ")
	printList(head1)
	
	fmt.Println("List2: ")
	printList(head2)
	
	
	solutionDummyHead := &ListNode{}
	
	
	curNode1 := head1
	curNode2 := head2
	
	for {
		var node *ListNode
		
		if (curNode1 == nil && curNode2 == nil) {
			break
		} else if (curNode1 == nil) {
			node = &ListNode{curNode2.Val, nil}
			curNode2 = curNode2.Next
		} else if (curNode2 == nil) {
			node = &ListNode{curNode1.Val, nil}
			curNode1 = curNode1.Next
		} else if (curNode1.Val < curNode2.Val) {
			node = &ListNode{curNode1.Val, nil}
			curNode1 = curNode1.Next
		} else {
			node = &ListNode{curNode2.Val, nil}
			curNode2 = curNode2.Next		
		}
		
		solutionDummyHead = addNodeToEnd(node, solutionDummyHead)
				
	}
	solution := solutionDummyHead.Next
	fmt.Println("Solution:")
	printList(solution)
	
	
	
}

