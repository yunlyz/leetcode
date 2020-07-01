package golang

import (
    "sort"
)

func findBestValue(arr []int, target int) int {
    sort.Ints(arr)
    sum := 0
    avg := target / len(arr)
    for key, value := range arr {
        if avg < value {
            t := float64((target - sum) / (len(arr) - key))
            if (t - float64(avg)) > 0.5 {
                return int(t) + 1
            }
            return int(t)
        }
        sum += value
    }

    return arr[len(arr)-1]
}
