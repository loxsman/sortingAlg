package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func fillArr(array []int) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range array {
		array[i] = rng.Intn(10000000)
	}
}

func bubleSort(array []int) {
	size := len(array)
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func shakerSort(array []int) {
	size := len(array)
	left := 0
	right := size - 1
	for left <= right {
		for i := left; i < right; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
			}
		}
		right--
		for j := right; j > left; j-- {
			if array[j-1] > array[j] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
		left++
	}
}

func selectionSort(array []int) {
	size := len(array)
	for i := 0; i < size-1; i++ {
		minInd := i
		for j := i + 1; j < size; j++ {
			if array[j] < array[minInd] {
				minInd = j
			}
		}
		array[i], array[minInd] = array[minInd], array[i]
	}
}

func insertionSort(array []int) {
	size := len(array)
	for i := 1; i < size; i++ {
		key := array[i]
		j := i - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j = j - 1
		}
		array[j+1] = key
	}
}

func quickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	size := len(array)
	pivotIndex := size / 2
	pivot := array[pivotIndex]

	var lessPivot []int
	var highPivot []int

	for i, k := range array {
		if i == pivotIndex {
			continue
		}
		if k <= pivot {
			lessPivot = append(lessPivot, k)
		} else {
			highPivot = append(highPivot, k)
		}
	}

	return append(append(quickSort(lessPivot), pivot), quickSort(highPivot)...)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func mergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	mid := len(array) / 2
	left := mergeSort(array[:mid])
	right := mergeSort(array[mid:])

	return merge(left, right)
}

func main() {
	size := 100000
	unSortArray := make([]int, size)
	fillArr(unSortArray)

	var wg sync.WaitGroup
	wg.Add(6)

	go func() {
		defer wg.Done()
		array := make([]int, size)
		copy(array, unSortArray)
		startBuble := time.Now()
		bubleSort(array)
		durationBuble := time.Since(startBuble)
		fmt.Printf("Bubble sort: %d ms\n", durationBuble.Milliseconds())
	}()

	go func() {
		defer wg.Done()
		array := make([]int, size)
		copy(array, unSortArray)
		startShaker := time.Now()
		shakerSort(array)
		durationShaker := time.Since(startShaker)
		fmt.Printf("Shaker sort: %d ms\n", durationShaker.Milliseconds())
	}()

	go func() {
		defer wg.Done()
		array := make([]int, size)
		copy(array, unSortArray)
		startSelection := time.Now()
		selectionSort(array)
		durationSelection := time.Since(startSelection)
		fmt.Printf("Selection sort: %d ms\n", durationSelection.Milliseconds())
	}()

	go func() {
		defer wg.Done()
		array := make([]int, size)
		copy(array, unSortArray)
		startInsertion := time.Now()
		insertionSort(array)
		durationInsertion := time.Since(startInsertion)
		fmt.Printf("Insertion sort: %d ms\n", durationInsertion.Milliseconds())
	}()

	go func() {
		defer wg.Done()
		array := make([]int, size)
		copy(array, unSortArray)
		startQuick := time.Now()
		array = quickSort(array)
		durationQuick := time.Since(startQuick)
		fmt.Printf("Quick sort: %d ms\n", durationQuick.Milliseconds())
	}()

	go func() {
		defer wg.Done()
		array := make([]int, size)
		copy(array, unSortArray)
		startMerge := time.Now()
		array = mergeSort(array)
		durationMerge := time.Since(startMerge)
		fmt.Printf("Merge sort: %d ms\n", durationMerge.Milliseconds())
	}()

	wg.Wait()
}
