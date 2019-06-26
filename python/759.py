#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com>

import heapq

INF = 1 << 32 - 1


class Solution(object):

    def employeeFreeTime(self, schedule):
        """
        Using heap.
        """

        h = []
        for i, employee in enumerate(schedule):
            if len(employee) > 0:
                heapq.heappush(h, (employee[0], i, 0))

        head, tail, worktime = -INF, -INF, []
        while len(h) > 0:
            (start, end), id, index = heapq.heappop(h)

            if index < len(schedule[id]) - 1:
                heapq.heappush(h, (schedule[id][index + 1], id, index + 1))

            if start > tail:
                if head != tail:
                    worktime.append([head, tail])
                head, tail = start, end
                continue

            tail = max(tail, end)

        if head != tail:
            worktime.append([head, tail])

        r = []
        for i in range(1, len(worktime)):
            r.append([worktime[i - 1][1], worktime[i][0]])
        return r

    def employeeFreeTime_dac(self, schedule):
        """
        Divide & Conquer
        """

        def union(a, b):
            if len(a) == 0:
                return b
            elif len(b) == 0:
                return a

            c = a + b
            c.sort()

            head, tail, r = -INF, -INF, []
            for start, end in c:
                if start > tail:
                    if head != tail:
                        r.append([head, tail])
                    head, tail = start, end
                    continue
                tail = max(tail, end)

            if head != tail:
                r.append([head, tail])

            return r

        def find(schedule):
            if len(schedule) == 1:
                return schedule[0]
            elif len(schedule) == 2:
                return union(schedule[0], schedule[1])

            mid = len(schedule) / 2
            return union(find(schedule[:mid]), find(schedule[mid:]))

        worktime = find(schedule)
        if len(worktime) < 2:
            return []

        r = []
        for i in range(1, len(worktime)):
            r.append([worktime[i - 1][1], worktime[i][0]])
        return r
