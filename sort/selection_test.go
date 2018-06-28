package sort_test

import (
	"reflect"
	"testing"

	"github.com/johnnyluo/basic-algorithm/sort"
)

func TestSelectionSort(t *testing.T) {
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
			result, err := sort.Selection(item.src, item.c)
			if nil != err {
				st.Error(err)
				return
			}
			st.Logf("result:%+v\n", result)
			if !reflect.DeepEqual(item.expectedResult, result) {
				st.Errorf("we expected the result to be %+v, however we got %+v", item.expectedResult, result)
			}
		})
	}
}
