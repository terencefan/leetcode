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

    def isSymmetric(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        if not root:
            return True

        l, r = root.left, root.right
        st = [(l, r)]

        while st:
            l, r = st.pop()
            if l is None and r is None:
                continue
            elif l is None or r is None:
                return False
            if l.val != r.val:
                return False
            st.append((l.left, r.right))
            st.append((l.right, r.left))
        return True


if __name__ == '__main__':
    from utils import Tree
    t = Tree('541#1#42#2#')
    print Solution().isSymmetric(t.root)
