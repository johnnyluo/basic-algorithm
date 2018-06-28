package sort

import (
	"errors"
)

// Insert sort
// too much allocation
func Insert(src []interface{}, c Comparable) ([]interface{}, error) {
	if nil == c {
		return nil, errors.New("no comparable")
	}
	totalLen := len(src)
	if totalLen <= 1 {
		return src, nil
	}
	result := make([]interface{}, 0, totalLen)
	result = append(result, src[0])
	for i := 1; i < totalLen; i++ {
		flag := false
		for j := i - 1; j >= 0; j-- {
			//fmt.Printf("i:%d,j:%d,Result:%+v\n", i, j, result)
			r, err := c(src[i], result[j])
			if nil != err {
				return nil, err
			}
			if r >= 0 {
				result = append(append(result[:j], src[i]), result[j:]...)
				flag = true
				break
			}
		}
		if !flag {
			result = append([]interface{}{src[i]}, result[:]...)
		}
	}
	return result, nil
}

// Insert1 second implementation
// maximum time complexity o(n*n)
func Insert1(src []interface{}, c Comparable) ([]interface{}, error) {
	if nil == c {
		return nil, errors.New("no comparable")
	}
	totalLen := len(src)
	if totalLen <= 1 {
		return src, nil
	}
	result := make([]interface{}, 0, totalLen)
	result = append(result, src[0])
	for i := 1; i < totalLen; i++ {
		result = append(result, src[i])
		for j := i; j > 0; j-- {
			r, err := c(result[j], result[j-1])
			if nil != err {
				return nil, err
			}
			if r == -1 {
				result[j], result[j-1] = result[j-1], result[j]
			}
		}
	}
	return result, nil
}
