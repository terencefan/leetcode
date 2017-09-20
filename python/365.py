#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>


class Solution(object):

    def do_actions(self, a, b):
        # fill a
        if a < self.x:
            yield self.x, b
        # fill b
        if b < self.y:
            yield a, self.y
        # empty a
        if a > 0:
            yield 0, b
        # empty b
        if b > 0:
            yield a, 0
        # pour into a from b
        d = min(b, self.x - a)
        if d > 0:
            yield a + d, b - d
        # pour into b from a
        d = min(a, self.y - b)
        if d > 0:
            yield a - d, b + d

    def canMeasureWater_dfs(self, x, y, z):
        """
        :type x: int
        :type y: int
        :type z: int
        :rtype: bool
        """

        if z == 0:
            return True

        self.x = x
        self.y = y

        m = set()
        s = [(0, 0)]
        while s:
            a, b = s.pop()
            m.add((a, b))
            for p, q in self.do_actions(a, b):
                if (p, q) in m:
                    continue
                if p == z or q == z or p + q == z:
                    return True
                s.append((p, q))
        return False

    def canMeasureWater(self, x, y, z):

        if x + y < z:
            return False
        if x == z or y == z or x + y == z:
            return True

        if x < y:
            x, y = y, x
        while x % y:
            x, y = y, x % y
        return z % y == 0

r = Solution().canMeasureWater(104659, 104677, 142424)
print(r)
