package p297;

public class P297 {
    public static void main(String args[]) {
        TreeNode n1 = new TreeNode(-1);
        n1.left = new TreeNode(0);
        n1.right = new TreeNode(1);

        String s = new Codec().serialize(n1);
        TreeNode n = new Codec().deserialize(s);
        System.out.println(n);
    }
}
