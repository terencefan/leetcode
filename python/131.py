#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>


class Solution(object):

    def partition(self, s):
        """
        :type s: str
        :rtype: List[List[str]]
        """
        n = len(s)
        a = [[0] * n for i in range(n)]
        for k in range(n):
            for i in range(0, n - k):
                j = i + k
                if i == j:
                    a[i][j] = 1
                    continue
                if s[i] != s[j]:
                    continue
                if i + 1 == j:
                    a[i][j] = 1
                else:
                    a[i][j] = a[i + 1][j - 1]

        st, res = [([], 0)], []
        while st:
            b, i = st.pop()
            if i == n:
                res.append(b)
            for j in range(i, n):
                if a[i][j]:
                    st.append((b + [s[i:j + 1]], j + 1))
        return res


res = Solution().partition('efe')
print(res)
