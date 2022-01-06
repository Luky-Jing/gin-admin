package conv

import "strconv"

func ParseStringSliceToUint64(s []string) []uint64 {
	iv := make([]uint64, len(s))
	for i, v := range s {
		iv[i], _ = strconv.ParseUint(v, 10, 64)
	}
	return iv
}

func ParseStringSliceToInt64(s []string) []int64 {
	iv := make([]int64, len(s))
	for i, v := range s {
		iv[i], _ = strconv.ParseInt(v, 10, 64)
	}
	return iv
}
