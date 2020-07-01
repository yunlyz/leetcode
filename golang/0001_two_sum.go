package golang

// twoSum 两数之和
// 解题思路：1.暴力循环法、2.两次hash法、3.一次hash法
// 一次hash法，利用hash表特性，key存在值，value存储index
// 判断已经遍历的列表中是否存在和当前值相加等于target的值
// 如果存在返回prev index和当前index，否则返回nil
func twoSum(nums []int, target int) []int {
    hash := map[int]int{}
    for key, value := range nums {
        if index, ok := hash[target-value]; ok == true {
            return []int{index, key}
        }
        hash[value] = key
    }
    return nil
}
