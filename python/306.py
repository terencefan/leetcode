#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def _is(self, num, astr, bstr):

        if astr != '0' and astr[0] == '0':
            return False
        if bstr != '0' and bstr[0] == '0':
            return False

        c = int(astr) + int(bstr)
        cstr = str(c)
        cstrl = len(cstr)

        if num[:cstrl] == cstr:
            if cstrl == len(num):
                return True
            return self._is(num[cstrl:], bstr, cstr)

    def isAdditiveNumber(self, num):
        l = len(num)

        if l < 3:
            return False

        for i in range(1, l - 1):
            for j in range(i + 1, l):
                astr, bstr = num[:i], num[i:j]
                if self._is(num[j:], astr, bstr):
                    return True
        return False


if __name__ == '__main__':
    print Solution().isAdditiveNumber('11234')
