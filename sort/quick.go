package sort

import "fmt"

// Comparable compare stuff
// when a < b return -1
// when a == b return 0
// when a > b return 1
type Comparable func(a interface{}, b interface{}) (int, error)

// Quick sort
func Quick(input []interface{}, comparable Comparable) ([]interface{}, error) {
	totalLengh := len(input)
	if totalLengh <= 1 {
		return input, nil
	}

	median := input[totalLengh/2]
	if median == nil {
		fmt.Println(input)
	}
	left := make([]interface{}, 0, totalLengh)
	right := make([]interface{}, 0, totalLengh)
	middle := make([]interface{}, 0, totalLengh)
	for _, item := range input {
		result, err := comparable(item, median)
		if nil != err {
			return nil, err
		}
		switch result {
		case -1:
			left = append(left, item)
		case 0:
			middle = append(middle, item)
		case 1:
			right = append(right, item)
		}
	}
	left, err := Quick(left, comparable)
	if nil != err {
		return nil, err
	}
	right, err = Quick(right, comparable)
	if nil != err {
		return nil, err
	}
	return append(append(left, middle...), right...), nil
}
