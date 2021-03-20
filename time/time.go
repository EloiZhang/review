package main

import (
	"fmt"
	"time"
)

func main() {
	t, _ := time.Parse(
		time.RFC3339,
		"2015-12-09T14:52:47Z")
	fmt.Println(t)

	time2 := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local")    //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, time2, loc)
	timestamp := tmp.Unix()    //转化为时间戳 类型是int64
	fmt.Println(timestamp)
}
