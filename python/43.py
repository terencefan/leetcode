#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def multiply(self, num1, num2):
        """
        :type num1: str
        :type num2: str
        :rtype: str
        """

        temp, maxlength = [], 0
        for index, i in enumerate(reversed(num2)):
            k, r = 0, [0] * index

            for j in reversed(num1):
                k, t = divmod(int(j) * int(i) + k, 10)
                r.append(t)
            if k:
                r.append(k)
            temp.append(r)
            maxlength = max(maxlength, len(r))

        res, k = [], 0
        for i in range(maxlength):
            r = sum(j[i] for j in temp if len(j) > i) + k
            k, t = divmod(r, 10)
            res.append(t)
        if k:
            res.append(k)
        res = ''.join(map(str, reversed(res)))
        while res.startswith('0') and res != '0':
            res = res[1:]
        return res


print Solution().multiply('123', '0')
