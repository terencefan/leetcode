package leetcode.p6

class Solution {
    fun convert(s: String, numRows: Int): String {
        if (numRows == 1) return s

        var sb = StringBuilder()
        val k = 2 * numRows - 2

        for (i in 0 until numRows) {
            for (j in i until s.length step k) {
                sb.append(s[j])
                if (i != 0 && i != numRows - 1) {
                    val mirror = j + k - 2 * i
                    if (mirror < s.length) {
                        sb.append(s[mirror])
                    }
                }
            }
        }
        return sb.toString()
    }
}
