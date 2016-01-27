#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    vector<int> plusOne(vector<int>& digits) {

        int pos = digits.size(), carry = 1;
        while(--pos>=0 && carry) {
            carry = digits[pos] == 9?1:0;
            digits[pos] = (digits[pos] + 1) % 10;
        }

        if(carry) {
            digits[0] = 1;
            digits.push_back(0);
        }

        return digits;

    }
};
