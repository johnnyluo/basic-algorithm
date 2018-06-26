package sort

// Cocktail shake
func Cocktail(src []interface{}, c Comparable) ([]interface{}, error) {
	totalLen := len(src)
	if totalLen <= 1 {
		return src, nil
	}
	for i := 0; i < totalLen/2; i++ {
		left := 0
		right := totalLen - 1
		for left <= right {
			result, err := c(src[left], src[left+1])
			if nil != err {
				return nil, err
			}
			if result == 1 {
				src[left], src[left+1] = src[left+1], src[left]
			}
			left++

			result, err = c(src[right], src[right-1])
			if nil != err {
				return nil, err
			}
			if result == -1 {
				src[right], src[right-1] = src[right-1], src[right]
			}
			right--
		}
	}
	return src, nil
}
