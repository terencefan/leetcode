package leetcode.p150

import java.util.Stack

class Solution {
    fun evalRPN(tokens: Array<String>): Int {
        val s = Stack<Int>()

        tokens.forEach { token ->
            when (token) {
                "+", "-", "*", "/" -> {
                    val b = s.pop()
                    val a = s.pop()
                    when (token) {
                        "+" -> s.push(a + b)
                        "-" -> s.push(a - b)
                        "*" -> s.push(a * b)
                        "/" -> s.push(a / b)
                    }
                }
                else -> s.push(token.toInt())
            }
        }
        return s.pop()
    }
}