package golang

func levelOrderBottom(root *TreeNode) [][]int {
    ans := make([][]int, 0)

    if root == nil {
        return ans
    }

    stack := make([][]int, 0)
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)

    for len(queue) != 0 {
        list := make([]int, 0)
        tmp := make([]*TreeNode, 0)
        for _, node := range queue {
            list = append(list, node.Val)
            if node.Left != nil {
                tmp = append(tmp, node.Left)
            }
            if node.Right != nil {
                tmp = append(tmp, node.Right)
            }
        }
        stack = append(stack, list)
        queue = tmp
    }

    for i := len(stack) - 1; i >= 0; i-- {
        ans = append(ans, stack[i])
    }

    return ans
}
