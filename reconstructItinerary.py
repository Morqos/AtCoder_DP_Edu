# Problem here: https://leetcode.com/problems/reconstruct-itinerary/

class Solution:
  def findItinerary(self, tickets: List[List[str]]) -> List[str]:
    adj = {}
    itinerary = []
    n = len(tickets)
    
    for _from, _to in tickets:
      adj[_from] = adj.get(_from, []) + [_to]
      
    def dfs(_from, adj, path) -> bool:
      nonlocal n
      
      if _from not in adj or not adj[_from]:
        if len(path) == n + 1:
          itinerary.append(path)
          return True
        else:
          return False
      
      adj_sorted_list = sorted(adj[_from])
      for _to in adj_sorted_list:
        adj[_from].remove(_to)
        is_okay = dfs(_to, adj, path + [_to])
        if is_okay:
          return True
        adj[_from] += [_to]
      return False
    
    _ = dfs("JFK", adj, ["JFK"])
    return itinerary[0]