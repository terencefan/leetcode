#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

from collections import deque


class Solution(object):
    def combinationSum(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """

        candidates.sort()
        st, length, result = deque([(0, 0, [])]), len(candidates), []

        while st:
            ptr, ans, res = st.popleft()
            if ans == target:
                result.append(res)
            if ptr == length:
                continue
            for i in range(ptr, length):
                if i > ptr and candidates[i] == candidates[i - 1]:
                    continue
                num = candidates[i]
                if ans + num > target:
                    break
                st.append((i + 1, ans + num, res + [num]))

        return result


if __name__ == '__main__':
    print Solution().combinationSum([10, 1, 2, 7, 6, 1, 5], 8)
