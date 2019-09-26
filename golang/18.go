package main

import (
	"fmt"
	"io/ioutil"
	"log"
	//	"regexp"
	//	"strings"
)

var m [100][100]int

// https://blog.csdn.net/qq_38362049/article/details/80960374

func dfs1(x int, y int, n int) int {
	if x == n {
		return 0
	}
	ans1 := dfs1(x+1, y, n)
	ans2 := dfs1(x+1, y+1, n)
	if ans1 > ans2 {
		return ans1 + m[x][y]
	} else {
		return ans2 + m[x][y]
	}
}

func dfs2(a [][]int) int {
	fmt.Println("hello world")
	fmt.Println("hello world", a)
	return 1
}

func dfs3(n int) int {

	f := [100][100]int{{0}}

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

	//	fmt.Println(f)
	return f[0][0]
}

func dfs4(n int) int {

	f := [100][100]int{{0}}
	f[0][0] = m[0][0]        //初始化f[0][0]为第0层的值；
	for i := 1; i < n; i++ { //更新之后14层的权值；
		for j := 0; j <= i; j++ { //每层数字为i+1个；
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
	//for i := 0; i <= n; i++ {
	for i := 0; i < n; i++ {
		if f[n-1][i] > ans {
			ans = f[n-1][i]
		}
	} //循环找到最后一层最大的权值
	//	fmt.Printf("%d\n", ans)
	return ans
}

func main() {
	//fmt.Println("Problem 18")

	fileName := "p08_triangle.txt"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("File contents: \n%s", content)

	//	s := regexp.MustCompile("\n")
	//	lines := s.Split(strings.Trim(content, " |\t|\n|\r\n"), -1)
	//	fmt.Println(lines)

	//m := [100][100]int{{0}}
	m = [100][100]int{{0}}
	i, j := 0, 0
	for _, v := range content {
		//fmt.Println(v, v-'0')
		//		fmt.Println(v)

		if v == '\n' {
			i++
			j = 0
			continue
		}

		if v == ' ' {
			j++
			continue
		}
		//m[i][j] = m[i][j]*10 + int(v-'0')
		if m[i][j] != 0 {
			m[i][j] = m[i][j]*10 + int(v-'0')
		} else {
			m[i][j] = int(v - '0')
		}

	}

	//	fmt.Println(m)

	for i := 0; i < 40; i++ {
		for j := 0; j < 40; j++ {
			if m[i][j] != 0 {
				fmt.Print(m[i][j], " ")
			}
		}
		fmt.Println("")
	}

	fmt.Println("dfs1: ", dfs1(0, 0, 15))

	fmt.Println("dfs3: ", dfs3(15))
	fmt.Println("dfs4: ", dfs4(15))
}
