---
title: 代码片段
top: 4
---
[[toc]]

# Introduction to Data Structure
## Binary Tree
### Traverse A Tree
#### [Preorder Traversal](https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/928/)
> Chez Scheme

```scheme
; Definition for a binary tree node.
; (define TreeNode
;     (case-lambda 
;         [(x) (TreeNode x `() `())]
;         [(x left) (TreeNode x left `())]
;         [(x left right) (cons x (cons left right))]))
(define (preorderTraversal root)
    (if (null? root)
        (list)
        (let ((val (car root))
            (left (if (null? (cdr root)) `() (car (cdr root))))
            (right (if (null? (cdr root)) `() (cdr (cdr root)))))
            (append (list val) (preorderTraversal left) (preorderTraversal right)))))
```
> c++


```cpp
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> opt;
        if (!root)
            return opt;
        if (root->val)
            opt.push_back(root->val);
        if (root->left) {
            vector<int> opt_left = this->preorderTraversal(root->left);
            opt.insert(opt.end(), opt_left.begin(), opt_left.end());
        }
        if (root->right) {
            vector<int> opt_right = this->preorderTraversal(root->right);
            opt.insert(opt.end(), opt_right.begin(), opt_right.end());
        }
        return opt;
    }
};
```
> python

```python
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def preorderTraversal(self, root: 'TreeNode') -> 'List[int]':
        opt = []
        if root is None:
            return opt
        if root.val is not None:
            opt.append(root.val)
        if root.left is not None:
            opt += self.preorderTraversal(root.left)
        if root.right is not None:
            opt += self.preorderTraversal(root.right)
        return opt
```
> php

```php
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */
class Solution {

    /**
     * @param TreeNode $root
     * @return Integer[]
     */
    function preorderTraversal($root) {
        $opt = [];
        if (!$root)
            return $opt;
        if ($root->val)
            array_push($opt, $root->val);
        if ($root->left)
            $opt = array_merge($opt, $this->preorderTraversal($root->left));
        if ($root->right)
            $opt = array_merge($opt, $this->preorderTraversal($root->right));
        return $opt;
    }
}
```
#### [Inorder Traversal](https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/929/)
> c++

```cpp
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> opt;
        if (!root)
            return opt;
        if (root->left) {
            vector<int> opt_left = this->inorderTraversal(root->left);
            opt.insert(opt.end(), opt_left.begin(), opt_left.end());
        }
        if (root->val)
            opt.push_back(root->val);
        if (root->right) {
            vector<int> opt_right = this->inorderTraversal(root->right);
            opt.insert(opt.end(), opt_right.begin(), opt_right.end());
        }
        return opt;
    }
};
```
> python

```python
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def inorderTraversal(self, root: 'TreeNode') -> 'List[int]':
        opt = []
        if root is None:
            return opt
        if root.left is not None:
            opt += self.inorderTraversal(root.left)
        if root.val is not None:
            opt.append(root.val)
        if root.right is not None:
            opt += self.inorderTraversal(root.right)
        return opt
```
> php

```php
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */
class Solution {

    /**
     * @param TreeNode $root
     * @return Integer[]
     */
    function inorderTraversal($root) {
        $opt = [];
        if (!$root)
            return $opt;
        if ($root->left)
            $opt = array_merge($opt, $this->inorderTraversal($root->left));
        if ($root->val)
            array_push($opt, $root->val);
        if ($root->right)
            $opt = array_merge($opt, $this->inorderTraversal($root->right));
        return $opt;
    }
}
```
#### [Postorder Traversal](https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/930/)
> c++

```cpp
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> opt;
        if (!root)
            return opt;
        if (root->left) {
            vector<int> opt_left = this->postorderTraversal(root->left);
            opt.insert(opt.end(), opt_left.begin(), opt_left.end());
        }
        if (root->right) {
            vector<int> opt_right = this->postorderTraversal(root->right);
            opt.insert(opt.end(), opt_right.begin(), opt_right.end());
        }
        if (root->val) {
            opt.push_back(root->val);
        }
        return opt;
    }
};
```
> python

```python
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def postorderTraversal(self, root: 'TreeNode') -> 'List[int]':
        opt = []
        if root is None:
            return opt
        if root.left is not None:
            opt += self.postorderTraversal(root.left)
        if root.right is not None:
            opt += self.postorderTraversal(root.right)
        if root.val is not None:
            opt.append(root.val)
        return opt
```
> php

```php
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */
class Solution {

    /**
     * @param TreeNode $root
     * @return Integer[]
     */
    function postorderTraversal($root) {
        $opt = [];
        if (!$root)
            return $opt;
        if ($root->left)
            $opt = array_merge($opt, $this->postorderTraversal($root->left));
        if ($root->right)
            $opt = array_merge($opt, $this->postorderTraversal($root->right));
        if ($root->val)
            array_push($opt, $root->val);
        return $opt;
    }
}
```
#### [Level Order Traversal](https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/931/)
> c++

```cpp
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector<TreeNode*> q;
        vector<TreeNode*>::iterator iter;
        vector<vector<int>> opt;
        if (NULL == root)
            return opt;
        q.push_back(root);
        
        while(q.size() > 0) {
            vector<TreeNode*> q2;
            vector<int> opt2;
            for(iter=q.begin(); iter!=q.end(); iter++) {
                opt2.push_back((*iter)->val);
                if (NULL != (*iter)->left)
                    q2.push_back((*iter)->left);
                if (NULL != (*iter)->right)
                    q2.push_back((*iter)->right);
            }
            q = q2;
            if (opt2.size() > 0)
                opt.push_back(opt2);
        }
        return opt;
    }
};
```
> python

```python
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def levelOrder(self, root: 'TreeNode') -> 'List[List[int]]':
        opt = []
        if root is None:
            return opt
        queue = [root]
        while queue:
            opt_2 = []
            queue_2 = []
            for node in queue:
                opt_2.append(node.val)
                if node.left:
                    queue_2.append(node.left)
                if node.right:
                    queue_2.append(node.right)
            queue = queue_2
            if opt_2:
                opt.append(opt_2)
        return opt
```
> php

```php
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */
class Solution {

    /**
     * @param TreeNode $root
     * @return Integer[][]
     */
    function levelOrder($root) {
        $opt = [];
        if (!$root)
            return $opt;
        $queue = [$root];
        while($queue) {
            $queue2 = [];
            $opt2 = [];
            foreach ($queue as $node) {
                array_push($opt2, $node->val);
                if ($node->left)
                    array_push($queue2, $node->left);
                if ($node->right)
                    array_push($queue2, $node->right);
            }
            if ($opt2)
                array_push($opt, $opt2);
            $queue = $queue2;
        }
        return $opt;
    }
}
```