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

int main() {
    int a[4] = {2, 7, 11, 15};
    vector<int> nums(a, a+3);
    vector<int> result = Solution().twoSum(nums, 9);
    for(int i=0;i<result.size();i++) {
        printf("%d ", result[i]);
    }
    return 0;
}
