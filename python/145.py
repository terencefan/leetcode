#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    # recursive solution.
    # def postorderTraversal(self, root):
    #     if not root:
    #         return []
    #     return self.postorderTraversal(root.left) + \
    #         self.postorderTraversal(root.right) + \
    #         [root.val]

    # iterative with 2 stacks
    # def postorderTraversal(self, root):
    #     if not root:
    #         return []

    #     q, st = [root], []
    #     while q:
    #         node = q.pop()
    #         if node.left:
    #             q.append(node.left)
    #         if node.right:
    #             q.append(node.right)
    #         st.append(node)

    #     res = []
    #     while st:
    #         res.append(st.pop().val)
    #     return res

    # iterative with 1 stack and flag
    def postorderTraversal(self, root):

        node, st = root, []
        while node or st:

            while node:
                st.append((node, False))
                node = node.left

            if st:
                node, flag = st.pop()
                if flag:
                    print node.val
                    node = None
                else:
                    st.append((node, True))
                    node = node.right


if __name__ == '__main__':
    from utils import Tree
    print Solution().postorderTraversal(Tree('132##54').root)
