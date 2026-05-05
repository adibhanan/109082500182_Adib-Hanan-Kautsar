package main

import "fmt"

type set [2022]int

func exist(t set, n int, val int) bool {
	for i := 0; i < n; i++ {
		if t[i] == val {
			return true
		}
	}
	return false
}

func inputSet(t *set, n *int) {
	var x int
	*n = 0

	for {
		fmt.Scan(&x)

		if exist(*t, *n, x) {
			break
		}

		(*t)[*n] = x
		*n++
	}
}

func findIntersection(T1 set, T2 set, n1 int, n2 int, T3 *set, h *int) {
	*h = 0

	for i := 0; i < n1; i++ {
		if exist(T2, n2, T1[i]) {
			(*T3)[*h] = T1[i]
			*h++
		}
	}
}

func printSet(t set, n int) {
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(t[i])
	}
	fmt.Println()
}

func main() {
	var s1, s2, s3 set
	var n1, n2, n3 int

	inputSet(&s1, &n1)
	inputSet(&s2, &n2)

	findIntersection(s1, s2, n1, n2, &s3, &n3)

	printSet(s3, n3)
}
