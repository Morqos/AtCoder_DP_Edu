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
  vector<ll> dp(w + 1);
  
  for(int i = 0; i < n; ++i)
  {
    scanf("%d %d", &weights[i], &values[i]);
  }

  for(int j = 0; j < n; ++j)
  {
    for(int i = w - 1; i >= 0; --i)
    {
      if(i + weights[j] <= w && (dp[i] != 0 || i == 0))
        dp[i + weights[j]] = max(dp[i + weights[j]], dp[i] + values[j]);
    }
  }

  ll ans = -1;
  for(int i = 0; i <= w; ++i)
  {
    ans = max(ans, dp[i]);
  }

  cout << ans << endl;


  return 0;
}