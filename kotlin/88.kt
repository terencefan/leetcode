package leetcode.p88

class Solution {
    fun merge(nums1: IntArray, m: Int, nums2: IntArray, n: Int): Unit {
        for (i in m - 1 downTo 0) {
            nums1[i + n] = nums1[i]
        }
        var i = m
        var j = 0
        for (k in 0 until m + n) {
            if (j >= n) {
                nums1[k] = nums1[i++]
            } else if (i >= m + n) {
                nums1[k] = nums2[j++]
            } else if (nums1[i] < nums2[j]) {
                nums1[k] = nums1[i++]
            } else {
                nums1[k] = nums2[j++]
            }
        }
    }
}