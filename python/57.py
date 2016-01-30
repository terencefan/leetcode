#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Interval(object):
    def __init__(self, s=0, e=0):
        self.start = s
        self.end = e

    def __repr__(self):
        return '(%d - %d)' % (self.start, self.end)


class Solution(object):
    def insert(self, intervals, newInterval):
        """
        :type intervals: List[Interval]
        :rtype: List[Interval]
        """
        st = []
        intervals.append(newInterval)
        for interval in sorted(intervals, key=lambda x: x.start):
            if st and interval.start <= st[-1].end:
                st[-1].end = max(st[-1].end, interval.end)
            else:
                st.append(interval)
        return st


if __name__ == '__main__':
    print Solution().insert([
        Interval(4, 9),
        Interval(1, 5),
        Interval(10, 16),
    ], Interval(8, 17))
