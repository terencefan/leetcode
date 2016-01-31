#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

from utils import (
    Tree,
)


class Solution(object):

    def isValidBST(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        pre, p, st = None, root, []

        while p or st:
            while p:
                st.append(p)
                p = p.left
            if st:
                p = st.pop()
                if pre and p.val <= pre:
                    return False
                pre = p.val
                p = p.right
        return True


if __name__ == '__main__':
    t = Tree([10, 5, 15, '#', '#', 6, 20])
    print Solution().isValidBST(t.root)
