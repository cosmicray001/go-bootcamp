// https://leetcode.com/problems/contains-duplicate/
package main
func containsDuplicate(nums []int) bool {
    mp := make(map[int]bool)
    for _, val := range nums {
        if mp[val] {
            return true
        }
        mp[val] = true
    }
    return false
}