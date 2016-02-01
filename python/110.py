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
    def isBalanced(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        res, _ = self.check(root)
        return res

    def check(self, root):
        if not root:
            return True, 0
        l, ld = self.check(root.left)
        r, rd = self.check(root.right)
        return l and r and abs(ld - rd) < 2, max(ld, rd) + 1


if __name__ == '__main__':
    from utils import Tree
    print Solution().isBalanced(Tree('1223333444444##55').root)
