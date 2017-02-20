#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>


class Solution(object):
    def candy(self, ratings):
        """
        :type ratings: List[int]
        :rtype: int
        """
        n = len(ratings)

        _as, _ds, res = 1, 1, 1
        for i in range(1, n):
            if ratings[i - 1] > ratings[i]:
                _ds += 1
                continue

            if _ds > 1:
                res += _ds * (_ds - 1) / 2
                if _ds > _as:
                    res += _ds - _as
                _as = _ds = 1

            if ratings[i - 1] < ratings[i]:
                _as += 1
                res += _as
            else:
                _as = _ds = 1
                res += 1

        res += _ds * (_ds - 1) / 2
        if _ds > _as:
            res += _ds - _as
        return res


if __name__ == '__main__':
    Solution().candy([1, 3, 4, 3, 2, 1])
    Solution().candy([4, 2, 3, 4, 1])
