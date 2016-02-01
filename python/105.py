#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

from utils import TreeNode


class Solution(object):

    # recursive solution.
    # def buildTree(self, preorder, inorder):
    #     if not preorder:
    #         return None

    #     root = TreeNode(preorder[0])
    #     pos = inorder.index(root.val)
    #     root.left = self.buildTree(preorder[1: pos + 1], inorder[:pos])
    #     root.right = self.buildTree(preorder[pos + 1:], inorder[pos + 1:])
    #     return root

    def buildTree(self, preorder, inorder):
        if not preorder:
            return None

        root = TreeNode(preorder[0])

        st, j = [root], 0
        for i in range(1, len(preorder)):
            parent, node = None, TreeNode(preorder[i])
            while st and st[-1].val == inorder[j]:
                parent = st.pop()
                j += 1
            if parent:
                parent.right = node
            else:
                st[-1].left = node
            st.append(node)
        return root


print Solution().buildTree([1, 2, 3], [2, 1, 3])
