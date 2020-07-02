package golang

func isBalanced(root *TreeNode) bool {
    if maxDepthV2(root) == -1 {
        return false
    }
    return true
}

func maxDepthV2(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := maxDepthV2(root.Left)
    right := maxDepthV2(root.Right)
    if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
        return -1
    }
    if left > right {
        return left + 1
    }
    return right + 1
}
