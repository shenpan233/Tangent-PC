/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/3/11 22:29
  @Notice:  中间处理
*/

package GuLog

import (
	"bytes"
	"strings"
)

const (
	red = iota
	blue
	green
	yellow
	grey
)

func enterColor(text string, color ColorFun) string {
	list := strings.Split(text, "\n")
	buffer := bytes.NewBufferString("")
	for _, sub := range list {
		buffer.WriteString(color(sub) + "\n")
	}
	return buffer.String()

}
