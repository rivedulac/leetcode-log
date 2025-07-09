/*
26. Remove Duplicates from Sorted Array
https://leetcode.com/problems/remove-duplicates-from-sorted-array/
Easy

[Description]
Given an integer array `nums` sorted in non-decreasing order, remove the
duplicates in-place such that each unique element appears only once. The
relative order of the elements should be kept the same. Then return the number
of unique elements in `nums`.

[Result]
Date: 2025-07-09
Runtime: 0ms, Beats 100.00%
Memory: 6.26MB, Beats 60.38%
*/

func removeDuplicates(nums []int) int {
    i := 0
    for j := 0; j < len(nums); j++ {
        if nums[j] != nums[i] {
            i++
            nums[i] = nums[j] 
        }
    }
    return i + 1
}
