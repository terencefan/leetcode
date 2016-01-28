#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

MAX_INT = 2147483647


class Solution:
    # @param {integer} dividend
    # @param {integer} divisor
    # @return {integer}

    def divide(self, dividend, divisor):

        if not divisor:
            return MAX_INT
        if not dividend:
            return 0

        operator = (dividend > 0) is (divisor > 0)
        dividend = abs(dividend)
        divisor = abs(divisor)

        result, step, scale = 0, 1, divisor
        while scale < dividend:
            step = step << 1
            scale = scale << 1

        while scale > 0:
            if dividend >= scale:
                dividend -= scale
                result += step
            scale = scale >> 1
            step = step >> 1

        return min(result if operator else -result, MAX_INT)

print Solution().divide(53, -2)
