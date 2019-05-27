package p341;

import java.util.ArrayList;
import java.util.List;

public class Main {

    public static void main(String[] args) {
        List<NestedInteger> l3 = new ArrayList<>();

        List<NestedInteger> l2 = new ArrayList<>();
        l2.add(new NestedIntegerItem(l3));

        List<NestedInteger> l1 = new ArrayList<>();
        l1.add(new NestedIntegerItem(l2));
        l1.add(new NestedIntegerItem(3));

        NestedIterator it = new NestedIterator(l1);
        System.out.println(it.hasNext());
        System.out.println(it.next());
        System.out.println(it.hasNext());
    }

}
