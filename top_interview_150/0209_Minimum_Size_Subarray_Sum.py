"""
209. Minimum Size Subarray Sum
https://leetcode.com/problems/minimum-size-subarray-sum/
Medium

[Description]
Given an array of positive integers `nums` and a positive integer `target`,
return the minimal length of a subarray whose sum is greater than or equal to
`target`. If there is no such subarray, return `0` instead.

[Result]
Date: 2025-07-10
Runtime: 11ms, Beats 46.59%
Memory: 28.28MB, Beats 80.66%
"""

from typing import List

class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        if not nums or sum(nums) < target:
            return 0
        
        n = len(nums)
        minlen = n
        numsum = nums[0]
        l = 0
        r = 0
        
        while r < n:
            if minlen == 1:
                break
            # When the sum of subarray is smaller than target,
            # add one more element to the subarray.
            if numsum < target:
                if r == n - 1:
                    break
                r += 1
                numsum += nums[r]
            # When the sum of subarray meets the condition,
            # remove the leftmost element from the subarray.
            else:
                minlen = min(minlen, r - l + 1)
                numsum -= nums[l]
                l += 1
        return minlen
