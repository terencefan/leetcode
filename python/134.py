#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def canCompleteCircuit(self, gas, cost):
        """
        :type gas: List[int]
        :type cost: List[int]
        :rtype: int
        """

        if sum(gas) < sum(cost):
            return -1

        n = len(gas)
        diff = [gas[i] - cost[i] for i in range(n)]

        res, pos = 0, 0

        for i in range(n):
            res += diff[i]
            if res < 0:
                res = 0
                pos = i + 1
        return pos

if __name__ == '__main__':
    print Solution().canCompleteCircuit(
        [5, 7, 0, 10],
        [4, 1, 2, 3],
    )
