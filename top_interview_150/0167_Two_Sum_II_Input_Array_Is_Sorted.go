/*
167. Two Sum II - Input Array Is Sorted
https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
Medium

Given a 1-indexed array of integers `numbers` that is already sorted in
non-decreasing order, find two numbers such that they add up to a specific
target number. Let these two numbers be `numbers[index1]` and `numbers[index2]`
where `1 <= index1 < index2 <= numbers.length`.

Return the indices of the two numbers, `index1` and `index2`, added by one as an
integer array `[index1, index2]` of length 2.

The tests are generated such that there is exactly one solution. You may not use
the same element twice.

[Result]
Date: 2025-07-09
Runtime: 0ms, Beats 100.00%
Memory: 7.86MB, Beats 71.27%
*/

func twoSum(numbers []int, target int) []int {
    l := 0
    r := len(numbers) - 1

    for l < r {
        sum := numbers[l] + numbers[r]
        if sum == target {
            return []int{l + 1, r + 1}
        }
        if sum < target {
            l++
        } else {
            r--
        }
    }
    return []int{}
}
