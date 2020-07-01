package golang

import (
    "sort"
)

// threeSum 三数之和
// 解题思路：1.排序+双指针
// 边界条件: 1. 当前值大于0，则三数之和不可能为零; 2. 当前值与前一个索引值相等，表示会重复，跳过
//      3. res = cur + left + right == 0，如果left或right的下一个值相等，表示会重复
func threeSum(nums []int) [][]int {
    ans := [][]int{}
    sort.Ints(nums)
    var left, right int
    for i := 0; i < len(nums)-1; i++ {
        left = i + 1
        right = len(nums) - 1
        if nums[i] > 0 {
            return ans
        }
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        for right > left {
            zero := nums[i] + nums[left] + nums[right]
            if zero == 0 {
                ans = append(ans, []int{nums[i], nums[left], nums[right]})
                if right > left && nums[left] == nums[left+1] {
                    left++
                }
                for right > left && nums[right] == nums[right-1] {
                    right--
                }
                left++
                right--
            } else if zero > 0 {
                right--
            } else if zero < 0 {
                left++
            }
        }
    }
    return ans
}
