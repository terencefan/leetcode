#include <iostream>
using namespace std;

class Solution {
public:
    int lengthOfLastWord(string s) {
        int result = 0;
        for(int i=s.length() - 1;i>=0;i--) {
            if(s[i] == ' ') {
                if(result) {
                    break;
                }
            } else {
                result += 1;
            }
        }
        return result;
    }
};

int main() {
    int result = Solution().lengthOfLastWord("aa ");
    cout<<result<<endl;
    return 0;
}
