public class Main {

    private static void s5() {
        String a = new S5.Solution().longestPalindrome("aaabaaaa");
        System.out.println(a);
    }

    private static void s647() {
        int n = new S647.Solution().countSubstrings("abc");
        System.out.println(n);
    }

    private static void s904() {
        int n = new S904.Solution().totalFruit(new int[]{1,2,1});
        System.out.println(n);
    }

    public static void main(String[] args) {
        s904();
    }
}
