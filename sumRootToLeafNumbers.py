# Problem here: https://leetcode.com/problems/sum-root-to-leaf-numbers/

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
  def sumNumbers(self, root: Optional[TreeNode]) -> int:
    
    paths = []
    def dfs(node, path):
      if not node:
        return
      if not node.left and not node.right:
        paths.append(path + str(node.val))
      dfs(node.left, path  + str(node.val))
      dfs(node.right, path  + str(node.val))
    
    dfs(root, '')
    
    print(paths)
    sum_paths = 0
    for str_path in paths:
      sum_paths += int(str_path)
    
    return sum_paths