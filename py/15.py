import numpy as np
dp=np.zeros((50,50))
dp[0][0]=1
for i in range(1,41):
	dp[i][0]=1
	for j in range(1,41):
		dp[i][j]=dp[i-1][j-1]+dp[i-1][j]
print(dp[40][20])
