#include <bits/stdc++.h>

typedef long long ll; 
using namespace std;

int INF = 1e9 + 7;

int main()
{
  int n, w;
  scanf("%d %d", &n, &w);

  vector<int> weights(n);
  vector<int> values(n);

  ll maxValue = 0;
  
  for(int i = 0; i < n; ++i)
  {
    scanf("%d %d", &weights[i], &values[i]);
    maxValue += values[i];
  }

  vector<ll> dp(maxValue + 1, INF);
  dp[0] = 0;

  for(int j = 0; j < n; ++j)
  {
    for(int i = maxValue - 1; i >= 0; --i)
    {
      if(
        i + values[j] <= maxValue &&
        (dp[i] != 0 || i == 0) &&
        dp[i] + weights[j] <= w
      )
        dp[i + values[j]] = min(dp[i + values[j]], dp[i] + weights[j]);
    }
  }

  int ans = -1;
  for(int i = maxValue; i >= 0; --i)
  {
    if(dp[i] <= w && dp[i] != 0)
    {
      ans = i;
      break;
    }
  }

  cout << ans << endl;


  return 0;
}