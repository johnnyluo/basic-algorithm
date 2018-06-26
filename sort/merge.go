package sort

func merge(left, right []interface{}, c Comparable) ([]interface{}, error) {
	lenLeft := len(left)
	lenRight := len(right)
	result := make([]interface{}, 0, lenLeft+lenRight)
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...), nil
		}
		if len(right) == 0 {
			return append(result, left...), nil
		}
		r, err := c(left[0], right[0])
		if nil != err {
			return nil, err
		}
		if r <= 0 {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}

	}
	return result, nil
}

// Merge merge sort implementation
func Merge(src []interface{}, c Comparable) ([]interface{}, error) {
	totalLen := len(src)
	if totalLen <= 1 {
		return src, nil
	}
	middle := totalLen / 2
	left, err := Merge(src[:middle], c)
	if nil != err {
		return nil, err
	}
	right, err := Merge(src[middle:], c)
	if nil != err {
		return nil, err
	}

	return merge(left, right, c)
}
