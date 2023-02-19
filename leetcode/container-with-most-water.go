// https://leetcode.com/problems/container-with-most-water/
package main
func maxArea(height []int) int {
    ans, lo, hi := 0, 0, len(height) - 1
    for lo < hi {
        val := (hi - lo) * min(height[lo], height[hi])
        ans = max(ans, val)
        if height[lo] < height[hi] {
            lo += 1
        } else {
            hi -= 1
        }
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
