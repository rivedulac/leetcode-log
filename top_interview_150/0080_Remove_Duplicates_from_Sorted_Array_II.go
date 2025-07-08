"""
80. Remove Duplicates from Sorted Array II
https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
Medium

[Description]
Given an integer array `nums` sorted in non-decreasing order, remove some
duplicates in-place such that each element appears at most twice. The relative
order of the elements should be kept the same.

[Result]
Date: 2025-07-09
Runtime: 3ms, Beats 96.84%
Memory: 7.30MB, Beats 34.39%
"""

func removeDuplicates(nums []int) int {
    i := 0

    for j:= 0; j < len(nums); j++ {
        if j < 2 || nums[j] > nums[i - 2] {
            nums[i] = nums[j]
            i += 1
        }
    }
    return i
}
