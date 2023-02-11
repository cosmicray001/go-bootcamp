// https://leetcode.com/problems/counting-bits/
package main
func countBits(n int) []int {
    ans := make([]int, n + 1)
    ans[0] = 0
    for i := 1; i < n + 1; i++ {
        ans[i] = ans[i >> 1] + (i & 1)
    }
    return ans
}
