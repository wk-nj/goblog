package types

import (
	"goblog/pkg/logger"
	"strconv"
)

// Int64ToString 将 int64 转换为 string
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// Uint64ToString 将 uint64 转换为 string
func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

func StringToUint64(str string) uint64 {
	ui,err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		logger.LogError(err)
	}
	return ui
}