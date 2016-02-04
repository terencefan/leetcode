#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    steps = [
        (8, 0b00000000111111110000000011111111, 0b11111111000000001111111100000000),
        (4, 0b00001111000011110000111100001111, 0b11110000111100001111000011110000),
        (2, 0b00110011001100110011001100110011, 0b11001100110011001100110011001100),
        (1, 0b01010101010101010101010101010101, 0b10101010101010101010101010101010)
    ]

    def reverseBits(self, n):
        n = (n >> 16) | (n << 16)
        for bit, m1, m2 in self.steps:
            n = ((n >> bit) & m1) | ((n << bit) & m2)
        return n

    def reverseBits2(self, n):
        res = 0
        for i in range(32):
            res = res << 1
            res += n & 1
            n = n >> 1
        return res


print Solution().reverseBits(43261596)
