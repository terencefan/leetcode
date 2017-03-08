// Forward declaration of isBadVersion API.
bool isBadVersion(int version);

class Solution {
public:
    int firstBadVersion(int n) {
        long i = 1;
        long j = n;
        long k;
        while (i != j) {
            k = (i + j) / 2;
            if (isBadVersion(k)) {
                j = k;
            } else {
                i = k + 1;
            }
        }
        return i;
    }
};
