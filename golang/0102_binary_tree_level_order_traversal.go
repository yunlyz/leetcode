package golang

// levelOrder 二叉树层次遍历
// 利用计算器计算每层的长度，然后根据长度遍历当前层节点
// 双队列，一个全局队列和一个局部队列
func levelOrder(root *TreeNode) [][]int {
    ans := make([][]int, 0)

    if root == nil {
        return nil
    }

    queue := make([]*TreeNode, 0)
    queue = append(queue, root)

    for len(queue) > 0 {
        tmp := make([]*TreeNode, 0)
        list := make([]int, 0)
        for _, node := range queue {
            list = append(list, node.Val)
            if node.Left != nil {
                tmp = append(tmp, node.Left)
            }
            if node.Right != nil {
                tmp = append(tmp, node.Right)
            }
        }
        ans = append(ans, list)
        queue = tmp
    }

    return ans
}
