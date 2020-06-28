package golang

// preorderTraversal 二叉树的前序遍历
// 非递归解题思路：利用栈后进先出特性，上一次访问的节点在栈的最顶层。
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    ans := make([]int, 0)
    stack := make([]*TreeNode, 0)

    for root != nil || len(stack) != 0 {
        for root != nil {
            ans = append(ans, root.Val)
            stack = append(stack, root)
            root = root.Left
        }
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        root = node.Right
    }

    return ans
}
