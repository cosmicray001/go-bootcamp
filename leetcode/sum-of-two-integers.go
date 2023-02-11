// https://leetcode.com/problems/sum-of-two-integers/
package main
func getSum(a int, b int) int {
    for b != 0 {
        c := a & b
        a = a ^ b
        b = c << 1
    }
    return a
}
