package copy

import (
	"fmt"
	"reflect"
)

var _ = fmt.Println

func OrderCopy(src interface{}, dst interface{}) {
	srcType := reflect.TypeOf(src)
	dstType := reflect.TypeOf(dst)
	if srcType.Kind() != reflect.Ptr || dstType.Kind() != reflect.Ptr {
		fmt.Println("----------")
		return
	}
	srcTypeElem := srcType.Elem()
	dstTypeElem := dstType.Elem()
	srcCount := srcTypeElem.NumField()
	dstCount := dstTypeElem.NumField()
	minCount := srcCount
	if dstCount < minCount {
		minCount = dstCount
	}
	if minCount == 0 {
		return
	}
	srcValue := reflect.ValueOf(src).Elem()
	dstValue := reflect.ValueOf(dst).Elem()
	for i := 0; i < minCount; i++ {
		srcFieldType := srcTypeElem.Field(i)
		dstFieldType := dstTypeElem.Field(i)
		srcFieldValue := srcValue.Field(i)
		dstFieldValue := dstValue.Field(i)
		srcFieldTypeKind := srcFieldType.Type.Kind()
		if srcFieldTypeKind != reflect.Ptr && (srcFieldTypeKind != reflect.Struct && srcFieldValue.Type() != dstFieldValue.Type()) {
			// fmt.Println(srcFieldValue.Type())
			// fmt.Println(dstFieldValue.Type())
			return
		}
		if srcFieldTypeKind == reflect.Struct {
			var _ = dstFieldType
			newValue := reflect.New(dstFieldValue.Type())
			newValueElem := newValue.Elem()
			OrderCopy(&srcFieldValue, &newValueElem)
			// dstFieldValue.Set(newValue.Elem())
			reflect.ValueOf(dst).Elem().Field(i).Set(newValue.Elem())
			// OrderCopy(&srcFieldType, &dstFieldType)
			// dstFieldValue.Set(srcFieldValue)
		} else if srcFieldTypeKind == reflect.Ptr {
		} else if srcFieldTypeKind == reflect.Array {
		} else {
			dstFieldValue.Set(srcFieldValue)
			// fmt.Println(dstFieldValue)
		}
	}
}
