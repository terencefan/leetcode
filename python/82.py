#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def deleteDuplicates(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        dummy = pre = ListNode(0)
        dummy.next = head
        while head and head.next:
            if head.val != head.next.val:
                pre = pre.next
                head = head.next
                continue

            while head and head.next and head.val == head.next.val:
                head = head.next
            pre.next = head.next
            head = head.next
        return dummy.next
