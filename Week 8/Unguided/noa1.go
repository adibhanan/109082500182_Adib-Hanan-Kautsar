package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x int
	y int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func dalamLingkaran(t Titik, l Lingkaran) bool {
	dx := float64(t.x - l.pusat.x)
	dy := float64(t.y - l.pusat.y)
	jarak := math.Sqrt(dx*dx + dy*dy)

	return jarak <= float64(l.r)
}

func main() {
	var l1, l2 Lingkaran
	var t Titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)

	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)

	fmt.Scan(&t.x, &t.y)

	dalam1 := dalamLingkaran(t, l1)
	dalam2 := dalamLingkaran(t, l2)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
