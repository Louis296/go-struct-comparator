package go_struct_comparator

import "reflect"

func structCompare(a interface{}, b interface{}) map[string][]interface{} {
	DefaultTag := "compare_key"
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
			if ok2 && !reflect.DeepEqual(v1, v2) {
				result[tag.Get(DefaultTag)] = []interface{}{v1, v2}
			}
		}
	}
	return result
}
