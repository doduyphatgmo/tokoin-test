package utils

import "strconv"

func ParseUint64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

func ConvertStrListToMap(strList []string, strKeymap map[string]bool) {
	for _, str := range strList {
		strKeymap[str] = true
	}
}
