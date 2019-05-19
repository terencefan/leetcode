package p904;

public class Solution {
    public int totalFruit(int[] tree) {
        if (tree.length <= 1) {
            return 1;
        }

        int baskets[] = new int[2];
        baskets[0] = tree[0];
        baskets[1] = -1;

        int count = 1;
        int count_b = 0;
        int r = 0, t;

        int i;
        for (i = 1; i < tree.length; i++) {
            if (tree[i] != baskets[0]) {
                baskets[1] = tree[i];
                break;
            } else {
                count++;
            }
        }
        if (i == tree.length || i == tree.length - 1) {
            return tree.length;
        }

        for (;i < tree.length;i++) {
            count++;
            if (tree[i] == baskets[1]) {
                count_b++;
            } else if (tree[i] == baskets[0]) {
                t = baskets[0];
                baskets[0] = baskets[1];
                baskets[1] = t;
                count_b = 1;
            } else {
                count = count_b + 1;
                count_b = 1;
                baskets[0] = baskets[1];
                baskets[1] = tree[i];
            }
            r = r > count ? r : count;
        }
        return r;
    }
}
