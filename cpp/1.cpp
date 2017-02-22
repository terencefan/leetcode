#include <iostream>
#include <vector>
#include <unordered_map>
using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {

        unordered_map<int, int> hash;
        vector<int> result;

        for(int i=0;i<nums.size();i++) {
            int temp  = target - nums[i];
            if(hash.find(temp) != hash.end()) {
                result.push_back(hash[temp] + 1);
                result.push_back(i + 1);
                break;
            }
            hash[nums[i]] = i;
        }

        return result;
    }
};

void d(int* a) {
  int c = 3;
  a = &c;
}

int main() {
  int *a;
  int b = 2;
  a = &b;
  d(a);
}
