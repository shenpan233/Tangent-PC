package GuLog

import (
	"fmt"
)

//这样写好像有点笨
var guLog struct {
	color      bool
	timeFormat string
}

func Error(info, format string, v ...interface{}) {
	print(err, info, fmt.Sprintf(format, v...))
}

func Debug(info, format string, v ...interface{}) {
	print(debug, info, fmt.Sprintf(format, v...))
}

func Warm(info, format string, v ...interface{}) {
	print(warm, info, fmt.Sprintf(format, v...))
}

func Info(Info, format string, v ...interface{}) {
	print(infO, Info, fmt.Sprintf(format, v...))
}

func Notice(info, format string, v ...interface{}) {
	print(notice, info, fmt.Sprintf(format, v...))
}
