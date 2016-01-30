#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def simplifyPath(self, path):
        """
        :type path: str
        :rtype: str
        """

        st = []
        for i in path.split('/'):
            if not i or i == '.':
                continue
            elif i == '..':
                if st:
                    st.pop()
            else:
                st.append(i)
        return "/%s" % '/'.join(st)


print Solution().simplifyPath('/a/./b/../../c/..')
