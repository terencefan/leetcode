#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com>


class Node(object):
    def __init__(self, val, neighbors):
        self.val = val
        self.neighbors = neighbors


def clone(node):
    return Node(node.val, [])


"""
# Definition for a Node.
class Node(object):
    def __init__(self, val, neighbors):
        self.val = val
        self.neighbors = neighbors
"""


class Solution(object):
    def cloneGraph(self, node):
        """
        :type node: Node
        :rtype: Node
        """
        if node is None:
            return None

        root = clone(node)
        q = [(node, root)]

        visited = {id(node): root}
        while len(q) > 0:
            m, n = q.pop()
            for old in m.neighbors:
                if id(old) not in visited:
                    visited[id(old)] = clone(old)
                    q.append((old, visited[id(old)]))
                n.neighbors.append(visited[id(old)])
        return root
