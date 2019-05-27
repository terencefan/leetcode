package p341;

import java.util.List;

public class NestedIntegerItem implements NestedInteger {

    private List<NestedInteger> l;
    private int val;

    NestedIntegerItem(int val) {
        this.val = val;
    }

    NestedIntegerItem(List<NestedInteger> l) {
        this.l = l;
    }

    @Override
    public boolean isInteger() {
        return l == null;
    }

    @Override
    public Integer getInteger() {
        return val;
    }

    @Override
    public List<NestedInteger> getList() {
        return l;
    }
}
