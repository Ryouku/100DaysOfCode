## Motivation

This time, I wanted to experiment with Go func types: variadics, closures and multiple return types. Some of this concepts were uncommon in PHP language, therefore I'd like to get familiar with this concept.

## Expectations

To try and understand the possibilities of various Go func types.

## Result

```shell
» go run main.go
[1 10 1000] 1011
0
[1.1 10.1 1000.5] [10.1] [] 1011.7
hahaha
[10 15 20.5 20.1 10]
10 int
15 int
20.5 float64
20.1 main.height
10 main.width
1 5
2 7
3 10
4 14
5 19
6 25
7 32
8 40
9 49
10 59
```