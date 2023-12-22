package utils

import "strconv"

func ParseInt(param string) int64 {
	i, _ := strconv.ParseInt(param, 10, 64)
	return i
}
