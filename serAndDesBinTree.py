# Problem here: https://leetcode.com/problems/serialize-and-deserialize-binary-tree/submissions/

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Codec:

  def serialize(self, root):
    """Encodes a tree to a single string.

    :type root: TreeNode
    :rtype: str
    """
    if not root:
      return ''
    
    serialized_tree = [root.val]
    queue = [root]
    while queue:
      node = queue.pop(0)

      serialized_tree.append(node.left.val if node.left else None)
      serialized_tree.append(node.right.val if node.right else None)

      if node.left:
        queue.append(node.left)
      if node.right:
        queue.append(node.right)
    
    string_tree = ''
    for index in range(len(serialized_tree)):
      string_tree += str(serialized_tree[index])
      if index != len(serialized_tree) - 1:
        string_tree += ','
    
    return string_tree


  def deserialize(self, data):
    """Decodes your encoded data to tree.

    :type data: str
    :rtype: TreeNode
    """
    if not data:
      return None
    
    array_tree = data.split(',')
    queue = [TreeNode(int(array_tree.pop(0)))]

    first = True
    root = None

    while queue and array_tree:
      node = queue.pop(0)

      if not node:
        continue
      
      left_value_str = array_tree.pop(0)
      right_value_str = array_tree.pop(0)

      left = TreeNode(int(left_value_str)) if left_value_str != 'None' else None
      right = TreeNode(int(right_value_str)) if right_value_str != 'None' else None

      node.left = left
      node.right = right

      queue.append(left)
      queue.append(right)

      if first:
        first = False
        root = node
    
    return root
        

# Your Codec object will be instantiated and called as such:
# ser = Codec()
# deser = Codec()
# ans = deser.deserialize(ser.serialize(root))