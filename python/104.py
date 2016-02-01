#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):

    # recursive solution, 2%...
    # def maxDepth(self, root):
    #     """
    #     :type root: TreeNode
    #     :rtype: int
    #     """
    #     if not root:
    #         return 0
    #     l = self.maxDepth(root.left)
    #     r = self.maxDepth(root.right)
    #     return max(l, r) + 1

    # recurent solution, 17%
    # def maxDepth(self, root):
    #     a, depth = [root], 0
    #     while a:
    #         b, depth = [], depth + 1
    #         for node in a:
    #             if node.left:
    #                 b.append(node.left)
    #             if node.right:
    #                 b.append(node.right)
    #         a = b
    #     return depth

    def maxDepth(self, root):
        if not root:
            return 0

        depth, st = 0, [(root, 1)]
        while st:
            node, d = st.pop()
            if node.left:
                st.append((node.left, d + 1))
            if node.right:
                st.append((node.right, d + 1))
            depth = max(depth, d)
        return depth
