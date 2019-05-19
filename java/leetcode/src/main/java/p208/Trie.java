package p208;

class Trie {

    class Node {
        boolean end;
        Node[] arr;

        Node() {
            this.end = false;
            this.arr = new Node[26];
        }
    }

    private Node root = new Node();

    /** Initialize your data structure here. */
    public Trie() {
    }

    /** Inserts a word into the trie. */
    public void insert(String word) {
        Node node = root;
        for (char c : word.toCharArray()) {
            int index = c - 'a';
            if (node.arr[index] == null) {
                node.arr[index] = new Node();
            }
            node = node.arr[index];
        }
        node.end = true;
    }

    boolean search(String word, boolean end) {
        Node node = root;
        for (char c : word.toCharArray()) {
            int index = c - 'a';
            if (node.arr[index] == null) {
                return false;
            } else {
                node = node.arr[index];
            }
        }
        return end ? node.end == true : true;
    }

    /** Returns if the word is in the trie. */
    public boolean search(String word) {
        return search(word, true);
    }

    /** Returns if there is any word in the trie that starts with the given prefix. */
    public boolean startsWith(String prefix) {
        return search(prefix, false);
    }
}
