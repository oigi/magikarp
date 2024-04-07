package util

import "strconv"

/*
以下不等式是否成立：
t1 - t2 >= diff
*/
func JudgeTimeDiff(t1 int64, t2 string, diff int64) bool {
	t2_i, _ := strconv.Atoi(t2)
	t2_i64 := int64(t2_i)
	return t1-t2_i64 >= diff
}
