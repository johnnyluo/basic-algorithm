package main

import (
	"errors"
	"fmt"

	"github.com/johnnyluo/basic-algorithm/sort"
)

func main() {
	input := []interface{}{5, 4, 3, 2, 1, 6, 7, 8, 9, 22, 23, 24, 89, 122, 12, 13, 55, 34, 233, 234}
	result, err := sort.Heap(input, primitiveComparable)
	if nil != err {
		panic(err)
	}
	fmt.Println(result)
}
func primitiveComparable(a interface{}, b interface{}) (int, error) {
	if a == b {
		return 0, nil
	}
	switch a.(type) {
	case int:
		return intComparable(a.(int), b.(int)), nil
	case uint32:
		return uint32Comparable(a.(uint32), b.(uint32)), nil
	}
	return -2, errors.New("doesn't suppot type")
}
func uint32Comparable(a, b uint32) int {
	if a == b {
		return 0
	}
	if a > b {
		return 1
	}
	return -1
}

func intComparable(a, b int) int {
	if a == b {
		return 0
	}
	if a > b {
		return 1
	}
	return -1
}
