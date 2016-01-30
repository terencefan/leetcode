#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

from collections import defaultdict


class Solution(object):
    def groupAnagrams(self, strs):
        """
        :type strs: List[str]
        :rtype: List[List[str]]
        """

        def _hash(s):
            return ''.join(sorted(s))

        res = defaultdict(list)
        for s in strs:
            res[_hash(s)].append(s)

        return [sorted(x) for x in res.values()]


print Solution().groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"])
