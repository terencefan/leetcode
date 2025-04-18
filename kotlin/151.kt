package leetcode.p151

class Solution {
    fun reverseWords(s: String): String {
        return s.split(" ").filter({ it.isNotEmpty() }).reversed().joinToString(" ");
    }
}