#include <iostream>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
  public:
    ListNode* sortList(ListNode* head) {
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
  ListNode a(1), b(4), c(7), d(9), e(5);
  a.next = &b;
  b.next = &c;
  c.next = &d;
  d.next = &e;
  ListNode *head = Solution().sortList(&a);
  while(head) {
    cout<<head->val<<endl;
    head = head->next;
  }
  return 0;
}
