package golang

import (
    "math"
)

func maxPathSum(root *TreeNode) int {
    var pathMax = math.MinInt32
    var helper func(root *TreeNode) int
    helper = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        left := max(helper(root.Left), 0)
        right := max(helper(root.Right), 0)
        pathMax = max(pathMax, root.Val+left+right)
        return root.Val + max(left, right)
    }
    helper(root)
    return pathMax
}
