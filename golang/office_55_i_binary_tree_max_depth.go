package golang

func maxDepth(root *TreeNode) int {
    ans, left, right := 0, 0, 0
    if root == nil {
        return 0
    }
    ans += 1
    if root.Left != nil {
        left = maxDepth(root.Left)
    }
    if root.Right != nil {
        right = maxDepth(root.Right)
    }
    ans += max(left, right)
    return ans
}
