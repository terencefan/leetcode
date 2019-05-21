//
// Created by Terence Fan on 5/21/19.
//

#ifndef CPP_NODE_H
#define CPP_NODE_H

class Node {
public:
    int val;
    Node* next;
    Node* random;

    Node() {}

    Node(int _val, Node* _next, Node* _random) {
        val = _val;
        next = _next;
        random = _random;
    }
};

#endif //CPP_NODE_H
