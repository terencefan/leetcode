#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution:
    # @param {string} s the IP string
    # @return {string[]} All possible valid IP addresses

    def validate(self, s):
        if s[0] == '0' and s != '0':
            return False
        return 0 <= int(s) <= 255

    def restore(self, s, n=4):

        if n == 1:
            if self.validate(s):
                yield [s]
            return
        n -= 1

        p = max(0, len(s) - n * 3 - 1)
        q = min(3, len(s) - n)
        for i in range(p, q):
            ip = s[:i + 1]
            if self.validate(ip):
                for left in self.restore(s[i + 1:], n):
                    yield [ip] + left
        return

    def restoreIpAddresses(self, s):
        return ['.'.join(x) for x in self.restore(s)]


print Solution().restoreIpAddresses('10008')
