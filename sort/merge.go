package sort

func mergeSort(nums []int) []int {
    ans := make([]int, 0)
    if len(nums) <= 1 {
        return nums
    }
    if len(nums) == 2 {

    }

    left := mergeSort(nums[:len(nums)/2])
    right := mergeSort(nums[len(nums)/2:])
    
    ans = append(ans, left...)
    ans = append(ans, right...)

    return ans
}
