#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com>


def kmp(s, p):

    def build(p):
        f = [0] * len(p)

        k = 0
        for i in range(1, len(p)):
            k = k + 1 if p[i] == p[k] else 0
            f[i] = k
        return f

    f = build(p)

    i, j = 0, 0
    while i < len(s):
        if j == len(p):
            return i - len(p)
        if s[i] == p[j]:
            i += 1
            j += 1
        else:
            j = f[j - 1] if j > 0 else 0

    return i - len(p) if j == len(p) else -1


if __name__ == '__main__':
    r = kmp("abcdabcdabx", "abcdabx")
    print(r)
