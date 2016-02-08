#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

from collections import defaultdict


class Solution(object):

    def wordBreak(self, s, wordDict):
        """
        :type s: str
        :type wordDict: Set[str]
        :rtype: bool
        """

        n, wordMap = len(s), defaultdict(set)
        for word in wordDict:
            wordMap[len(word)].add(word)

        f = [True] + [False] * n
        for i in range(n):
            for l, words in wordMap.items():
                if i - l + 1 < 0 or not f[i - l + 1]:
                    continue
                t = s[i - l + 1:i + 1]
                for word in words:
                    if t == word:
                        f[i + 1] = True
                        break
        return f[n]


print Solution().wordBreak('dogs', set(['dog', 's', 'gs']))
