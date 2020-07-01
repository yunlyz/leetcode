package golang

import (
    "strings"
)

// isPalindromeStr 验证回文串
// 解题发放：1.双指针法
// 解题思路：利用双指针，判断前后是否问数字或字母，如果不是则移动指针。
func isPalindromeStr(s string) bool {
    s = strings.ToLower(s)
    left, right := 0, len(s)-1
    for right > left {
        for right > left && !isLetterAndNum(s[left]) {
            left++
        }
        for right > left && !isLetterAndNum(s[right]) {
            right--
        }
        if right > left {
            if s[left] != s[right] {
                return false
            }
            left++
            right--
        }
    }
    return true
}

func isLetterAndNum(ch byte) bool {
    return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
