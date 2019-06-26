#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com>


def find_less_than(arr, k):
    lo, hi = 1, len(arr)
    while lo < hi:
        mid = lo + (hi - lo) / 2
        if arr[mid] >= k:
            hi = mid
        else:
            lo = mid + 1
    return arr[lo - 1]


def find_less_equal(arr, k):
    lo, hi = 1, len(arr)
    while lo < hi:
        mid = lo + (hi - lo) / 2
        if arr[mid] > k:
            hi = mid
        else:
            lo = mid + 1
    return arr[lo - 1]


def find_more_than(arr, k):
    lo, hi = 0, len(arr) - 1
    while lo < hi:
        mid = lo + (hi - lo) / 2
        if arr[mid] > k:
            hi = mid
        else:
            lo = mid + 1
    return arr[lo]


def find_more_equal(arr, k):
    lo, hi = 0, len(arr) - 1
    while lo < hi:
        mid = lo + (hi - lo) / 2
        if arr[mid] >= k:
            hi = mid
        else:
            lo = mid + 1
    return arr[lo]


if __name__ == '__main__':
    r = find_more_equal([0, 1, 3, 7, 9, 10], 7)
    print(r)
