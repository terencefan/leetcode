#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def longestValidParentheses(self, s):
        """
        :type s: str
        :rtype: int
        """

        st, length = [], len(s)
        for i in range(length):
            if st and s[st[-1]] == '(' and s[i] == ')':
                st.pop()
            else:
                st.append(i)
        return max(map(lambda a, b: b - a - 1, [-1] + st, st + [length]))


print Solution().longestValidParentheses('()))()()()()))()')
