"""
15. 3Sum
https://leetcode.com/problems/3sum/
Medium

Given an integer array `nums`, return all the triplets `[nums[i], nums[j], nums[k]]`
such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`.

Notice that the solution set must not contain duplicate triplets.

[Result]
Date: 2025-07-09
Runtime: 419ms, Beats 94.62%
Memory: 20.72MB, Beats 36.73%
"""

from typing import List

class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        zero = 0
        pos, neg = list(), list()
        output = set()

        for num in nums:
            if not num:
                zero += 1
            elif num > 0:
                pos.append(num)
            else:
                neg.append(num)
        
        pos_set = set(pos)
        neg_set = set(neg)
        
        # Three zeros
        if zero >= 3:
            output.add((0, 0, 0))
        
        # One zero, one negative, one positive
        if zero:
            for num in pos_set:
                if num * -1 in neg_set:
                    output.add((num * -1, 0, num))
        
        # No zero, two negatives, one positive
        for i in range(len(neg)):
            for j in range(i + 1, len(neg)):
                target = -1 * (neg[i] + neg[j])
                if target in pos_set:
                    output.add(tuple(sorted([neg[i], neg[j], target])))
        
        # No zero, one negative, two positives
        for i in range(len(pos)):
            for j in range(i + 1, len(pos)):
                target = -1 * (pos[i] + pos[j])
                if target in neg_set:
                    output.add(tuple(sorted([target, pos[i], pos[j]])))
        return list(output)
