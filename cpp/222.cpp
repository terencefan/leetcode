#include<iostream>

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;

    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    int countNodes(TreeNode* root) {
        if (!root) return 0;
        TreeNode* p = root;
        int height = 0;
        while (p) {
            p = p->left;
            height++;
        }
        int res = (1 << (height - 1)) - 1;
        TreeNode* t = root;
        for (height = height - 2; height >= 0; height--) {
            p = t->left;
            for (int j = height; j > 0; j--) p = p->right;
            if (p != NULL) {
                res += 1 << height;
                t = t->right;
            } else {
                t = t->left;
            }
        }
        if (t != NULL) res += 1;
        return res;
    }
};
