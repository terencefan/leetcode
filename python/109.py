#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

from utils import (  # noqa
    List,
    ListNode,
    TreeNode,
)

# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):

    def sortedListToBST(self, head):
        """
        :type head: ListNode
        :rtype: TreeNode
        """
        size, node = 0, head
        while node:
            size += 1
            node = node.next
        return self.generate(head, size)

    def generate(self, head, size):
        if not head or not size:
            return None

        i = j = head
        n = 1
        while j and j.next and n < size:
            i = i.next
            j = j.next.next
            n += 2
        node = TreeNode(i.val)
        node.left = self.generate(head, size / 2)
        node.right = self.generate(i.next, size - size / 2 - 1)
        return node


print Solution().sortedListToBST(List([1, 2, 3]).head)
