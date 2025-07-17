package techniques

import (
	"container/list"
	"fmt"
)

// Node represents a node in a tree structure.
// In practical terms, this could model anything hierarchical, such as a file system, organizational structure, or any tree-like data.
type Node struct {
	Value    int
	Children []*Node
}

// BreadthFirstSearch traverses the tree level by level, starting from the root node.
// It checks each node's value against a target value.
// This is classic breadth-first search (BFS) technique implementation.
// --
// What is BFS?
// BFS is a graph/tree traversal algorithm that explores all nodes at the current depth before moving to the next level.
// It's implemented using a queue (FIFO structure) to process nodes in the order they are discovered.
//
// BFS is particularly useful for:
// - Finding the shortest path in unweighted graphs.
// - Exploring nodes closest to the root first.
// - Level-order processing of tree-like structures
//
// This implementation is general-purpose and can be reused both in algorithms and data structures contexts.
// Time Complexity: O(n), where n is the number of nodes in the tree.
// Space Complexity: O(w), where w is the maximum width of the tree (the maximum number of nodes at any level).
func BreadthFirstSearch(root *Node, target int) *Node {
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		// Get the element in front of the queue (FIFO).
		// remove the element from the queue to process it
		element := queue.Front()
		queue.Remove(element)

		current := element.Value.(*Node)
		fmt.Printf("Visiting current value: %d\n", current.Value)

		if current.Value == target {
			fmt.Printf("Found target value: %d\n", target)
			return current
		}

		for _, child := range current.Children {
			queue.PushBack(child)
		}

	}

	fmt.Println("No target value found")
	return nil
}
