#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def levelOrder(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        if not root:
            return []

        res, nodes = [], [root]
        while nodes:
            temp, r = [], []
            for node in nodes:
                r.append(node.val)
                if node.left:
                    temp.append(node.left)
                if node.right:
                    temp.append(node.right)
            res.append(r)
            nodes = temp
        return res


if __name__ == '__main__':
    from utils import Tree
    t = Tree('1234##5')
    print Solution().levelOrder(t.root)
