package leetcode.p155

import java.util.Stack

class MinStack() {

    private val s1 = Stack<Int>()

    private val s2 = Stack<Int>()

    fun push(v: Int) {
        s1.push(v)
        if (s2.isEmpty() || s2.peek() >= v) {
            s2.push(v)
        }
    }

    fun pop() {
        val v = s1.pop()
        if (s2.peek() == v) {
            s2.pop()
        }
    }

    fun top(): Int {
        return s1.peek()
    }

    fun getMin(): Int {
        return s2.peek()
    }

}

/**
 * Your MinStack object will be instantiated and called as such:
 * var obj = MinStack()
 * obj.push(`val`)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */