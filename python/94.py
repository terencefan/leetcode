#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

from utils import (
    Tree,
)


class Solution(object):

    def inorderTraversal(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        if not root:
            return []
        l = self.inorderTraversal(root.left)
        r = self.inorderTraversal(root.right)
        return l + [root.val] + r


if __name__ == '__main__':
    t = Tree('123##4##5')
    print t.root.right
    print Solution().inorderTraversal(t.root)
