// https://leetcode.com/problems/two-sum/
package main
func twoSum(nums []int, target int) []int {
    mp := make(map[int]int)
    for idx, val := range nums {
        x := target - val
        i, j := mp[x]
        if j {
            return []int{i, idx}
        }
        mp[val] = idx
    }
    return []int{}
}