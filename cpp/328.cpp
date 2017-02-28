#include <iostream>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
  public:
    ListNode* oddEvenList(ListNode* head) {
      if (!head || !head->next) {
        return head;
      }
      ListNode *l1, *l2, *h1, *h2, *p1, *p2;
      l1 = h1 = head;
      l2 = h2 = head->next;
      while (l2->next && l2->next->next) {
        p1 = l1;
        p2 = l2;
        l1 = l1->next->next;
        l2 = l2->next->next;
        p1->next = l1;
        p2->next = l2;
      }
      if (l2->next) {
        l1->next = l2->next;
        l2->next = NULL;
        l1->next->next = h2;
      } else {
        l1->next = h2;
      }
      return head;
    }
};

int main() {

  int i = 1;
  ListNode *h, *a;
  h = a = new ListNode(i);
  while (i < 5) {
    i++;
    a->next = new ListNode(i);
    a = a->next;
  }
  h = Solution().oddEvenList(h);
  while (h) {
    cout<<h->val<<endl;
    h = h->next;
  }
  return 0;
}
