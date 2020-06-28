package golang

// inorderTraversal 二叉树中序遍历
// 同样利用栈后进先出特性，记录每个访问节点历史。中止条件为：当前节点为空则弹出中序节点，同时遍历右子节点
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    ans := make([]int, 0)
    stack := make([]*TreeNode, 0)

    for root != nil || len(stack) != 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        ans = append(ans, node.Val)
        root = node.Right
    }

    return ans
}
