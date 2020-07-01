package golang

// climbStairs 爬楼梯
// 解题思路：1.动态规划
//  ①动态转移方程式：f(x) = f(x-1) + f(x-2)
//  ②边界条件：f(0) = 1、f(1) = 1、f(2) = f(1) + f(0) == 2
func climbStairs(n int) int {
    ans := make([]int, n+1)
    ans[0] = 1
    ans[1] = 1
    for i := 2; i <= n; i++ {
        ans[i] = ans[i-1] + ans[i-2]
    }
    return ans[n]
}