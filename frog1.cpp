#include <bits/stdc++.h>
 
using namespace std;

int main()
{
  int n;
  scanf("%d", &n);

  vector<int> massi(n);
  for(int i = 0; i < n; ++i)
  {
    scanf("%d", &massi[i]);
  }

  vector<int> dp(n);
  dp[1] = abs(massi[0] - massi[1]);

  for(int i = 0; i < n - 1; ++i)
  {
    dp[i + 1] = min(dp[i + 1], dp[i] + abs(massi[i] - massi[i + 1]));

    if(i < n - 2)
    {
      dp[i + 2] = dp[i] + abs(massi[i] - massi[i + 2]);
    }
  }

  cout << dp[n - 1] << endl;


  return 0;
}