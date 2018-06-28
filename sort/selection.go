package sort

import "errors"

// Selection sort
// o(n*n)
func Selection(src []interface{}, c Comparable) ([]interface{}, error) {
	if nil == c {
		return nil, errors.New("no comparable")
	}
	totalLen := len(src)
	if totalLen <= 1 {
		return src, nil
	}
	result := make([]interface{}, 0, totalLen)
	for i := 0; i < totalLen; i++ {
		idx, err := findMin(src, c)
		if nil != err {
			return nil, err
		}
		result = append(result, src[idx])
		src = append(src[:idx], src[idx+1:]...)
	}
	return result, nil
}
func findMin(src []interface{}, c Comparable) (int, error) {
	if len(src) == 1 {
		return 0, nil
	}
	idx := 0
	value := src[idx]
	for i := 1; i < len(src); i++ {
		r, err := c(value, src[i])
		if nil != err {
			return 0, err
		}
		if r == 1 {
			idx = i
			value = src[i]
		}
	}
	return idx, nil
}
