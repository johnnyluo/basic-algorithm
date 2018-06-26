package sort_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/johnnyluo/basic-algorithm/sort"
)

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
func TestQuickSort(t *testing.T) {
	inputs := []struct {
		name           string
		src            []interface{}
		c              sort.Comparable
		expectedResult []interface{}
	}{
		{
			name:           "first",
			src:            []interface{}{9, 8, 7, 6, 5, 4, 3, 2, 1},
			c:              primitiveComparable,
			expectedResult: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:           "second",
			src:            []interface{}{9, 8, 7, 7, 5, 4, 3, 2, 1},
			c:              primitiveComparable,
			expectedResult: []interface{}{1, 2, 3, 4, 5, 7, 7, 8, 9},
		},
	}
	for _, item := range inputs {
		t.Run(item.name, func(st *testing.T) {
			result, err := sort.Quick(item.src, item.c)
			if nil != err {
				st.Error(err)
				return
			}
			if !reflect.DeepEqual(item.expectedResult, result) {
				st.Errorf("we expected the result to be %+v, however we got %+v", item.expectedResult, result)
			}
		})
	}
}
