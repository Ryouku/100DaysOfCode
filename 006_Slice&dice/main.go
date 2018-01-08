package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	// Arrays

	// create float64 type array, cause int is far too popular..
	var a [10]float64
	fmt.Println("Empty array:", a, reflect.TypeOf(a))

	// populate array with data..
	for index := 0; index < len(a); index++ {
		// some dummy number
		x := 3.14159265359

		// do some float rounding / conversion magic
		// TODO: optimise this bit
		f := strconv.FormatFloat(float64(x*float64(index+1)), 'f', 3, 64)
		ff, _ := strconv.ParseFloat(f, 64)

		// assign newly created float number into float array
		a[index] = ff
	}

	fmt.Println("Populated array:", a, reflect.TypeOf(a))

	// small inline array initialisation using literal notation
	// and get back to the int's :)
	aa := [2]int{10, 20}
	fmt.Println("Small array:", aa, reflect.TypeOf(aa))

	// multi-dimensional-arrayish-data
	// init empty array
	var ma [2][3][4]int
	fmt.Println("ma:", ma)
	fmt.Println("ma len:", len(ma))
	fmt.Println("ma len[1]:", len(ma[1]))
	fmt.Println("ma len[1][1]:", len(ma[1][1]))

	// do some pre-population
	for y := 0; y < len(ma); y++ {
		for r := 0; r < len(ma[y]); r++ {
			for e := 0; e < len(ma[y][r]); e++ {
				ma[y][r][e] = (y + r + e) * (r + 1)
			}
		}
	}

	// some data were inserted...
	fmt.Println("Populated ma:", ma)

	// SLICES

	s := make([]string, 4)
	fmt.Println("Slice s:", s, reflect.TypeOf(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"

	fmt.Println("Slice s:", s, reflect.TypeOf(s))

	sl := s[(len(s) / 2):(len(s)/2 + 1)]
	fmt.Println(sl)

	sl = s[:(len(s)/2 + 1)]
	fmt.Println(sl)

	sl = s[(len(s)/2 + 1):]
	fmt.Println(sl)

	// slices slice
	twoS := make([][]int, 2)
	fmt.Println("Two dim-slice:", twoS, reflect.TypeOf(twoS))

	for o := 0; o < len(twoS); o++ {
		n := o + 1
		twoS[o] = make([]int, n)

		for q := 0; q < n; q++ {
			twoS[o][q] = o + q
		}
	}

	fmt.Println("Two dim-slice w'data:", twoS, reflect.TypeOf(twoS))

	// MAPS

	invoices := make(map[int]float64)

	// acconts and their invoice amount
	invoices[4563728] = 134.56
	invoices[2986480] = 999.99

	fmt.Println("Map:", invoices, reflect.TypeOf(invoices))

	i1 := invoices[4563728]
	fmt.Println("Map i1:", i1, reflect.TypeOf(i1))

	// paid some extra...
	invoices[4563728] = 18934.56

	i1 = invoices[4563728]
	fmt.Println("Map i1:", i1, reflect.TypeOf(i1))

	invoiceList := make(map[int]map[int]float64)
	fmt.Println("Map:", invoiceList, reflect.TypeOf(invoiceList))

	// little bit annoying to do this.., as I've already declared it at line 113
	invoiceList[4563728] = map[int]float64{}
	invoiceList[4563728][1] = 12.30
	invoiceList[4563728][2] = 112.30
	invoiceList[4563728][3] = 132.30

	fmt.Println("Map:", invoiceList, reflect.TypeOf(invoiceList))

	// data manipulations

	// dynamic size array
	x := [...]int{1, 2, 3, 4, 5, 6, 7, 10}
	fmt.Println("Array x", x, reflect.TypeOf(x))
	fmt.Println(cap(x))

	// slice from array
	c := x[5:]
	fmt.Println("Array c", c, reflect.TypeOf(c), cap(c))

	// modify original(!) array value by using slice
	c[1] = 33
	fmt.Println("Array x", x, reflect.TypeOf(x))

	// copy of array
	vv := &x
	vv[2] = 1000
	fmt.Println("Array vv", vv, reflect.TypeOf(vv))

}
