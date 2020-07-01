package golang

import (
    "strconv"
)

func addBinary(a string, b string) string {
    var ans = ""
    var carry = 0
    length := max(len(a), len(b))
    for i := 0; i < length; i++ {
        if i < len(a) {
            carry += int(a[len(a)-i-1] - '0')
        }
        if i < len(b) {
            carry += int(b[len(b)-i-1] - '0')
        }
        ans = strconv.Itoa(carry%2) + ans
        carry /= 2
    }
    if carry > 0 {
        ans = "1" + ans
    }

    return ans
}
