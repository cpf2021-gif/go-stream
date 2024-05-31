package main

import (
	"fmt"
	"strconv"

	"github.com/cpf2021-gif/go-stream"
)

type student struct {
	name string
	age  int
}

func main() {
	nums := []int{
		1, 1, 9, 4, 4, 6, 3, 5, 15,
	}

	// stream[[]E, E, V], E -> V
	it := stream.StreamOf[[]int, int, string](nums).
		Filter(func(i int, v int) bool {
			return v%2 == 1
		}). // -> 1, 1, 9, 3, 5, 15
		Sorted(func(a, b any) bool {
			return a.(int) < b.(int)
		}). // -> 1, 1, 3, 5, 9, 15
		Distinct(func(v any) any {
			return v.(int) % 5
		}).      // -> 1, 3, 5, 9
		Tail(3). // -> 3, 5, 9
		Map(func(i, v int) string {
			return strconv.Itoa(v)
		}) // -> [3, 5, 9]

	it.ForEach(func(i int, s string) {
		fmt.Println(i, s)
	})

	// Chunk
	fmt.Println(
		stream.Chunk[[]string, string, string](it.ToSlice(), 2).
			ToSlice()) // -> [[3, 5], [9]]

	// Groupby
	stream.GroupBy[[]int, int, int](nums, func(i, v int) any {
		return v % 3
	}).
		Map(func(i int, v []int) int {
			sum := 0
			for _, num := range v {
				sum += num
			}
			return sum
		}).
		ForEach(func(i int, v int) {
			fmt.Println(i, v)
		})

	// []struct
	students := []student{
		{
			name: "xpf",
			age:  43,
		},
		{
			name: "xypf",
			age:  11,
		},
		{
			name: "xyg",
			age:  22,
		},
	}

	stream.StreamOf[[]student, student, student](students).
		Sorted(func(a, b any) bool {
			return a.(student).age < b.(student).age
		}).
		ForEach(func(i int, s student) {
			fmt.Println(i, s)
		})
}
