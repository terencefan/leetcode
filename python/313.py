#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>


class Solution(object):
    def nthSuperUglyNumber(self, n, primes):
        """
        :type n: int
        :type primes: List[int]
        :rtype: int
        """
        r, index = [1], [0] * len(primes)
        for i in range(1, n):
            r.append(r[index[0]] * primes[0])
            ks = [0]
            for j in range(1, len(primes)):
                m = r[index[j]] * primes[j]
                if m < r[i]:
                    ks, r[i] = [j], m
                elif m == r[i]:
                    ks.append(j)
            for k in ks:
                index[k] += 1
        return r[n - 1]
