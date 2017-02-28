#include <iostream>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
  public:
    ListNode* insertionSortList(ListNode* head) {
      ListNode res(0), *pre, *cur, *tmp;
      cur = head;
      while (cur) {
        pre = &res;
        while (pre->next && pre->next->val < cur->val) {
          pre = pre->next;
        }
        tmp = cur->next;
        cur->next = pre->next;
        pre->next = cur;
        cur = tmp;
      }
      return res.next;
    }
};

int main() {
  ListNode a(1), b(2), c(3), d(4);
  a.next = &c;
  c.next = &b;
  b.next = &d;
  Solution().insertionSortList(&a);
  return 0;
}
