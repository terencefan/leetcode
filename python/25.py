#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

    def __repr__(self):
        return 'ListNode {}'.format(self.val)


class Solution(object):

    def reverseKGroup(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """

        pre, node, t = None, head, [None] * k
        while True:

            for i in range(k):
                if not node:
                    return head
                t[i] = node
                node = node.next

            print t, pre

            if not pre:
                head = t[k - 1]
            else:
                pre.next = t[k - 1]

            for i in range(k - 1, 0, -1):
                t[i].next = t[i - 1]
            t[0].next = node
            pre = t[0]

        return head


head = ListNode(1)
head.next = ListNode(2)
head.next.next = ListNode(3)
head.next.next.next = ListNode(4)

head = Solution().reverseKGroup(head, 2)
while head:
    print head.val
    head = head.next
