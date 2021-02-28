package utils

import "strconv"

func ParseUint64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}
