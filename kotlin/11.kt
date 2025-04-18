package leetcode.p11

class Solution {
    fun maxArea(height: IntArray): Int {
        var i = 0
        var j = height.size - 1

        var maxArea = 0
        while (i < j) {
            maxArea = maxOf(maxArea, (j - i) * minOf(height[i], height[j]))
            if (height[i] < height[j]) {
                i++
            } else {
                j--
            }
        }
        return maxArea
    }
}