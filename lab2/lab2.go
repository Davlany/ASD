package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type TestData struct {
	SwapNum    int
	CompareNum int
	TimeStamp  time.Duration
}

func main() {
	arr := RandomizeArray(1000000)
	td := ShellSort(arr, 1000000)
	fmt.Println(arr)
	fmt.Println(td)
}

func RandomizeArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.IntN(10000000)
	}
	return arr
}

func ShellSort(data []int, size int) TestData {
	start := time.Now()
	var testData TestData
	for i := size / 2; i > 0; i /= 2 {
		for j := i; j < size; j++ {
			testData.CompareNum++
			for k := j - i; k >= 0 && data[k] > data[k+i]; k -= i {
				temp := data[k]
				data[k] = data[k+i]
				data[k+i] = temp
				testData.SwapNum++
			}
		}
	}
	elapsed := time.Since(start)
	testData.TimeStamp = elapsed
	return testData
}

func SelectionSort(data []int) TestData {
	var testData TestData
	listLen := len(data)
	start := time.Now()
	for i := 0; i < listLen-1; i++ {
		minIndex := i
		for j := i; j < listLen; j++ {
			testData.CompareNum++
			if data[j] < data[minIndex] {
				minIndex = j
				testData.SwapNum++
			}
		}
		di := data[i]
		data[i] = data[minIndex]
		data[minIndex] = di
	}
	elapsed := time.Since(start)
	testData.TimeStamp = elapsed
	return testData
}
