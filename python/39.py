#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def combinationSum(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """

        candidates.sort()
        st, length, result = [(0, 0, [])], len(candidates), []

        while st:
            ptr, ans, res = st.pop()
            if ans > target:
                continue
            if ans == target:
                result.append(res)
            if ptr == length:
                continue
            for i in reversed(range(ptr, length)):
                num = candidates[i]
                st.append((i, ans + num, res + [num]))

        return result


if __name__ == '__main__':
    print Solution().combinationSum([2, 3, 6, 7], 7)
