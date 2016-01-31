#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def isSameTree(self, p, q):
        """
        :type p: TreeNode
        :type q: TreeNode
        :rtype: bool
        """
        if p is None:
            return q is None
        elif q is None:
            return p is None
        elif p.val != q.val:
            return False
        return self.isSameTree(p.left, q.left) and \
            self.isSameTree(p.right, q.right)
