#include <iostream>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
};


class Solution {
public:
    void reorderList(ListNode* head) {
      // Boundary conditions.
      if (head == NULL || head->next == NULL) {
        return;
      }

      ListNode *l1, *l2, *tmp;
      l1 = l2 = head;
      // Get the mid Node of the list
      while (l2->next && l2->next->next) {
        l1 = l1->next;
        l2 = l2->next->next;
      }
      // Save the first Node of the second half list.
      l2 = l1->next;
      // Cut the list.
      l1->next = NULL;

      // Reverse the second half list, *l1 is used for recording previous Node.
      l1 = NULL;
      while (l2) {
        tmp = l2->next;
        l2->next = l1;
        l1 = l2;
        l2 = tmp;
      }

      // Merge the two half lists
      l2 = l1;
      l1 = head;
      while (l2) {
        tmp = l2->next;
        l2->next = l1->next;
        l1->next = l2;
        l1 = l1->next->next;
        l2 = tmp;
      }

    }
};
