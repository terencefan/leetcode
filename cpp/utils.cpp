#include<iostream>

using namespace std;

struct ListNode {
  int val;
  ListNode *next;

  ListNode(int x) {
    val = x;
    next = NULL;
  }

};

ListNode* generateList(int count) {
  if (count < 1) {
    return NULL;
  }
  int i = 1;
  ListNode *h, *p;
  h = p = new ListNode(1);
  while (i < count) {
    i += 1;
    p->next = new ListNode(i);
    p = p->next;
  }
  return h;
}

int main() {
  ListNode* h = generateList(12);
  while(h) {
    cout<<h->val<<endl;
    h = h->next;
  }
  return 0;
}
