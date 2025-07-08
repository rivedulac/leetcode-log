"""
26. Remove Duplicates from Sorted Array
https://leetcode.com/problems/remove-duplicates-from-sorted-array/
Easy

[Description]
Given an integer array `nums` sorted in non-decreasing order, remove the
duplicates in-place such that each unique element appears only once. The
relative order of the elements should be kept the same. Then return the number
of unique elements in `nums`.

[Result]
Date: 2025-07-08
Runtime: 0ms, Beats 100.00%
Memory: 18.89MB, Beats 74.45%
"""
from typing import List

class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        i, j = 0, 1
        n = len(nums)
        while j < n:
            while j < n and nums[i] == nums[j]:
                j += 1
            if j == n:
                break
            # print(i, j, nums)
            i += 1
            nums[i] = nums[j]
        return i + 1
