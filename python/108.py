#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

from utils import TreeNode


# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):

    def sortedArrayToBST(self, nums):
        """
        :type nums: List[int]
        :rtype: TreeNode
        """
        self.nums = nums
        return self.generate(0, len(nums) - 1)

    def generate(self, l, r):
        if l > r:
            return None
        mid = (l + r) / 2
        node = TreeNode(self.nums[mid])
        node.left = self.generate(l, mid - 1)
        node.right = self.generate(mid + 1, r)
        return node


print Solution().sortedArrayToBST([1, 2, 3])
