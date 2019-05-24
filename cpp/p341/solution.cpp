#include "nested_integer.h"

#include <stack>

typedef vector<NestedInteger> NestedList;

class NestedIterator {
private:
    stack<NestedList> stackList;
    stack<NestedList::iterator*> stackIter;

public:
    NestedIterator(vector<NestedInteger> &nestedList) {
        push(nestedList, nestedList.begin());
    }

    void push(NestedList list, NestedList::iterator *it) {
        while (*it != list.end()) {
            stackList.push(list);
            stackIter.push(it);
            if ((*it)->isInteger()) {
                break;
            } else {
                list = it->getList();
                it = *it->getList().begin();
            }
        }
    }

    int next() {
        auto it = stackIter.top();
        auto list = stackList.top();

        stackList.pop();
        stackIter.pop();

        if (it->isInteger()) {
            int r = it->getInteger();
            it++;
            push(list, it);
            return r;
        } else {
        }
        return r;

    }

    bool hasNext() {
        return stackList.size() > 0;
    }
};

/**
 * Your NestedIterator object will be instantiated and called as such:
 * NestedIterator i(nestedList);
 * while (i.hasNext()) cout << i.next();
 */
