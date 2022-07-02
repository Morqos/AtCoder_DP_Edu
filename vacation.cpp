#include <bits/stdc++.h>
 
using namespace std;

int INF = 1e9 + 7;

int main()
{
  int n;
  scanf("%d", &n);

  vector<vector<int>> points(n, vector<int>(3));
  vector<vector<int>> dp(n, vector<int>(3));
  for(int i = 0; i < n; ++i)
  {
    scanf("%d %d %d", &points[i][0], &points[i][1], &points[i][2]);
  }

  dp[0][0] = points[0][0];
  dp[0][1] = points[0][1];
  dp[0][2] = points[0][2];

  for(int i = 0; i < n - 1; ++i)
  {
    for(int j = 0; j < 3; ++j)
    {
      for(int k = 1; k < 3; ++k)
      {
        dp[i + 1][(j + k) % 3] = max(dp[i + 1][(j + k) % 3], dp[i][j] + points[i + 1][(j + k) % 3]); 
      }
    }
  }

  cout << max(dp[n - 1][0], max(dp[n - 1][1], dp[n - 1][2])) << endl;

  return 0;
}