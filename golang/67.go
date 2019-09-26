package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

var m [150][150]int

func dfs3(n int) int {

	f := [150][150]int{{0}}

	for i := 0; i < n; i++ {
		f[n-1][i] = m[n-1][i]
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			f[i][j] = m[i][j]
			if f[i+1][j] > f[i+1][j+1] {
				f[i][j] += f[i+1][j]
			} else {
				f[i][j] += f[i+1][j+1]
			}
		}
	}

	return f[0][0]
}

func dfs4(n int) int {

	f := [150][150]int{{0}}
	f[0][0] = m[0][0]
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				// the left column can only access its accesstor in the same position
				f[i][j] = f[i-1][j]
			} else {
				if f[i-1][j-1] > f[i-1][j] {
					f[i][j] = f[i-1][j-1]
				} else {
					//if f[i-1][j] > f[i-1][j-1] {
					f[i][j] = f[i-1][j]
				}
			}
			f[i][j] += m[i][j]
		}
	}

	ans := 0
	for i := 0; i <= n; i++ {
		if f[n-1][i] > ans {
			ans = f[n-1][i]
		}
	}
	return ans
}

func main() {

	fileName := "p067_triangle.txt"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	m = [150][150]int{{0}}
	i, j := 0, 0
	for _, v := range content {

		if v == '\n' {
			i++
			j = 0
			continue
		}

		if v == ' ' {
			j++
			continue
		}

		if m[i][j] != 0 {
			m[i][j] = m[i][j]*10 + int(v-'0')
		} else {
			m[i][j] = int(v - '0')
		}

	}

	for i := 0; i < 100; i++ {
		for j := 0; j <= i; j++ {
			if m[i][j] != 0 {
				fmt.Print(m[i][j], " ")
			}
		}
		fmt.Println("")
	}

	fmt.Println("dfs3: ", dfs3(100))
	fmt.Println("dfs4: ", dfs4(100))
}
