package golang

// postorderTraversal 二叉树后序遍历
// 非递归实现：遍历左子节点，直到叶子节点；
// 如果是页面节点，必然会访问；
// 并且上次访问的节点和当前节点的右子节点一样时，表示右子节点已经访问完，则访问该节点值并且出栈
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    ans := make([]int, 0)
    stack := make([]*TreeNode, 0)
    var lastVisit *TreeNode
    for root != nil || len(stack) != 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        node := stack[len(stack)-1]
        if node.Right == nil || node.Right == lastVisit {
            stack = stack[:len(stack)-1]
            ans = append(ans, node.Val)
            lastVisit = node
        } else {
            root = node.Right
        }
    }

    return ans
}
