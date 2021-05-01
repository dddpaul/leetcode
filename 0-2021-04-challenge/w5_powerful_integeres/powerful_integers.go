package main

import (
	"fmt"
	"math"
)

/**
 * https://leetcode.com/explore/challenge/card/april-leetcoding-challenge-2021/597/week-5-april-29th-april-30th/3726/
 */

type Solver interface {
	PowerfulIntegers(x int, y int, bound int) []int
}

type DirectSolver struct{}

func (ds DirectSolver) PowerfulIntegers(x int, y int, bound int) []int {

	log := func(a int, b int) int {
		if a == 1 {
			return 0
		}
		return int(math.Floor(math.Log(float64(b)) / math.Log(float64(a))))
	}

	set := make(map[int]interface{})
	exists := struct{}{}

	max_i := log(x, bound)
	max_j := log(y, bound)
	fmt.Printf("max_i = %d, max_j = %d", max_i, max_j)

	for i := 0; i <= max_i; i++ {
		for j := 0; j <= max_j; j++ {
			if num := int(math.Pow(float64(x), float64(i)) + math.Pow(float64(y), float64(j))); num <= bound {
				fmt.Printf("i = %d, j = %d, num = %d\n", i, j, num)
				set[num] = exists
			}
		}
	}

	arr := make([]int, 0)
	for k := range set {
		arr = append(arr, k)
	}

	return arr
}
