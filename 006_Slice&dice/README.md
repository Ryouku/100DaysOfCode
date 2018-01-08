## Motivation

This time I wanted to play with arrays, slices and maps, as this is the basics of data containment and manipulation using hashes, arrays, slices and etc.

## Expectations

To get familiar with such data types and to understant, when and what is more suitable for the task at hand.

## Result

```shell
» go run main.go
Empty array: [0 0 0 0 0 0 0 0 0 0] [10]float64
Populated array: [3.142 6.283 9.425 12.566 15.708 18.85 21.991 25.133 28.274 31.416] [10]float64
Small array: [10 20] [2]int
ma: [[[0 0 0 0] [0 0 0 0] [0 0 0 0]] [[0 0 0 0] [0 0 0 0] [0 0 0 0]]]
ma len: 2
ma len[1]: 3
ma len[1][1]: 4
Populated ma: [[[0 1 2 3] [2 4 6 8] [6 9 12 15]] [[1 2 3 4] [4 6 8 10] [9 12 15 18]]]
Slice s: [   ] []string
Slice s: [a b c d] []string
[c]
[a b c]
[d]
Two dim-slice: [[] []] [][]int
Two dim-slice w'data: [[0] [1 2]] [][]int
Map: map[4563728:134.56 2986480:999.99] map[int]float64
Map i1: 134.56 float64
Map i1: 18934.56 float64
Map: map[] map[int]map[int]float64
Map: map[4563728:map[1:12.3 2:112.3 3:132.3]] map[int]map[int]float64
Array x [1 2 3 4 5 6 7 10] [8]int
8
Array c [6 7 10] []int 3
Array x [1 2 3 4 5 6 33 10] [8]int
Array vv &[1 2 1000 4 5 6 33 10] *[8]int
```