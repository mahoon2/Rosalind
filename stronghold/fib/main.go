package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("fib.in")

	if err != nil {
		panic(err)
	}

	var n, k int
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ = strconv.Atoi(scanner.Text())

	dp := [2]int{1, 1}

	if n <= 2 {
		fmt.Printf("1\n")
		return
	}

	for i := range n - 2 {
		dp[i%2] = k*dp[i%2] + dp[(i+1)%2]
	}
	fmt.Printf("%d\n", dp[(n-1)%2])
}
