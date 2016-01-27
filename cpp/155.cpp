#include <iostream>
#include <stack>
using namespace std;

class MinStack {

    protected:
        stack<int> s1;
        stack<int> s2;

    public:

        void push(int x) {
            s1.push(x);
            if(s2.empty() || x <= getMin()) {
                s2.push(x);
            }
        }

        void pop() {
            if(top() == getMin()) {
                s2.pop();
            }
            s1.pop();
        }

        int top() {
            return s1.top();
        }

        int getMin() {
            return s2.top();
        }

};

int main() {
    MinStack s;
    s.push(0);
    s.push(1);
    s.push(0);
    s.getMin();
    s.pop();
    s.getMin();
    cout<<"finish"<<endl;
}
