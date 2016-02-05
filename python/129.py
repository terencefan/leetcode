#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def sumNumbers(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if not root:
            return 0

        res, st = 0, [(root, 0)]
        while st:
            node, val = st.pop()
            val = val * 10 + node.val
            if not node.left and not node.right:
                res += val
            if node.left:
                st.append((node.left, val))
            if node.right:
                st.append((node.right, val))
        return res


if __name__ == '__main__':
    from utils import Tree
    print Solution().sumNumbers(Tree('123').root)
