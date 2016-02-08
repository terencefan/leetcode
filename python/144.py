#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    # recursive solution.
    # def preorderTraversal(self, root):

    #     if not root:
    #         return []
    #     return [root.val] + \
    #         self.preorderTraversal(root.left) + \
    #         self.preorderTraversal(root.right)

    # iterative solution.
    def preorderTraversal(self, root):
        if not root:
            return []

        node, st, res = root, [], []
        while node:
            while node:
                res.append(node.val)
                if node.left:
                    st.append(node)
                node = node.left
            if st:
                node = st.pop()
                node = node.right
        return res


if __name__ == '__main__':
    from utils import Tree
    print Solution().preorderTraversal(Tree('132##54').root)
