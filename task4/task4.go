package main

import (
	"fmt"
	"math/rand"
)

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

func main() {
	slice := []int{3, 44, 11, 23, 51, 94, 5526, 87, 29, 4972, 6777, 42124, 48, 62, 30, 22, 1, 98, 1111}
	fmt.Println(quicksort(slice))
}