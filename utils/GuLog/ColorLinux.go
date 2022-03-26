/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/3/11 22:28
  @Notice:  Linux的处理
*/

package GuLog

import "bytes"

type ColorFun func(text string) string

var right = func() []byte {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteByte(0x1B)
	buffer.WriteString("[0m")
	return buffer.Bytes()
}()

var ColorMap = map[int]ColorFun{
	red: func(text string) string {
		buffer := bytes.NewBuffer(nil)
		buffer.WriteByte(0x1B)
		buffer.WriteString("[1;31m")
		buffer.WriteString(text)
		buffer.Write(right)
		return buffer.String()
	},
	blue: func(text string) string {
		buffer := bytes.NewBuffer(nil)
		buffer.WriteByte(0x1B)
		buffer.WriteString("[1;34m")
		buffer.WriteString(text)
		buffer.Write(right)
		return buffer.String()
	},
	yellow: func(text string) string {
		buffer := bytes.NewBuffer(nil)
		buffer.WriteByte(0x1B)
		buffer.WriteString("[1;43;30m")
		buffer.WriteString(text)
		buffer.Write(right)
		return buffer.String()
	},
	green: func(text string) string {
		buffer := bytes.NewBuffer(nil)
		buffer.WriteByte(0x1B)
		buffer.WriteString("[1;32m")
		buffer.WriteString(text)
		buffer.Write(right)
		return buffer.String()
	},
	grey: func(text string) string {
		buffer := bytes.NewBuffer(nil)
		buffer.WriteByte(0x1B)
		buffer.WriteString("[1;47;30m")
		buffer.WriteString(text)
		buffer.Write(right)
		return buffer.String()
	},
}
