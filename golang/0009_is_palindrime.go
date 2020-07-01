package golang

// isPalindrome 回文数
// 解题思路：1.逆转数字比较两个数字是否相等
// 逆转思路：采用模10 + 值乘以10的算法将数字逆转
// 我们只要逆转一半的数字就可以判断是否为回文数，当数字个数为奇数时，逆转数字除以10后在判断相等
func isPalindrome(x int) bool {
    if x < 0 || (x%10 == 0 && x != 0) {
        return false
    }
    reverse := 0
    for x > reverse {
        reverse = reverse*10 + x%10
        x /= 10
    }

    return x == reverse || reverse/10 == x
}
