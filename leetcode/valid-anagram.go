// https://leetcode.com/problems/valid-anagram/
package main
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    mp := make(map[string]int)
    for i := 0; i < len(s); i++ {
        mp[string(s[i])] += 1
    }
    for i := 0; i < len(t); i++ {
        mp[string(t[i])]--
        if mp[string(t[i])] < 0 {
            return false
        }
    }
    return true
}
