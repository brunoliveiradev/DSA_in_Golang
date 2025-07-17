package techniques

import "fmt"

// DepthFirstSearchUsingStack performs DFS using an explicit stack.
// This avoids the risk of stack overflow in very deep trees.
// ---
// What is DFS?
// DFS (Depth-First Search) explores as far as possible along each branch before backtracking.
// It uses a stack (LIFO structure) to keep track of nodes to visit next.
// It typically uses recursion or a stack if done iteratively.
//
// DFS is useful for:
// - Exploring any deep tree structure.
// - Topological sorting
// - Path finding where depth matters more than breadth (e.g. finding a path in a maze or the deepest node in a tree).
//
// Time Complexity: O(n), where n is the number of nodes in the tree.
// Space Complexity: O(h), where h is the height of the tree (due to the stack).
func DepthFirstSearchUsingStack(root *Node, target int) *Node {
	if root == nil {
		return nil
	}

	stack := []*Node{root}

	// Iterate while there are nodes in the stack
	for len(stack) > 0 {
		n := len(stack) - 1 // Get the last element in the stack (LIFO)
		current := stack[n] // Get the current node to process
		stack = stack[:n]   // Remove the last element from the stack

		fmt.Printf("Visiting current value: %d\n", current.Value)

		if current.Value == target {
			fmt.Printf("Found target value: %d\n", target)
			return current
		}

		// Push all children onto the stack in reverse order to maintain the correct order of processing
		// Left-to-right traversal is recommended for consistency because stacks are LIFO.
		// This means we add the rightmost child first so it gets processed last represent by len(current.Children) - 1.
		// The i >= 0 loop ensures we traverse from the last child to the first child and i-- ensures we process the leftmost child first.
		for i := len(current.Children) - 1; i >= 0; i-- {
			child := current.Children[i]
			fmt.Printf("Adding child value: %d to stack\n", child.Value)
			stack = append(stack, child) // Add the child to the stack
		}
	}

	fmt.Println("No target value found")
	return nil
}

// DepthFirstSearchUsingRecursion performs DFS using recursion.
// This recursive version is elegant and easy to understand,
// though in production systems with deep trees, you must be mindful
// of stack overflows and consider an iterative version instead.
func DepthFirstSearchUsingRecursion(root *Node, target int) *Node {
	if root == nil {
		return nil
	}

	fmt.Printf("Visiting current value: %d\n", root.Value)

	if root.Value == target {
		fmt.Printf("Found target value: %d\n", target)
		return root
	}

	for _, child := range root.Children {
		result := DepthFirstSearchUsingRecursion(child, target)
		if result != nil {
			return result // Return if the target is found in the subtree
		}
	}

	fmt.Println("No target value found")
	return nil
}
