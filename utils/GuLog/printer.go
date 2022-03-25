/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		打印
* @Creat:   2021/10/29 0029 22:46
 */

package GuLog

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func dump(color int, Tag, msg string) {
	pc, file, line, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	funcName = funcName[strings.LastIndex(funcName, ".")+1:]
	if guLog.isWindowConsole {
		//颜色反射+goroutine缓解下读写锁的性能问题
		go WinPrint(file, funcName, Tag, msg, line, winColor[linuxColor(color)])
	} else {
		buffer := bytes.NewBufferString(time.Now().Format(guLog.timeFormat))
		buffer.WriteString(" " + file)
		buffer.WriteString(":" + strconv.Itoa(line))
		fun := ColorMap[color]
		buffer.WriteString(fun(" [" + Tag + "] "))
		buffer.WriteString(funcName)
		buffer.WriteString("\n")
		buffer.WriteString(enterColor(msg, fun))
		fmt.Print(buffer)
	}
}

func Clear() {
	os.Stdout.Sync()
}
