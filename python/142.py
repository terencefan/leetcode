# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>

# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution(object):
    def detectCycle(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        l1, l2 = head, head
        a, v, i = None, None, 0
        while l1 and l2 and l1.next and l2.next.next:
            i += 1
            l1 = l1.next
            l2 = l2.next.next
            if l1.val == l2.val:
                if not a:
                    a = i
                    v = l1.val
                    continue
            if l1.val == v:
                loop = i - a
                print(a, loop)
                break
        return None


if __name__ == '__main__':
    from utils import ListNode
    x = ListNode(1)
    y = ListNode(5)
    z = ListNode(-3)
    a = ListNode(3)
    b = ListNode(2)
    c = ListNode(0)
    d = ListNode(-4)
    x.next = y
    y.next = z
    z.next = a
    a.next = b
    b.next = c
    c.next = d
    d.next = b
    print(a)
    c = Solution().detectCycle(x)
    print(c)
