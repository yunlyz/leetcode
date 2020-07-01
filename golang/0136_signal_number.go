package golang

// singleNumber 只出现一个的数字
// 解题思路：1.hash表、2.加减法、3.异或运算、4.队列
func singleNumber(nums []int) int {
    signle := 0
    for _, num := range nums {
        signle ^= num
    }
    return signle
}
