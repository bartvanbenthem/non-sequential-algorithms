package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 50_000_000

type DataType interface {
	~float64 | ~int | ~string
}

func IsSorted[T DataType](data []T) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

func InsertSort[T DataType](data []T) {
	i := 1
	for i < len(data) {
		h := data[i]
		j := i - 1
		for j >= 0 && h < data[j] {
			data[j+1] = data[j]
			j -= 1
		}
		data[j+1] = h
		i += 1
	}
}

func Merge[T DataType](left, right []T) []T {
	result := make([]T, len(left)+len(right))
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {

			result[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}
	return result
}

func MergeSort[T DataType](data []T) []T {
	if len(data) > 100 {
		middle := len(data) / 2
		left := data[:middle]
		right := data[middle:]
		data = Merge(MergeSort(right), MergeSort(left))
	} else {
		InsertSort(data)
	}
	return data
}

func main() {
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}

	start := time.Now()
	result := MergeSort[float64](data)
	elapsed := time.Since(start)
	fmt.Println("Elapsed time for MergeSort = ", elapsed)
	fmt.Println("Is sorted: ", IsSorted(result))
}
