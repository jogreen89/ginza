// binary_search.go
//
// A Go implementation of a Binary Search algorithm
// and the findmin() iterative search procedure.
// (c) 2016 zubernetes

package main

import "fmt"

func main() {
	a := []int{2, 3, 5, 6, 7, 11, 13}
	n := findmin(a)
	fmt.Printf("The findmin function returns %d.\n",n)

	key := 6
	n = BinarySearch(a, key, 0, 7) 
	fmt.Printf("The BinarySearch function"+
			   " returns the position of %d at position %d.\n", key, n-1)
}

func BinarySearch(a []int, key int, low int, high int) int {
	// return the position of the key in the array
	if (low < high) {
		mid := (low + high) / 2
		if (a[mid] == key) {
			fmt.Printf("a[mid] is %d\n", a[mid])
			return mid
		} else if (a[mid] > key) {
			return BinarySearch(a, key, low, mid-1)
		} else if (a[mid] < key) {
			return BinarySearch(a, key, mid+1, high)
		}
	}
	return 0;
}

func findmin(a []int) int {
	// return the minimum element in a []int
	// worst case time complexity of findmin is theta(n)
	min := 0
	n := len(a)
	for i := 0; i < n - 1; i++ {
		if (i == 0) {
			min = a[i] 
		}
		if (a[i] > a[i+1]) {
			min = a[i+1]
		}
	}
	return min
}
