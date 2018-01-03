package main

import (
	"fmt"
	"reflect"
)

type (
	height float64
	width  int64
)

func variadicInts(numbers ...int) {
	fmt.Print(numbers, " ")
}

func variadicFloats(floats ...float64) {
	fmt.Print(floats, " ")
}

func variadicSum(numbers ...int) {
	total := 0
	for _, num := range numbers {
		total += num
	}

	fmt.Println(total)
}

func variadicSumFloats(numbers ...float64) {
	var total float64
	for _, num := range numbers {
		total += num
	}

	fmt.Println(total)
}

func vTask(taskName string, i ...interface{}) {
	fmt.Println(taskName)
	fmt.Println(i)

	for _, y := range i {
		fmt.Println(y, reflect.TypeOf(y))
	}
}

func cClose() func() (int, int) {
	x := 0
	z := 4
	return func() (int, int) {
		x++
		z += x
		return x, z
	}
}

func main() {
	n := []int{1, 10, 1000}
	variadicInts(n...)
	variadicSum(n...)
	variadicSum()

	f := []float64{1.1, 10.10, 1000.5}
	variadicFloats(f...)
	variadicFloats(10.1)
	variadicFloats()
	variadicSumFloats(f...)

	var customHeight height = 20.1
	var customWidth width = 10

	vTask("hahaha", 10, 15, 20.5, customHeight, customWidth)

	cc := cClose()
	for index := 0; index < 10; index++ {
		fmt.Println(cc())
	}
}
