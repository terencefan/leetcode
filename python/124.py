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

    def maxPathSum(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        self._res = root.val
        self._maxPath(root)
        return self._res

    # a: res
    # b: longest sum from root
    def _maxPath(self, root):
        if root is None:
            return 0
        left = self._maxPath(root.left)
        right = self._maxPath(root.right)
        left = 0 if left < 0 else left
        right = 0 if right < 0 else right
        self._res = max(self._res, left + right + root.val)
        return max(left, right) + root.val


if __name__ == '__main__':
    from utils import Tree
    print Solution().maxPathSum(Tree([-1, '#', 9, -6, 3, '#', '#', '#', '#', -2]).root)
