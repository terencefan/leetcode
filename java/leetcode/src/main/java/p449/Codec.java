package p449;

import java.nio.ByteBuffer;
import java.util.LinkedList;

public class Codec {

    private static ByteBuffer buf = ByteBuffer.allocate(4);

    private void serialize(StringBuilder sb, TreeNode node) {
        if (node == null) {
            sb.append('$');
        } else {
            buf.putInt(0, node.val);
            sb.append(new String(buf.array()));
        }
    }

    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {
        StringBuilder sb = new StringBuilder();

        LinkedList<TreeNode> q = new LinkedList<>();
        q.add(root);

        while (q.size() > 0) {
            int length = q.size();
            for (int i = 0; i < length; i++) {
                TreeNode node = q.pollFirst();
                serialize(sb, node);
                if (node == null) {
                    continue;
                }
                q.add(node.left);
                q.add(node.right);
            }
        }
        return sb.toString();
    }

    private TreeNode deserialize(ByteBuffer sb) {
        if (sb.get() == '$') {
            return null;
        }
        sb.position(sb.position() - 1);
        return new TreeNode(sb.getInt());
    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {

        ByteBuffer buf = ByteBuffer.wrap(data.getBytes());
        if (buf.get() == '$') {
            return null;
        }
        buf.position(buf.position() - 1);

        TreeNode root = new TreeNode(buf.getInt());
        LinkedList<TreeNode> q = new LinkedList<>();
        q.push(root);

        while (q.size() > 0) {
            int length = q.size();
            for (int i = 0; i < length; i++) {
                TreeNode node = q.pollFirst();
                node.left = deserialize(buf);
                if (node.left != null) {
                    q.add(node.left);
                }
                node.right = deserialize(buf);
                if (node.right != null) {
                    q.add(node.right);
                }
            }
        }
        return root;
    }
}
