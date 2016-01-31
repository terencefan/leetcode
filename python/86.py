#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

from utils import (
    List,
    ListNode,
)


class Solution(object):
    def partition(self, head, x):
        """
        :type head: ListNode
        :type x: int
        :rtype: ListNode
        """

        dummy1 = node1 = ListNode(0)
        dummy2 = node2 = ListNode(0)
        while head:
            if head.val < x:
                node1.next = ListNode(head.val)
                node1 = node1.next
            else:
                node2.next = ListNode(head.val)
                node2 = node2.next
            head = head.next

        if not dummy1.next:
            return dummy2.next
        else:
            node1.next = dummy2.next
            return dummy1.next


if __name__ == '__main__':
    l = List([1, 4, 3, 2, 5, 2])
    print Solution().partition(l.head, 3)
