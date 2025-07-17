package techniques

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDepthFirstSearchUsingRecursion_FindsTargetInTree(t *testing.T) {
	root := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}
	node4 := &Node{Value: 4}
	node5 := &Node{Value: 5}
	root.Children = []*Node{node2, node3}
	node2.Children = []*Node{node4, node5}

	result := DepthFirstSearchUsingRecursion(root, 5)
	assert.NotNil(t, result)
	assert.Equal(t, 5, result.Value)
}

func TestDepthFirstSearchUsingRecursion_ReturnsNilIfTargetNotFound(t *testing.T) {
	root := &Node{Value: 1}
	node2 := &Node{Value: 2}
	root.Children = []*Node{node2}

	result := DepthFirstSearchUsingRecursion(root, 99)
	assert.Nil(t, result)
}

func TestDepthFirstSearchUsingRecursion_ReturnsRootIfTargetIsRoot(t *testing.T) {
	root := &Node{Value: 42}
	result := DepthFirstSearchUsingRecursion(root, 42)
	assert.NotNil(t, result)
	assert.Equal(t, 42, result.Value)
}

func TestDepthFirstSearchUsingRecursion_ReturnsNilIfRootIsNil(t *testing.T) {
	result := DepthFirstSearchUsingRecursion(nil, 1)
	assert.Nil(t, result)
}

func TestDepthFirstSearchUsingRecursion_FindsTargetInDeepTree(t *testing.T) {
	root := &Node{Value: 1}
	current := root
	for i := 2; i <= 10; i++ {
		child := &Node{Value: i}
		current.Children = []*Node{child}
		current = child
	}
	result := DepthFirstSearchUsingRecursion(root, 10)
	assert.NotNil(t, result)
	assert.Equal(t, 10, result.Value)
}

func TestDepthFirstSearchUsingStack_FindsTargetInTree(t *testing.T) {
	root := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}
	node4 := &Node{Value: 4}
	node5 := &Node{Value: 5}
	root.Children = []*Node{node2, node3}
	node2.Children = []*Node{node4, node5}

	result := DepthFirstSearchUsingStack(root, 5)
	assert.NotNil(t, result)
	assert.Equal(t, 5, result.Value)
}

func TestDepthFirstSearchUsingStack_ReturnsNilIfTargetNotFound(t *testing.T) {
	root := &Node{Value: 1}
	node2 := &Node{Value: 2}
	root.Children = []*Node{node2}

	result := DepthFirstSearchUsingStack(root, 99)
	assert.Nil(t, result)
}

func TestDepthFirstSearchUsingStack_ReturnsRootIfTargetIsRoot(t *testing.T) {
	root := &Node{Value: 42}
	result := DepthFirstSearchUsingStack(root, 42)
	assert.NotNil(t, result)
	assert.Equal(t, 42, result.Value)
}

func TestDepthFirstSearchUsingStack_ReturnsNilIfRootIsNil(t *testing.T) {
	result := DepthFirstSearchUsingStack(nil, 1)
	assert.Nil(t, result)
}

func TestDepthFirstSearchUsingStack_FindsTargetInDeepTree(t *testing.T) {
	root := &Node{Value: 1}
	current := root
	for i := 2; i <= 10; i++ {
		child := &Node{Value: i}
		current.Children = []*Node{child}
		current = child
	}
	result := DepthFirstSearchUsingStack(root, 10)
	assert.NotNil(t, result)
	assert.Equal(t, 10, result.Value)
}
