package utils

import "strconv"

func ParseUint64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

//func ConvertSliceToSet(slice []interface{}) map[interface{}]bool {
//	var set = make(map[interface{}]bool)
//	for _, item := range slice {
//		set[item] = true
//	}
//	return set
//}
