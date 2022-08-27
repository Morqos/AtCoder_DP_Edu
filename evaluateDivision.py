# Problem here: https://leetcode.com/problems/evaluate-division/

class Solution:
  def calcEquation(self, equations: List[List[str]], values: List[float], queries: List[List[str]]) -> List[float]:
    adj = {}

    for index in range(len(equations)):
      _from, _to = equations[index][0], equations[index][1]
      adj[_from] = adj.get(_from, []) + [[_to, values[index]]]
      adj[_to] = adj.get(_to, []) + [[_from, (1/values[index])]]

    answer = []

    def bfs(_from, _to):

      # (node, result_so_far)
      queue = [[_from, 1.0]]
      while queue:
        node_value, result_so_far = queue.pop(0)

        if node_value == _to:
          return result_so_far

        list_adj = sorted(adj[node_value])
        for adj_node in list_adj:
          queue.append([adj_node[0], adj_node[1] * result_so_far])
          adj[node_value].remove(adj_node)

      return -1


    for i in range(len(queries)):
      _from, _to = queries[i][0], queries[i][1]
      if _from not in adj or _to not in adj:
        answer.append(-1)
      elif _from == _to:
        answer.append(1)
      else:
        copy_adj = copy.deepcopy(adj)
        res = bfs(_from, _to)
        if res < 0:
          res = -1
        answer.append(res)
      adj = copy_adj

    return answer
