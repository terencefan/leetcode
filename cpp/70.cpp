#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    int climbStairs(int n) {
        int i = 0, a = 1, b = 1, t;
        while(i++<n) {
            if(i == 2) {
                b = 2;
            } else if(i > 2) {
                t = b;
                b = a + b;
                a = t;
            }

        }
        return b;
    };
};

int main() {
    cout<<Solution().climbStairs(2)<<endl;
    return 0;
}
