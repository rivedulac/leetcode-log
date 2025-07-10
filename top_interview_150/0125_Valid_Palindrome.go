/*
125. Valid Palindrome
https://leetcode.com/problems/valid-palindrome/
Easy

A phrase is a palindrome if, after converting all uppercase letters into
lowercase letters and removing all non-alphanumeric characters, it reads the
same forward and backward. Alphanumeric characters include letters and numbers.

Given a string `s`, return `true` if it is a palindrome, or `false` otherwise.

[Result]
Date: 2025-07-09
Runtime: 0ms, Beats 100.00%
Memory: 4.62MB, Beats 75.56%
*/

package main

import (
    "fmt"
    "strings"
)

func isAlphaNumeric(r byte) bool {
    return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func isPalindrome(s string) bool {
    l := 0
    r := len(s) - 1

    for l < r {
        for l < r && !isAlphaNumeric(s[l]) {
            l++
        }
        for l < r && !isAlphaNumeric(s[r]) {
            r--
        }
        if l == r {
            break
        }
        if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
            return false
        }
        l++
        r--
    }
    return true
}

func main() {
    s := "A man, a plan, a canal: Panama"
    fmt.Println(isPalindrome(s))
}
