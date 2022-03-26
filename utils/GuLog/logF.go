package GuLog

import "fmt"

//这样写好像有点笨
var guLog struct {
	color           bool
	timeFormat      string
	isWindowConsole bool
}

func ErrorF(msg string, v ...interface{}) {
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	dump(red, "E", msg)
}

func DebugF(msg string, v ...interface{}) {
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	dump(grey, "D", msg)

}

func WarmF(msg string, v ...interface{}) {
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	dump(yellow, "W", msg)

}

func InfoF(msg string, v ...interface{}) {
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	dump(blue, "I", msg)

}

func NoticeF(msg string, v ...interface{}) {
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	dump(green, "N", msg)
}
