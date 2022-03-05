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
	"strings"
	"time"
)

func print(Type int8, info, msg string) {
	strLogType := "U"
	var (
		left, right string
	)
	{
		switch Type {
		case warm:
			strLogType = warmSign
			if guLog.color {
				left = fmt.Sprintf("%c[1;43;30m", 0x1B)
				right = fmt.Sprintf("%c[0m", 0x1B)
			}
			break
		case err:
			strLogType = errorSign
			if guLog.color {
				left = fmt.Sprintf("%c[31m", 0x1B)
				right = fmt.Sprintf("%c[0m", 0x1B)
			}
			break
		case debug:
			strLogType = debugSign
			if guLog.color {
				left = fmt.Sprintf("%c[1;47;30m", 0x1B)
				right = fmt.Sprintf("%c[0m", 0x1B)
			}
			break
		case infO:
			strLogType = infoSign
			if guLog.color {
				left = fmt.Sprintf("%c[1;34m", 0x1B)
				right = fmt.Sprintf("%c[0m", 0x1B)
			}
			break
		case notice:
			strLogType = noticeSign
			if guLog.color {
				left = fmt.Sprintf("%c[1;32m", 0x1B)
				right = fmt.Sprintf("%c[0m", 0x1B)
			}
			break
		}
	}
	msg = msgLoader(msg, left, right)
	fmt.Printf("%s %s[%s]  %s %s\n%s\n", time.Now().Format(guLog.timeFormat), left, strLogType, info, right, msg)
}

func msgLoader(msg string, left, right string) string {
	buffer := bytes.NewBufferString("")
	lit := strings.Split(msg, "\n")
	for _, kid := range lit {
		buffer.WriteString(left)
		buffer.WriteString(kid)
		buffer.WriteString(right)
		buffer.WriteString("\n")
	}
	//buffer.WriteString(left)
	//buffer.WriteString(*msg)
	//buffer.WriteString(right)
	return buffer.String()
}
