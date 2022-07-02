#include <bits/stdc++.h>
 
using namespace std;

int INF = 1e9 + 7;

int main()
{
  int n, k;
  scanf("%d %d", &n, &k);

  vector<int> massi(n);
  for(int i = 0; i < n; ++i)
  {
    scanf("%d", &massi[i]);
  }

  vector<int> dp(n, INF);
  dp[0] = 0;

  for(int i = 0; i < n; ++i)
  {
    for(int j = 1; j <= k; ++j)
    {
      if(i + j < n)
      {
        dp[i + j] = min(dp[i + j], abs(massi[i] - massi[i + j]) + dp[i]);
      }
    }
  }

  cout << dp[n - 1] << endl;

  return 0;
}