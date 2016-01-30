#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def fullJustify(self, words, maxWidth):
        """
        :type words: List[str]
        :type maxWidth: int
        :rtype: List[str]
        """

        def _format(fragments, length):
            length, num = len(fragments), maxWidth - length
            if length == 1:
                return fragments[0].ljust(maxWidth)
            avg, left = divmod(num, length - 1)
            for i in range(length - 1):
                fragments[i] += ' ' * (avg + 1 if i < left else avg)
            return ''.join(fragments)

        res, length, fragments = [], len(words[0]), [words[0]]

        for word in words[1:]:
            if length + len(word) + len(fragments) > maxWidth:
                res.append(_format(fragments, length))
                length, fragments = len(word), [word]
            else:
                fragments.append(word)
                length += len(word)

        res.append(' '.join(fragments).ljust(maxWidth))
        return res


print Solution().fullJustify([
    "This", "is", "an", "example", "of", "text", "justification."
], 16)
