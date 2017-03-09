#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>


class Solution(object):

    def hIndex(self, citations):
        """
        :type citations: List[int]
        :rtype: int
        """
        citations.sort()
        n = len(citations)
        for i in range(1, n + 1, 1):
            if citations[n - i] < i:
                return i - 1
        return 0


a = Solution().hIndex([0, 1])
print(a)
