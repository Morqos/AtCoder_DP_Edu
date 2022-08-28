# Problem here: https://leetcode.com/problems/longest-increasing-path-in-a-matrix/

class Solution:
  def longestIncreasingPath(self, matrix: List[List[int]]) -> int:
    def getAdj(cell):
      list_adj = []
      row, col = cell[0], cell[1]

      if row != 0:
        list_adj += [[row - 1, col]]

      if row != len(matrix) - 1:
        list_adj += [[row + 1, col]]

      if col != 0:
        list_adj += [[row, col - 1]]

      if col != len(matrix[0]) - 1:
        list_adj += [[row, col + 1]]

      return list_adj


    def runAlgorithm(start_cell):

      # visited is a matrix n x m initialized with False
      dp = [[1 for _ in range(len(matrix[0]))] for _ in range(len(matrix))]
      visited = [[False for _ in range(len(matrix[0]))] for _ in range(len(matrix))]


      max_value = 1
      queue = [[start_cell[0], start_cell[1]]]
      while queue:
        curr = queue.pop(0)
        visited[curr[0]][curr[1]] = True
        list_adj = getAdj(curr)

        for cell in list_adj:
          cell_row, cell_col = cell[0], cell[1]
          if not (matrix[curr[0]][curr[1]] < matrix[cell_row][cell_col]) and \
                  not visited[cell_row][cell_col]:
              queue += [[cell_row, cell_col]]

          elif matrix[curr[0]][curr[1]] < matrix[cell_row][cell_col] and \
                  dp[cell_row][cell_col] < dp[curr[0]][curr[1]] + 1:
            dp[cell_row][cell_col] = dp[curr[0]][curr[1]] + 1
            max_value = max(max_value, dp[cell_row][cell_col])
            queue += [[cell_row, cell_col]]

      return max_value


    answer = runAlgorithm([0, 0])
    
    return answer