package p341;

import java.util.ArrayDeque;
import java.util.Deque;
import java.util.Iterator;
import java.util.List;


public class NestedIterator implements Iterator<Integer> {

    private Deque<List<NestedInteger>> stack;
    private Deque<Integer> stackIndex;

    NestedInteger current;

    public NestedIterator(List<NestedInteger> nestedList) {
        stack = new ArrayDeque<>();
        stackIndex = new ArrayDeque<>();
        stack.push(nestedList);
        stackIndex.push(0);
    }

    private void adjust() {
        while (stack.size() > 0) {
            List<NestedInteger> l = stack.peek();
            int index = stackIndex.pop();

            if (index >= l.size()) {
                stack.pop();
            } else {
                NestedInteger it = l.get(index);
                if (it.isInteger()) {
                    stackIndex.push(index);
                    current = it;
                    return;
                } else {
                    stackIndex.push(index + 1);
                    stack.push(it.getList());
                    stackIndex.push(0);
                }
            }
        }
    }

    @Override
    public Integer next() {
        if (current == null) {
            adjust();
        }
        int r = current.getInteger();
        stackIndex.push(stackIndex.pop() + 1);
        current = null;
        return r;
    }

    @Override
    public boolean hasNext() {
        adjust();
        return stack.size() > 0;
    }
}

/**
 * Your NestedIterator object will be instantiated and called as such:
 * NestedIterator i = new NestedIterator(nestedList);
 * while (i.hasNext()) v[f()] = i.next();
 */