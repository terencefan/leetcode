package p430;

import java.util.ArrayDeque;
import java.util.Deque;
import java.util.Map;

class Solution {
    public Node flatten(Node head) {
        if (head == null) {
            return null;
        }

        Deque<Node> stack = new ArrayDeque<>();
        stack.push(head);

        Node dummy = new Node();
        Node pre = dummy;

        while (stack.size() > 0) {
            Node node = stack.pop();

            node.prev = pre;
            pre.next = node;

            while (node != null) {
                if (node.child != null) {
                    if (node.next != null) {
                        stack.push(node.next);
                    }
                    node.next = node.child;
                    node.next.prev = node;
                    node.child = null;
                }

                pre = node;
                node = node.next;
            }
        }

        dummy.next.prev = null;
        return dummy.next;
    }
}