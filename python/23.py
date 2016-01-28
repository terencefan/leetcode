#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

import heapq


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):

    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """

        minHeap = []
        res = node = ListNode(-1)
        for i in range(len(lists)):
            if not lists[i]:
                continue
            val, lists[i] = lists[i].val, lists[i].next
            heapq.heappush(minHeap, (val, i))

        while minHeap:
            val, i = heapq.heappop(minHeap)
            node.next = ListNode(val)
            node = node.next
            if lists[i]:
                val, lists[i] = lists[i].val, lists[i].next
                heapq.heappush(minHeap, (val, i))

        return res.next


if __name__ == '__main__':
    print Solution().mergeKLists([[1, 3, 5], [2, 9, 10], [6, 11]])
