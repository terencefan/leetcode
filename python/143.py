#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>

from utils import ListNode


class Solution(object):

    def reorderList(self, head):
        """
        :type head: ListNode
        :rtype: void Do not return anything, modify head in-place instead.
        """
        if not head or not head.next:
            return head

        l1 = l2 = head
        while l2.next and l2.next.next:
            l1 = l1.next
            l2 = l2.next.next
        l2 = l1.next
        l1.next = None
        l1 = None

        while l2:
            t = l2.next
            l2.next = l1
            l1 = l2
            l2 = t
        l2 = l1
        l1 = head

        while l2:
            t = l2.next
            l2.next = l1.next
            l1.next = l2
            l1 = l1.next.next
            l2 = t


head = ListNode(1)
head.next = ListNode(2)
head.next.next = ListNode(3)
# head.next.next.next = ListNode(4)
Solution().reorderList(head)
while head:
    print(head)
    head = head.next
