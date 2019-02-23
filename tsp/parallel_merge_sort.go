package tsp

import (
	"math"
	"sync"
	// "fmt"
)


func Merge_Sort(s []float64) []float64 {
	
	// nearest power of 2 to len(s)
	var k float64 = math.Ceil(math.Log2(float64(len(s))))
	
	// number of numbers to be added to the slice
	var diff int = int(math.Exp2(k) - float64(len(s)))

	// looping to add the deficit
	for i:=0;i<diff;i++ {
		s = append(s, -1)
	}

	// call normal mergesort for smaller arrays
	if(len(s) > 1024) {
		normal_mergesort(s)
	} else {
		parallel_mergesort(s)
	}

	return s[diff:]
	// fmt.Printf("%v", s)
}

func normal_mergesort(s []float64) {

	// straight-forward implementation of mergesort

	if len(s) > 1 {
		middle := len(s) / 2
		// Split in Middle and continue
		normal_mergesort(s[:middle])
		normal_mergesort(s[middle:])
		merge(s, middle)
	}
}


func parallel_mergesort(s []float64) {
	// fmt.Println("using parallel")

	if len(s) > 1 {
		var len int = len(s)
		var middle int = len / 2

		var wg sync.WaitGroup
		wg.Add(2)

		// parallely merging the smaller slices
		// defers to make-sure completion
		go func() {
			defer wg.Done()
			parallel_mergesort(s[:middle])
		}()

		go func() {
			defer wg.Done()
			parallel_mergesort(s[middle:])
		}()

		wg.Wait()
		merge(s, middle)
	}
}

func merge(s []float64, middle int) {
	helper := make([]float64, len(s))
	copy(helper, s)

	var tempLeft int = 0
	var tempRight int = middle
	var current int = 0
	high := len(s) - 1

	for tempLeft <= middle-1 && tempRight <= high {
		if helper[tempLeft] <= helper[tempRight] {
			s[current] = helper[tempLeft]
			tempLeft++
		} else {
			s[current] = helper[tempRight]
			tempRight++
		}
		current++
	}

	for tempLeft <= middle-1 {
		s[current] = helper[tempLeft]
		current++
		tempLeft++
	}
}