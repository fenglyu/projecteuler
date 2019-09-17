package main

import (
	"fmt"
)

func main() {
	dp := [21][21]int64{}

	//	fmt.Println(dp)
	for i := 0; i <= 20; i++ {
		for j := 0; j <= 20; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] += dp[i][j-1] + dp[i-1][j]
			}
		}
	}

	fmt.Println(dp[20][20])
}
