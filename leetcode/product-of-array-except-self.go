// https://leetcode.com/problems/product-of-array-except-self/
package main
func productExceptSelf(nums []int) []int {
    le := len(nums)
    ans := make([]int, le)
    ans[0] = 1
    for i := 1; i < le; i++ {
        ans[i] = ans[i - 1] * nums[i - 1]
    }
     r := 1
     for i := le - 1; i >= 0; i-- {
         ans[i] *= r
         r *= nums[i]
     }
     return ans
}