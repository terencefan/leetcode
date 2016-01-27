#include <iostream>
#include <vector>
using namespace std;

class Solution {
    public:
        void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
            int i = m - 1;
            int j = n - 1;
            int k = m + n - 1;

            while(i >= 0 && j >= 0) {
                printf("%d %d %d\n", i, j, k);
                if(nums1[i] > nums2[j]) {
                    nums1[k--] = nums1[i--];
                } else {
                    nums1[k--] = nums2[j--];
                }
            }

            while(j >= 0) {
                nums1[k--] = nums2[j--];
            }

            // for(int i=0;i<nums1.size();i++) {
            //     printf("%d ", nums1[i]);
            // }

        }

};

int main() {
    vector<int> nums1(8), nums2;
    nums1[0] = 1;nums1[1] = 2;nums1[2] = 7;nums1[3] = 9;
    nums2.push_back(2);
    nums2.push_back(8);
    nums2.push_back(9);
    nums2.push_back(10);
    Solution().merge(nums1, 4, nums2, 4);
    return 0;
}
