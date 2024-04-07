package util

import "strings"

// FillUserId 填充用户id
func FillUserId(userId string) string {
	return LFill(userId, 20)
}

func LFill(num string, bit int) string {
	if len(num) >= bit {
		return num
	}

	return strings.Repeat("0", bit-len(num)) + num
}
