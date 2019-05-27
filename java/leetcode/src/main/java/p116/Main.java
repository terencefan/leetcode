package p116;

public class Main {

    public static void main(String[] args) {
        Node n2 = new Node(2, null, null, null);
        Node n3 = new Node(3, null, null, null);
        Node n1 = new Node(3, n2, n3, null);
        new Solution().connect(n1);
    }

}
