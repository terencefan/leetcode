#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def zigzagLevelOrder(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        if not root:
            return []

        res, nodes = [], [root]
        while True:
            temp, r = [], []
            for node in nodes:
                if not node:
                    continue
                r.append(node.val)
                temp += [node.left, node.right]
            if not r:
                break
            if len(res) % 2:
                r.reverse()
            res.append(r)
            nodes = temp
        return res


if __name__ == '__main__':
    from utils import Tree
    t = Tree('12345')
    print Solution().zigzagLevelOrder(t.root)
