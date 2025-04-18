package leetcode.p205

import kotlin.collections.mutableMapOf

class Solution {
    fun isIsomorphic(s: String, t: String): Boolean {
        if (s.length != t.length) return false

        var m = mutableMapOf<Char, Char>()

        s.forEachIndexed { i, c ->
            if (m.containsKey(c)) {
                if (m[c] != t[i]) return false
            } else {
                if (m.containsValue(t[i])) return false
                m[c] = t[i]
            }
        }
        return true
    }
}