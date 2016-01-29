#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Validator(set):

    def add(self, x):
        if x == '.':
            return
        if x in self:
            raise Exception('num %s duplicated' % x)
        super(Validator, self).add(x)


class Solution(object):

    def isValidSudoku(self, board):
        try:
            self.validate(board)
            return True
        except:
            return False

    def validate(self, board):

        for i in xrange(9):
            row, column, block = Validator(), Validator(), Validator()
            rs, cs = divmod(i, 3)
            for j in xrange(9):
                row.add(board[i][j])
                column.add(board[j][i])
                block.add(board[rs * 3 + j / 3][cs * 3 + j % 3])


if __name__ == '__main__':
    board = [
        "618425379",
        "495637812",
        "327891645",
        "549713286",
        "276958431",
        "831246597",
        "164589723",
        "783162954",
        "952374168",
    ]
    print Solution().isValidSudoku(board)
