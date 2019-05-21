#include "node.h"

#include <map>
#include <iostream>

using namespace std;

class Solution {
public:
    Node* copyRandomList(Node* head) {
        map<Node*, Node*> map;

        Node *copiedHead = new Node(-1, nullptr, nullptr);

        Node* node = head;
        Node* copied = copiedHead;

        while (node != nullptr) {
            Node* next = new Node(node->val, nullptr, nullptr);
            map.insert(make_pair(node, next));
            node = node->next;
            copied->next = next;
            copied = copied->next;
        }

        node = head;
        copied = copiedHead->next;
        while (copied != nullptr) {
            if (node->random != nullptr) {
                auto it = map.find(node->random);
                copied->random = it->second;
            }
            node = node->next;
            copied = copied->next;
        }

        return copiedHead->next;
    }
};

int main() {
    Node* head = new Node(-1, nullptr, nullptr);
    Node* copied = (new Solution())->copyRandomList(head);
}
