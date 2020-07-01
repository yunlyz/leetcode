package golang

// longestCommonPrefix 最长公共前缀
// 解题思路：1.横向扫描、2.纵向扫描、3.分治
// 纵向查找解法：以第一个字符串为基准，以及比较所有字符串的第一个字符到最后一个字符
//          比较过程中，字符串长度小于当前索引值，该字符串为最长公共前缀，不会再有比这个更长的前缀
//          匹配到不相等的索引值，返回前面str[:index]
func longestCommonPrefix(strs []string) string {
    if len(strs) <= 0 {
        return ""
    }
    for index, char := range strs[0] {
        for _, str := range strs {
            if index >= len(str) {
                return str
            }
            if byte(char) != str[index] {
                return str[:index]
            }
        }
    }
    return strs[0]
}
