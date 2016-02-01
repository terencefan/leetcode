#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

from utils import (
    Tree,
)

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):

    def chain(self, root):
        node, left, right = root, root.left, root.right
        if left:
            node.left = None
            node.right, node = self.chain(left)
        if right:
            node.right, node = self.chain(right)
        return root, node

    def flatten(self, root):
        """
        :type root: TreeNode
        :rtype: void Do not return anything, modify root in-place instead.
        """
        if not root:
            return
        root, _ = self.chain(root)
        print root


if __name__ == '__main__':
    Solution().flatten(Tree('1#23').root)
