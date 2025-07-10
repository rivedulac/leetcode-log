/*
209. Minimum Size Subarray Sum
https://leetcode.com/problems/minimum-size-subarray-sum/
Medium

[Description]
Given an array of positive integers `nums` and a positive integer `target`,
return the minimal length of a subarray whose sum is greater than or equal to
`target`. If there is no such subarray, return `0` instead.

[Result]
Date: 2025-07-10
Runtime: 0ms, Beats 100.00%
Memory: 9.35MB, Beats 65.16%
*/

package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
    l, r := 0, 0
    numsum := nums[0]
    n := len(nums)
    minlen := n + 1
    for r < n {
        if numsum < target {
            if r == n - 1 {
                break
            }
            r++
            numsum += nums[r]
        } else {
            newlen := r - l + 1
            if minlen > newlen {
                minlen = newlen
            }
            numsum -= nums[l]
            l++
        }
    }
    if minlen == n + 1 {
        return 0
    }
    return minlen
}

func main() {
    target := 7
    nums := []int{2, 3, 1, 2, 4, 3}
    fmt.Println(minSubArrayLen(target, nums))
}
