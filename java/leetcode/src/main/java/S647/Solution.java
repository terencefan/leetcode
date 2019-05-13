package S647;

public class Solution {
    private char[] s;
    private int[] len;

    private final char SPLITTER = '|';

    private void expand(String str) {
        int l = str.length() * 2 + 1;
        s = new char[l];
        len = new int[l];

        for (int i = 0; i < str.length(); i++) {
            s[2 * i] = SPLITTER;
            s[2 * i + 1] = str.charAt(i);
        }
        s[l - 1] = SPLITTER;
    }

    private void calculate() {
        int maxRight = 0;
        int mid = 0;
        int l = len.length;

        for (int i = 0; i < len.length; i++) {
            int j;
            if (i < maxRight) {
                j = Math.min(maxRight - i, len[mid - (i - mid)]);
            } else {
                j = 1;
            }
            for (; j <= Math.min(i, l - i - 1); j++) {
                if (s[i - j] != s[i + j]) {
                    break;
                }
            }
            len[i] = j;

            if (i + j - 1 > maxRight) {
                maxRight = i + j - 1;
                mid = i;
            }
        }
    }

    public int countSubstrings(String str) {
        expand(str);
        calculate();

        int n = 0;
        for (int i = 0; i < len.length; i++) {
            n += len[i] / 2;
        }
        return n;
    }
}