package tmpl

import (
	"fmt"
	"strconv"
	"time"
)

// StartCompute compute start time,return format string
// 2020-11-13 10:00
func StartCompute(start uint8) string {
	date := time.Now().Format("2006-01-02")
	return fmt.Sprintf("%s %s:00", date, strconv.Itoa(int(start)))
}

// EndCompute compute end time,return format string,like StartCompute
func EndCompute(end uint8) string {
	date := time.Now().Format("2006-01-02")
	return fmt.Sprintf("%s %s:00", date, strconv.Itoa(int(end)))
}

// DurationCompute compute holiday duration
func DurationCompute(start, end uint8) string {
	return fmt.Sprintf("0天%s时", strconv.Itoa(int(end-start)))
}

// ApplyTime return apply holiday time, default is yesterday 10:00
func ApplyTime() string {
	now := time.Now()
	return fmt.Sprintf("%d-%d-%d 10:00", now.Year(), now.Month(), now.Day())
}
