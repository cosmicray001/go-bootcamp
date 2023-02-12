// https://leetcode.com/problems/maximum-subarray/
package main

func maxSubArray(nums []int) int {
    ans, mx := nums[0], nums[0]

    for i := 1; i < len(nums); i++ {
        mx = max(nums[i] + mx, nums[i])
        ans = max(ans, mx)
    }
    return ans
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}
