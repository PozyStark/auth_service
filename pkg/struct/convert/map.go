package convert

import (
	"fmt"
	"reflect"
)

type ReflectStruct struct {
	Value reflect.Value
	Type reflect.Type
}

func (r ReflectStruct) GetReflects() (reflect.Value, reflect.Type) {
	return r.Value, r.Type
}

func GetReflectValues(s any) (ReflectStruct, error) {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if v.Kind() != reflect.Struct {
		return ReflectStruct{}, fmt.Errorf("expects a type structure for input but received: %s", s)
	}
	return ReflectStruct{Value: v, Type: t}, nil
}

func IsBool(i any) bool {
	switch i.(type) {
	case bool:
		return true
	default:
		return false
	}
}

func StructToMap(s any) (map[string]any, error) {

	resultMap := make(map[string]any)
	rs, err := GetReflectValues(s)

	if err != nil {
		return resultMap, err
	}
	v, t := rs.GetReflects()
	for i:=0; i<v.NumField(); i++ {
		fieldType := t.Field(i)
		fieldValue := v.Field(i)
		// fmt.Println(fieldType,"-",fieldValue)
		resultMap[fieldType.Name] = fieldValue.Interface()
	}

	return resultMap, nil
}

func StructToMapNotNull(s any) (map[string]any, error) {

	resultMap := make(map[string]any)
	rs, err := GetReflectValues(s)

	if err != nil {
		return resultMap, err
	}
	v, t := rs.GetReflects()
	for i:=0; i<v.NumField(); i++ {
		fieldType := t.Field(i)
		fieldValue := v.Field(i)
		if !fieldValue.IsZero(){
			resultMap[fieldType.Name] = fieldValue.Interface()
		}
	}

	return resultMap, nil
}