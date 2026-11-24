package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// AreMirrors reports whether two trees are mirror images of each other:
// equal roots, a's left mirroring b's right and a's right mirroring b's left.
func AreMirrors(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.Val == b.Val && AreMirrors(a.Left, b.Right) && AreMirrors(a.Right, b.Left)
}

// Mirror transforms a tree into its mirror image in place (swap children
// everywhere) and returns it — the constructive counterpart of the check.
func Mirror(n *TreeNode) *TreeNode {
	if n == nil {
		return nil
	}
	n.Left, n.Right = Mirror(n.Right), Mirror(n.Left)
	return n
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	a := tn(1, tn(2, nil, nil), tn(3, nil, nil))
	b := tn(1, tn(3, nil, nil), tn(2, nil, nil))
	fmt.Println(AreMirrors(a, b))
}
