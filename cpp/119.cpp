#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    vector<int> getRow(int k) {
        vector<int> res;
        for(int i=0; i<k; i++) {
            res.push_back(1);
            for(int j=i-1;j>0;j--){
                res[j] = res[j] + res[j - 1];
            }
        }
        return res;
    };
};

int main() {
    Solution().getRow(5);
    return 0;
}
