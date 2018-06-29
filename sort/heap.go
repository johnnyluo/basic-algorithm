package sort

import (
	"errors"
	"fmt"
)

// Heap sort
func Heap(src []interface{}, c Comparable) ([]interface{}, error) {

	totalLen := len(src)
	if totalLen <= 1 {
		return src, nil
	}
	if nil == c {
		return nil, errors.New("how to compare the value?")
	}
	result, err := heapify(src, c)
	if nil != err {
		return nil, err
	}
	fmt.Println(result)

	for j := (totalLen - 1); j >= 1; j-- {
		result[0], result[j] = result[j], result[0]
		result, err = down(result, c, 0, j-1)
		if nil != err {
			return nil, err
		}
		fmt.Printf("[%d] -- %+v\n", j, result)
	}

	return result, nil
}

func down(src []interface{}, c Comparable, first, last int) ([]interface{}, error) {
	i := first
	for {
		j1 := 2*i + 1
		if j1 >= last || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 <= last {
			r, err := c(src[j2], src[j1])
			if nil != err {
				return nil, err
			}
			if r == 1 {
				j = j2
			}
		}
		r1, err := c(src[j], src[i])
		if nil != err {
			return nil, err
		}
		if r1 <= 0 {
			break
		}
		src[j], src[i] = src[i], src[j]
		i = j // keep going
	}

	return src, nil
}
func heapify(src []interface{}, c Comparable) ([]interface{}, error) {
	totalLen := len(src)
	for i := 0; i < totalLen; i++ {
		result, err := shiftUp(src, c, i)
		if nil != err {
			return nil, err
		}
		src = result
	}
	return src, nil
}

func shiftUp(src []interface{}, c Comparable, idx int) ([]interface{}, error) {
	if idx == 0 {
		return src, nil
	}
	parent := (idx - 1) / 2
	r, err := c(src[idx], src[parent])
	if nil != err {
		return nil, err
	}
	if r == 1 {
		// swap
		src[parent], src[idx] = src[idx], src[parent]
	}
	return shiftUp(src, c, parent)
}
