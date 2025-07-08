"""
80. Remove Duplicates from Sorted Array II
https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
Medium

[Description]
Given an integer array `nums` sorted in non-decreasing order, remove some
duplicates in-place such that each element appears at most twice. The relative
order of the elements should be kept the same.

[Result]
Date: 2025-07-08
Runtime: 73ms, Beats 97.05%
Memory: 20.37MB, Beats 79.71%
"""

from typing import List

class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        i = 0
        j = 0
        n = len(nums)
        while j < n:
            if j < 2 or nums[i - 2] < nums[j]:
                nums[i] = nums[j]
                i += 1
            j += 1
        return i
