package p449;

public class Main {
    public static void main(String[] args) {
        TreeNode node = new TreeNode(1);
        node.left = new TreeNode(2);
        node.right = new TreeNode(9);
        node.right.right = new TreeNode(7);
        node.left.left = new TreeNode(3);
        node.left.right = new TreeNode(5);

        Codec codec = new Codec();
        String s = codec.serialize(node);
        node = codec.deserialize(s);
        System.out.println("hi");
    }
}
