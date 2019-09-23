package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// https://blog.csdn.net/qq_38362049/article/details/80960374

func dfs1(x int, y int, n int) int {
	if x == n {
		return 0
	}
	ans1 := dfs(x+1, y, n)
	ans2 := dfs(x+1, y+1, n)
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

func dfs3(a [][]int) int {
	fmt.Println("hello world")
	fmt.Println("hello world", a)
	return 1
}

func dfs4(a [][]int) int {
	fmt.Println("hello world")
	fmt.Println("hello world", a)
	return 1

}

func main() {
	fmt.Println("Problem 18")

	fileName := "p08_triangle.txt"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)

	fmt.Println("test")
}
