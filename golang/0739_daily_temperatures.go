package golang

// dailyTemperatures 每日温度
// 解题思路：1.暴力法、2.单调栈
// 利用栈的特性后进先出特性，比较前后两个元素的大小，从而得出index的差值
func dailyTemperatures(T []int) []int {
    stack := []int{}
    ans := make([]int, len(T))

    for key, value := range T {
        if len(stack) == 0 {
            stack = append(stack, key)
            continue
        }
        for len(stack) > 0 && T[stack[len(stack)-1]] < value {
            prevIndex := stack[len(stack)-1]
            ans[prevIndex] = key - prevIndex
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, key)
    }

    return ans
}
