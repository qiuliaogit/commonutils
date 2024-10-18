package commonutils

import (
	"testing"
	"time"
)

func Test_DateNum(t *testing.T) {
	timestamp := 1729206000
	tm := time.Unix(int64(timestamp), 0)
	dateNumUnix := FormatDateNum(tm)
	dateNumBJ := FormatDateNumForBeijing(tm)

	t.Log("==> unix", dateNumUnix)
	t.Log("==> beijing", dateNumBJ)

	//  ArraySort<[]int>(lst)
}
