/**
 * @Author: DollarKillerX
 * @Description: easylog.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午5:38 2020/1/3
 */
package easylog

import (
	"fmt"
)

type easyLog struct {
	level string // 登记 INFO ERROR WARNING    // 默认debug 全部打印
}

/**
 * 规则 level 打印登记
 * DEBUG INFO ERROR WARNING
 * DEBUG 为全部都答应详细信息
 * WARNING 其他都为向上打印错误信息
 */
var logc = &easyLog{level: "INFO"}

func SetLevel(level string) {
	logc.level = level
}

func PrintWarning(v ...interface{}) {
	sprintf := fmt.Sprintf(" %c[%d;%d;%dm%s WARNING %c[0m ", 0x1B, 1, 41, 36, "", 0x1B)
	switch logc.level {
	case "WARNING":
		s, i := getCaller(2)
		fmt.Println("time: ", getTimeString(), sprintf, " path: [", s, ":", i, "] ", v)
	default:
		fmt.Println("time: ", getTimeString(), sprintf, v)

	}
}

func PrintError(v ...interface{}) {
	sprintf := fmt.Sprintf(" %c[%d;%d;%dm%s ERROR %c[0m ", 0x1B, 1, 45, 30, "", 0x1B)
	switch {
	case logc.level == "ERROR" || logc.level == "WARNING":
		s, i := getCaller(2)
		fmt.Println("time: ", getTimeString(), sprintf, " path: [", s, ":", i, "] ", v)
	default:
		fmt.Println("time: ", getTimeString(), sprintf, v)
	}
}

func PrintInfo(v ...interface{}) {
	//switch {
	//case logc.level == "ERROR" || logc.level == "WARNING":
	//	sprintf := fmt.Sprintf(" %c[%d;%d;%dm%s INFO %c[0m ", 0x1B, 1, 42, 36, "", 0x1B)
	//	s, i := getCaller(2)
	//	fmt.Println("time: ", getTimeString(), sprintf, " path: ", s, ":", i, " ", v)
	//default:
	//	fmt.Println(v)
	//}
	fmt.Println("time: ", getTimeString(), "  INFO  ", v)
}
