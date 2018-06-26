package sort_test

import (
	"reflect"
	"testing"

	"github.com/johnnyluo/basic-algorithm/sort"
)

func TestMergeSort(t *testing.T) {
	inputs := []struct {
		name           string
		src            []interface{}
		c              sort.Comparable
		expectedResult []interface{}
	}{
		{
			name:           "1",
			src:            []interface{}{9, 0, 8, 7, 6, 5, 4, 3, 2, 1, 12},
			c:              primitiveComparable,
			expectedResult: []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 12},
		},
	}
	for _, item := range inputs {
		t.Run(item.name, func(st *testing.T) {
			result, err := sort.Merge(item.src, item.c)
			if nil != err {
				st.Error(err)
				return
			}
			if !reflect.DeepEqual(result, item.expectedResult) {
				st.Errorf("we expected result to be %+v,however we got %+v", item.expectedResult, result)
			}
		})
	}
}
