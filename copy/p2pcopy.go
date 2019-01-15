package copy

import (
	"fmt"
	"reflect"
)

func OrderCopy(src interface{}, dst interface{}) {
	srcType := reflect.TypeOf(src)
	dstType := reflect.TypeOf(dst)
	srcCount := srcType.Elem().NumField()
	dstCount := dstType.Elem().NumField()
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
		dv := dstValue.Elem().Field(i)
		sv := srcValue.Elem().Field(i)
		if !dv.IsValid() || dv.IsNil() || !dv.CanSet() {
			fmt.Println("can't set")
			continue
		}
		srcKind := srcType.Kind()
		if srcKind == reflect.Ptr {
		} else if srcKind == reflect.Array {
		} else if srcKind == reflect.Map {
		} else {
			dv.Set(sv)
		}
	}
}
