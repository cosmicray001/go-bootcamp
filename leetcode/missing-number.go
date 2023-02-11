// https://leetcode.com/problems/missing-number/
package main

func missingNumber(nums []int) int {
    ans := 0
    for idx, val := range nums {
        ans ^= idx
        ans ^= val
    }
    return ans ^ len(nums)
}
