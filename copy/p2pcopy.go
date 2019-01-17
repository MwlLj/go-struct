package copy

import (
	"fmt"
	"reflect"
)

var _ = fmt.Println

func OrderCopy(src interface{}, dst interface{}) {
	srcType := reflect.TypeOf(src)
	dstType := reflect.TypeOf(dst)
	// fmt.Println(dstType.Kind())
	if srcType.Kind() != reflect.Ptr || dstType.Kind() != reflect.Ptr {
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
			// newValue := reflect.New(dstFieldValue.Type().Elem())
			newValue := reflect.New(dstFieldValue.Type())
			// fmt.Println("2.   ", newValue.Type().Kind())
			// newValueElem := newValue.Elem()
			// fmt.Println(newValue)
			// fmt.Println(srcFieldValue)
			OrderCopy(&srcFieldValue, newValue)
			dstFieldValue.Set(newValue.Elem())
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

func DeepFields(ifaceType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField

	for i := 0; i < ifaceType.NumField(); i++ {
		v := ifaceType.Field(i)
		if v.Anonymous && v.Type.Kind() == reflect.Struct {
			fields = append(fields, DeepFields(v.Type)...)
		} else {
			fields = append(fields, v)
		}
	}

	return fields
}

func StructCopy(DstStructPtr interface{}, SrcStructPtr interface{}) {
	srcv := reflect.ValueOf(SrcStructPtr)
	dstv := reflect.ValueOf(DstStructPtr)
	srct := reflect.TypeOf(SrcStructPtr)
	dstt := reflect.TypeOf(DstStructPtr)
	if srct.Kind() != reflect.Ptr || dstt.Kind() != reflect.Ptr ||
		srct.Elem().Kind() == reflect.Ptr || dstt.Elem().Kind() == reflect.Ptr {
		panic("Fatal error:type of parameters must be Ptr of value")
	}
	if srcv.IsNil() || dstv.IsNil() {
		panic("Fatal error:value of parameters should not be nil")
	}
	srcV := srcv.Elem()
	dstV := dstv.Elem()
	srcfields := DeepFields(reflect.ValueOf(SrcStructPtr).Elem().Type())
	for _, v := range srcfields {
		if v.Anonymous {
			continue
		}
		dst := dstV.FieldByName(v.Name)
		src := srcV.FieldByName(v.Name)
		if !dst.IsValid() {
			continue
		}
		if src.Type() == dst.Type() && dst.CanSet() {
			dst.Set(src)
			continue
		}
		if src.Kind() == reflect.Ptr && !src.IsNil() && src.Type().Elem() == dst.Type() {
			dst.Set(src.Elem())
			continue
		}
		if dst.Kind() == reflect.Ptr && dst.Type().Elem() == src.Type() {
			dst.Set(reflect.New(src.Type()))
			dst.Elem().Set(src)
			continue
		}
	}
	return
}
