package techniques

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Node represents a node in a tree structure.
func TestBreadthFirstSearch_FindsTargetInTree(t *testing.T) {
	// Sample tree:
	//         1
	//       /   \
	//      2     3
	//    /  \     \
	//   4    5     6

	root := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}
	node4 := &Node{Value: 4}
	node5 := &Node{Value: 5}
	node6 := &Node{Value: 6}
	root.Children = []*Node{node2, node3}
	node2.Children = []*Node{node4, node5}
	node3.Children = []*Node{node6}

	result := BreadthFirstSearch(root, 5)
	assert.NotNil(t, result)
	assert.Equal(t, 5, result.Value)
}

func TestBreadthFirstSearch_ReturnsNilIfTargetNotFound(t *testing.T) {
	root := &Node{Value: 1}
	node2 := &Node{Value: 2}
	root.Children = []*Node{node2}

	result := BreadthFirstSearch(root, 99)
	assert.Nil(t, result)
}

func TestBreadthFirstSearch_ReturnsRootIfTargetIsRoot(t *testing.T) {
	root := &Node{Value: 42}
	result := BreadthFirstSearch(root, 42)
	assert.NotNil(t, result)
	assert.Equal(t, 42, result.Value)
}

func TestBreadthFirstSearch_ReturnsNilIfRootIsNil(t *testing.T) {
	result := BreadthFirstSearch(nil, 1)
	assert.Nil(t, result)
}

func TestBreadthFirstSearch_FindsTargetInWideTree(t *testing.T) {
	root := &Node{Value: 1}
	for i := 2; i <= 10; i++ {
		root.Children = append(root.Children, &Node{Value: i})
	}
	result := BreadthFirstSearch(root, 7)
	assert.NotNil(t, result)
	assert.Equal(t, 7, result.Value)
}
