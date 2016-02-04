#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def isValidSerialization(self, preorder):
        """
        :type preorder: str
        :rtype: bool
        """
        preorder = preorder.split(',')
        if preorder[0] == '#':
            if len(preorder) == 1:
                return True
            return False

        count = 1
        for i in preorder:
            if not count:
                return False
            if i == '#':
                count -= 1
            else:
                count += 1
        return count == 0
