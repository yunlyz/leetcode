package golang

// TreeNode 二叉树
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
