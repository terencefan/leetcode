#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution(object):

    def hasCycle(self, head):
        """
        :type head: ListNode
        :rtype: bool
        """

        if not head:
            return False

        i = j = head
        while j.next and j.next.next:
            i = i.next
            j = j.next.next
            if i.val == j.val:
                return True
        return False


if __name__ == '__main__':
    from utils import List
    print Solution().hasCycle(List([1, 2, 3, 4]).head)
