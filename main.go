package main

import (
	"fmt"
	"sort"
)

// Implement function that returns an average value of array
func task1(input *[]int) float64 {
	if len(*input) == 0 {
		return float64(0)
	}
	sum := 0
	for _, v := range *input {
		sum += v
	}
	return float64(sum) / float64(len(*input))
}

// function that returns the copy of the original slice in reverse order. The type of elements is int64.
func task2(input *[]int64) []int64 {
	output := []int64{}
	if len(*input) == 0 {
		return output
	}
	for i := len(*input) - 1; i >= 0; i-- {
		output = append(output, (*input)[i])
	}
	return output
}

// function that returns map values sorted in order of increasing keys.
func task3(input *map[int]string) []string {
	output := []string{}
	keys := []int{}
	for k := range *input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		output = append(output, (*input)[k])
	}
	return output
}

func main() {
	t1Input := []int{1, 2, 3, 4, 5, 6}
	task1Res := task1(&t1Input)
	fmt.Printf("Task 1 result: %f\n", task1Res)

	t2Input := []int64{1, 2, 5, 15}
	task2Res := task2(&t2Input)
	fmt.Printf("Task 1 result: %v\n", task2Res)

	fmt.Println()

	t3Input := map[int]string{2: "a", 0: "b", 1: "c"}
	task3Res := task3(&t3Input)
	fmt.Printf("Task 1 result: %v\n", task3Res)

	fmt.Println("done")
}
