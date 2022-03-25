/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/3/11 22:35
  @Notice:  不带格式化的自动增量
*/

package GuLog

import "fmt"

func Error(v ...interface{}) {
	msg := fmt.Sprint(v...)
	dump(red, "E", msg)
}
func Debug(v ...interface{}) {
	msg := fmt.Sprint(v...)
	dump(grey, "D", msg)
}

func Warm(v ...interface{}) {
	msg := fmt.Sprint(v...)
	dump(yellow, "W", msg)
}

func Info(v ...interface{}) {
	msg := fmt.Sprint(v...)
	dump(blue, "I", msg)
}

func Notice(v ...interface{}) {
	msg := fmt.Sprint(v...)
	dump(green, "N", msg)
}
