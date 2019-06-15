#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com>

MINUTE = 60
HOUR = 60 * MINUTE
DATE = 24 * HOUR

levels = {
    "Year": 0,
    "Month": 10,
    "Day": 20,
    "Hour": 30,
    "Minute": 40,
    "Second": 50,
}

monthdays = [0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
monthdays_integral = []

r = 0
for i in range(12):
    r += monthdays[i]
    monthdays_integral.append(r)


def gettimestamp(timestamp, gra):
    year, month, date, hour, minute, second = timestamp.split(":")
    year = int(year)
    month = int(month)
    date = int(date)
    hour = int(hour)
    minute = int(minute)
    second = int(second)

    r, level = 0, levels[gra]
    if level >= 0:
        r += year * 365 + (year - 1) / 4 if year > 2000 else 0
    if level >= 10:
        r += monthdays_integral[month - 1] + \
            (1 if year % 4 == 0 and month > 2 else 0)
    if level >= 20:
        r += date - 1
    r *= 3600 * 24

    if level >= 30:
        r += hour * 3600
    if level >= 40:
        r += minute * 60
    if level >= 50:
        r += second
    return r


def merge(A, B):
    i, j, r = 0, 0, []
    while i < len(A) and j < len(B):
        if A[i][0] < B[j][0]:
            r.append(A[i])
            i += 1
        else:
            r.append(B[j])
            j += 1
    while i < len(A):
        r.append(A[i])
        i += 1
    while j < len(B):
        r.append(B[j])
        j += 1
    return r


class LogSystem(object):

    def __init__(self):
        self.logs = []
        self.unsorted_logs = []

    def put(self, id, timestamp):
        """
        :type id: int
        :type timestamp: str
        :rtype: None
        """
        t = gettimestamp(timestamp, "Second")
        self.unsorted_logs.append((t, id))

    def retrieve(self, s, e, gra):
        """
        :type s: str
        :type e: str
        :type gra: str
        :rtype: List[int]
        """
        if len(self.unsorted_logs) > 0:
            self.unsorted_logs.sort()
            self.logs = merge(self.logs, self.unsorted_logs)
            self.unsorted_logs = []

        def find(timestamp):
            lo, hi = 0, len(self.logs)
            while lo < hi:
                mid = lo + (hi - lo) / 2
                if self.logs[mid][0] >= timestamp:
                    hi = mid
                else:
                    lo = mid + 1
            return lo

        p = gettimestamp(s, gra)

        q = gettimestamp(e, gra)
        if gra == "Second":
            q += 1
        elif gra == "Minute":
            q += MINUTE
        elif gra == "Hour":
            q += HOUR
        elif gra == "Day":
            q += DATE
        elif gra == "Month":
            month = int(e.split(":")[1])
            q += monthdays[month] * DATE
        elif gra == "Year":
            year = int(e.split(":")[0])
            q += (366 if year % 4 == 0 else 365) * DATE

        l, r = find(p), find(q)
        return [x[1] for x in self.logs[l:r]]


# Your LogSystem object will be instantiated and called as such:
# obj = LogSystem()
# obj.put(id,timestamp)
# param_2 = obj.retrieve(s,e,gra)
