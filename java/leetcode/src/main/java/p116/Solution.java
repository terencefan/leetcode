package p116;

import java.util.ArrayDeque;
import java.util.Deque;

class Solution {


    public Node connect(Node root) {
        if (root == null) {
            return null;
        }

        Deque<Node> queue = new ArrayDeque<>();
        queue.push(root);

        while (queue.size() > 0) {
            Node pre = null;
            int length = queue.size();
            for (int i = 0; i < length; i++) {
                Node node = queue.pollLast();
                if (node == null) {
                    break;
                }
                if (pre != null) {
                    pre.next = node;
                }
                if (node.left != null) {
                    queue.push(node.left);
                }
                if (node.right != null) {
                    queue.push(node.right);
                }
                pre = node;
            }
        }

        return root;
    }
}