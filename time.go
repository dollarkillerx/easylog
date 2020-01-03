/**
 * @Author: DollarKillerX
 * @Description: time 获取time
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午6:25 2020/1/3
 */
package easylog

import (
	"time"
)

var (
	// 日期str模板
	timeLayout = "2006-01-02 15:04:05"
	// 时区设置
	location, _ = time.LoadLocation("Asia/Shanghai")
)

func getTimeString() string {
	return time.Unix(time.Now().In(location).Unix(), 0).Format(timeLayout)
}
