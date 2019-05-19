package p297;

import java.nio.ByteBuffer;

public class Codec {

    private void serialize(StringBuilder buf, TreeNode node) {
        if (node == null) {
            buf.append((byte) 0);
            return;
        }
        buf.append((byte) 1);
        buf.append(node.val);
        buf.append('\0');
        serialize(buf, node.left);
        serialize(buf, node.right);
    }

    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {
        StringBuilder buf = new StringBuilder();
        serialize(buf, root);
        return buf.toString();
    }

    private TreeNode deserialize(ByteBuffer buf) {
        if (buf.get() == '0') {
            return null;
        }
        int val = 0;
        boolean minus = false;

        while (true) {
            byte b = buf.get();
            if (b == '-') {
                minus = true;
            } else if (b == '\0') {
                break;
            } else {
                val *= 10;
                val += b - '0';
            }
        }
        TreeNode node = new TreeNode(minus ? -val : val);
        node.left = deserialize(buf);
        node.right = deserialize(buf);
        return node;
    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        ByteBuffer buf = ByteBuffer.wrap(data.getBytes());
        return deserialize(buf);
    }
}
