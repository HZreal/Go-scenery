package main

import "fmt"

// 二叉树可以使用链表实现，也可以使用数组实现，但数组实现一般用来表示完全二叉树

// TreeNode 使用链表实现二叉树
type TreeNode struct {
	Left  *TreeNode
	Data  string
	Right *TreeNode
}

// CompleteBinaryTree 使用数组实现完全二叉树。
// type CompleteBinaryTree []string

func NewTreeNode(value string) *TreeNode {
	return &TreeNode{Data: value}
}

func MockOneTree() *TreeNode {
	root := NewTreeNode("root")
	root.Left = NewTreeNode("rootLeft")
	root.Right = NewTreeNode("rootRight")
	root.Left.Left = NewTreeNode("rootLeftLeft")
	root.Left.Right = NewTreeNode("rootLeftRight")
	root.Right.Left = NewTreeNode("rootRightLeft")
	root.Right.Right = NewTreeNode("rootRightRight")

	//                         root
	//                      /        \
	//              rootLeft          rootRight
	//            /      \               /     \
	// rootLeftLeft rootLeftRight rootRightLeft rootRightRight

	return root
}

// DFS 深度优先搜索 Depth-First Search。深度优先搜索会从根节点开始，尽可能深地搜索树的分支，直到到达叶子节点，然后回溯到之前的节点，以探索下一个分支。
// 有三种主要的遍历策略：先序遍历、中序遍历、和后序遍历

// 先序遍历：先访问根节点，再访问左子树，最后访问右子树。
func preOrderTraversal(tree *TreeNode) {
	if tree == nil {
		return
	}

	fmt.Println("tree.data -->", tree.Data)
	preOrderTraversal(tree.Left)
	preOrderTraversal(tree.Right)
}

// 中序遍历：先访问左子树，再访问根节点，最后访问右子树。
func midOrderTraversal(tree *TreeNode) {
	if tree == nil {
		return
	}

	midOrderTraversal(tree.Left)
	fmt.Println("tree.data -->", tree.Data)
	midOrderTraversal(tree.Right)
}

// 后序遍历：先访问左子树，再访问右子树，最后访问根节点。
func postOrderTraversal(tree *TreeNode) {
	if tree == nil {
		return
	}

	postOrderTraversal(tree.Left)
	postOrderTraversal(tree.Right)
	fmt.Println("tree.data -->", tree.Data)
}

// LevelOrderTraversal 层次遍历
// 也称广度优先搜索Breadth-First Search(BFS)，每一层从左到右访问每一个节点。
func LevelOrderTraversal() {

}

func test() {
	tree := MockOneTree()
	preOrderTraversal(tree)
	fmt.Println("--------------------------------")
	midOrderTraversal(tree)
	fmt.Println("--------------------------------")
	postOrderTraversal(tree)
	fmt.Println("--------------------------------")
	LevelOrderTraversal()
}

func main() {
	test()
}
