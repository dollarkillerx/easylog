/**
 * @Author: DollarKillerX
 * @Description: caller 获取堆栈
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午5:47 2020/1/3
 */
package easylog

import "runtime"

func getCaller(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0
	}
	n := 0

	// 获取包名
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line
}
