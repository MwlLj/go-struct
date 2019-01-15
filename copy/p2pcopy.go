package copy

import (
	"fmt"
	"reflect"
)

func OrderCopy(src interface{}, dst interface{}) {
	srcType := reflect.TypeOf(src)
	dstType := reflect.TypeOf(dst)
	srcCount := srcType.NumField()
	dstCount := dstType.NumField()
	minCount := srcCount
	if dstCount < minCount {
		minCount = dstCount
	}
	if minCount == 0 {
		return
	}
	srcValue := reflect.ValueOf(src)
	dstValue := reflect.ValueOf(dst)
	for i := 0; i < minCount; i++ {
		if !dstValue.CanSet() || !dstValue.IsValid() || dstValue.IsNil() {
			fmt.Println("can't set")
			continue
		}
		srcKind := srcType.Kind()
		if srcKind == reflect.Ptr {
		} else if srcKind == reflect.Array {
		} else if srcKind == reflect.Map {
		} else {
			dv := dstValue.Elem().FieldByIndex([]int{i})
			sv := srcValue.Elem().FieldByIndex([]int{i}).Elem()
			dv.Set(sv)
		}
	}
}
