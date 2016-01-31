#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

from utils import (
    ListNode,
    List,
)


# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):

    def reverseBetween(self, head, m, n):
        """
        :type head: ListNode
        :type m: int
        :type n: int
        :rtype: ListNode
        """
        pre = dummy = ListNode(0)
        dummy.next = head

        i = 1
        while i < m:
            pre, head = head, head.next
            i += 1
        bl, br = pre, head

        while i <= n:
            nxt = head.next
            head.next = pre
            pre, head = head, nxt
            i += 1
        bl.next = pre
        br.next = head

        return dummy.next


if __name__ == '__main__':
    print Solution().reverseBetween(List([1, 2, 3, 4, 5]).head, 4, 5)
