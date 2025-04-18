package leetcode.p2

import leetcode.common.ListNode

class Solution {
    fun addTwoNumbers(l1: ListNode?, l2: ListNode?): ListNode? {
        var i = l1
        var j = l2
        var carry = 0

        var dummy = ListNode(0)
        var prev = dummy

        while (i != null || j != null || carry > 0) {
            val node = ListNode((i?.`val` ?: 0) + (j?.`val` ?: 0) + carry)
            carry = node.`val` / 10
            node.`val` %= 10
            prev.next = node
            prev = prev.next!!
            i = i?.next
            j = j?.next
        }

        return dummy.next
    }
}
