package golang

// maxScoreSightseeingPair 最佳观光组合
// 解题思路：枚举法，遍历找到和为最大的组合
// 优化：利用一次遍历，找到当前index和index之前的最大值之和。
// 利用两个标记位，一个标记index之前的最大值，另一个标记当前index和index的和的最大值。
func maxScoreSightseeingPair(A []int) int {
    var ans = 0
    var mx = A[0] - 0
    for j := 1; j < len(A); j++ {
        mx = max(mx, A[j-1]+j-1)
        ans = max(ans, mx+A[j]-j)
    }
    return ans
}
