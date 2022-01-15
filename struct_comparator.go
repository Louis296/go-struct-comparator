package go_struct_comparator

import (
	"reflect"
)

const (
	DefaultTag = "compare_key"
)

func structCompare(a interface{}, b interface{}, upper string) map[string][]interface{} {
	typeA := reflect.TypeOf(a)
	typeB := reflect.TypeOf(b)
	valueA := reflect.ValueOf(a)
	valueB := reflect.ValueOf(b)
	tagMap := make(map[string]interface{})
	result := make(map[string][]interface{})
	for i := 0; i < typeA.NumField(); i++ {
		tag := typeA.Field(i).Tag
		_, ok := tag.Lookup(DefaultTag)
		if ok {
			tagMap[tag.Get(DefaultTag)] = valueA.Field(i).Interface()
		}
	}
	for i := 0; i < typeB.NumField(); i++ {
		tag := typeB.Field(i).Tag
		_, ok := tag.Lookup(DefaultTag)
		if ok {
			v1 := valueB.Field(i).Interface()
			v2, ok2 := tagMap[tag.Get(DefaultTag)]
			if typeA.Field(i).Type.Kind() == reflect.Struct {
				mergeMap(result, structCompare(v1, v2, generateKey(upper, tag.Get(DefaultTag))))
			} else if typeA.Field(i).Type.Kind() == reflect.Array || typeA.Field(i).Type.Kind() == reflect.Slice {
				mergeMap(result, arrayCompare(v1, v2, generateKey(upper, tag.Get(DefaultTag))))
			} else if ok2 && !reflect.DeepEqual(v1, v2) {
				result[generateKey(upper, tag.Get(DefaultTag))] = []interface{}{v1, v2}
			}
		}
	}
	return result
}

// todo
func arrayCompare(a, b interface{}, upper string) map[string][]interface{} {
	result := make(map[string][]interface{})
	return result
}

func mergeMap(a, b map[string][]interface{}) map[string][]interface{} {
	for k, v := range b {
		a[k] = v
	}
	return a
}

func Compare(a, b interface{}) map[string][]interface{} {
	return structCompare(a, b, "")
}

func generateKey(upper, key string) string {
	if upper == "" {
		return key
	} else {
		return upper + "-" + key
	}
}
