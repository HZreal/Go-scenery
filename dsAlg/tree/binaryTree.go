package main

import (
	"container/list"
	"fmt"
)

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
	return &TreeNode{
		Left:  nil,
		Data:  value,
		Right: nil,
	}
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
func preOrderTraversal(tree *TreeNode, result *[]string) {
	if tree == nil {
		return
	}

	*result = append(*result, tree.Data)
	preOrderTraversal(tree.Left, result)
	preOrderTraversal(tree.Right, result)

}

// 中序遍历：先访问左子树，再访问根节点，最后访问右子树。
func midOrderTraversal(tree *TreeNode, result *[]string) {
	if tree == nil {
		return
	}

	preOrderTraversal(tree.Left, result)
	*result = append(*result, tree.Data)
	preOrderTraversal(tree.Right, result)
}

// 后序遍历：先访问左子树，再访问右子树，最后访问根节点。
func postOrderTraversal(tree *TreeNode, result *[]string) {
	if tree == nil {
		return
	}

	preOrderTraversal(tree.Left, result)
	preOrderTraversal(tree.Right, result)
	*result = append(*result, tree.Data)
}

// LevelOrderTraversal 层次遍历
// 也称广度优先搜索Breadth-First Search(BFS)，每一层从左到右访问每一个节点。
func LevelOrderTraversal(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	// 初始化一个队列
	queue := list.New()
	queue.PushBack(root)
	var result []string

	for queue.Len() > 0 {
		// 取出队列头部的节点
		element := queue.Front()
		node := element.Value.(*TreeNode)
		queue.Remove(element)

		// 将当前节点的数据添加到结果数组中
		result = append(result, node.Data)

		// 将左子节点加入队列
		if node.Left != nil {
			queue.PushBack(node.Left)
		}

		// 将右子节点加入队列
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}

	return result
}

func test() {
	tree := MockOneTree()

	var preOrderResult []string
	preOrderTraversal(tree, &preOrderResult)
	fmt.Println("preOrderResult--------------------------------", preOrderResult)

	var midOrderResult []string
	midOrderTraversal(tree, &midOrderResult)
	fmt.Println("midOrderResult--------------------------------", midOrderResult)

	var postOrderResult []string
	postOrderTraversal(tree, &postOrderResult)
	fmt.Println("postOrderResult--------------------------------", postOrderResult)

	LevelOrderResult := LevelOrderTraversal(tree)
	fmt.Println("LevelOrderResult--------------------------------", LevelOrderResult)

}

func main() {
	test()
}
