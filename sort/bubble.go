package sort

import "errors"

// Bubble sort algorithm
func Bubble(src []interface{}, c Comparable) ([]interface{}, error) {
	if nil == c {
		return nil, errors.New("no comparable")
	}
	totalLen := len(src)
	if totalLen <= 1 {
		return src, nil
	}
	for i := 0; i < totalLen; i++ {

		for j := 0; j < totalLen; j++ {
			result, err := c(src[i], src[j])
			if nil != err {
				return nil, err
			}
			if result == -1 {
				src[i], src[j] = src[j], src[i]
			}
		}
	}
	return src, nil
}
