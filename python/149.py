#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

from collections import defaultdict


class Point(object):
    def __init__(self, a=0, b=0):
        self.x = a
        self.y = b


class Solution(object):

    def maxPoints(self, points):
        """
        :type points: List[Point]
        :rtype: int
        """
        n, res = len(points), 0
        for i in range(n):
            hashmap, vertical, same, t = defaultdict(int), 0, 0, 0
            for j in range(i + 1, n):
                pi, pj = points[i], points[j]
                if pi.x == pj.x:
                    if pi.y == pj.y:
                        same += 1
                    else:
                        vertical += 1
                else:
                    k = round(float(pi.y - pj.y) / (pi.x - pj.x), 2)
                    hashmap[k] += 1
                    t = max(t, hashmap[k])
            res = max(res, max(t, vertical) + 1 + same)
        return res


if __name__ == '__main__':
    print Solution().maxPoints([
        Point(1, 1),
        Point(2, 2),
        Point(2, 2),
        Point(1, 1),
        Point(1, 0),
        Point(1, 0),
        Point(1, 0),
    ])
