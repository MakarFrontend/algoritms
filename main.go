package main

import (
	"GoAlgoritms/sorting"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*Нужно для пробных запусков*/
func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	testSlice := createSlice(10000)
	mainStart := time.Now()
	go func() {
		start := time.Now()
		sorting.FastSort(testSlice)
		fmt.Println("Fast: ", time.Since(start).Milliseconds())
		wg.Done()
	}()

	go func() {
		start := time.Now()
		sorting.SortBuble(testSlice)
		fmt.Println("Buble: ", time.Since(start).Milliseconds())
		wg.Done()
	}()

	go func() {
		start := time.Now()
		sorting.SelectionSort(testSlice)
		fmt.Println("Selection: ", time.Since(start).Milliseconds())
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Общее время: ", time.Since(mainStart).Milliseconds())
}

/*Для создания слайса*/
func createSlice(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = rand.Intn(10000)
	}
	return res
}
