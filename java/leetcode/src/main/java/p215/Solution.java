package p215;

import java.util.PriorityQueue;
import java.util.Queue;

class Solution {
    public int findKthLargest(int[] nums, int k) {
        if (nums.length == 0 || k == 0) {
            return 0;
        }

        Queue<Integer> q = new PriorityQueue<>();
        for (int num : nums) {
            q.add(num);
            if (q.size() > k) {
                q.poll();
            }
        }

        return q.poll();
    }
}
