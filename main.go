package main

import (
	"cattlelogger/logger"
)

func main() {
	l, err := logger.NewLog("logs", "log.log", 10, 10, 10, true, "info")
	if err != nil {
		panic(err)
	}
	l.Sugar().Panicf("你好：%s", "world")
}
