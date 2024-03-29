package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const size = 50_000_000
const threshold = 5000

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

// CONCURRENT QUICKSORT
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

func Partition[T DataType](data []T) int {
	data[len(data)/2], data[0] = data[0], data[len(data)/2]
	pivot := data[0]
	mid := 0
	i := 1
	for i < len(data) {
		if data[i] < pivot {
			mid += 1
			data[i], data[mid] = data[mid], data[i]
		}
		i += 1
	}
	data[0], data[mid] = data[mid], data[0]
	return mid
}

func ConcurrentQuicksort[T DataType](data []T, wg *sync.WaitGroup) {
	for len(data) >= 30 {
		mid := Partition(data)
		var portion []T
		if mid < len(data)/2 {
			portion = data[:mid]

			data = data[mid+1:]
		} else {
			portion = data[mid+1:]
			data = data[:mid]
		}
		if len(portion) > threshold {
			wg.Add(1)
			go func(data []T) {
				defer wg.Done()
				ConcurrentQuicksort(data, wg)
			}(portion)
		} else {
			ConcurrentQuicksort(portion, wg)
		}
	}
	InsertSort(data)
}

func QSort[T DataType](data []T) {
	var wg sync.WaitGroup
	ConcurrentQuicksort(data, &wg)
	wg.Wait()
}

func main() {
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}

	start := time.Now()
	QSort[float64](data)
	elapsed := time.Since(start)
	fmt.Println("Elapsed time for concurrent quicksort = ", elapsed)
	fmt.Println("Is sorted: ", IsSorted(data))
}
