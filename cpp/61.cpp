#include <iostream>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
};


class Solution {
public:
    ListNode* rotateRight(ListNode* head, int k) {
      if(head == NULL || k == 0) {
        return head;
      }
      ListNode* node = head;
      ListNode* res = NULL;
      int count = 1;
      while(node->next) {
        node = node->next;
        count++;
      }
      node->next = head;
      k = count - k % count;
      while(k > 0) {
        node = node->next;
        k--;
      }
      res = node->next;
      node->next = NULL;
      return res;
    }
};
