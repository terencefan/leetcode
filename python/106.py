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

    def buildTree(self, inorder, postorder):
        if not postorder:
            return None

        root = TreeNode(postorder[-1])

        st, j = [root], len(inorder) - 1
        for i in range(len(postorder) - 2, -1, -1):
            parent, node = None, TreeNode(postorder[i])
            while st and st[-1].val == inorder[j]:
                parent = st.pop()
                j -= 1
            if parent:
                parent.left = node
            else:
                st[-1].right = node
            st.append(node)
        return root


print Solution().buildTree([-1], [-1])
