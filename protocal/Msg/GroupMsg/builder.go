/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/26 12:27
  @Notice:	_消息结构构造器
*/

package GroupMsg

import "fmt"

const (
	FormatPic = "[Pic="
	FormatAt  = "[At="
	FormatEnd = "]"
)

func buildPic(guid string) string {
	return FormatPic + guid + FormatEnd
}

func buildAt(uin uint32) string {
	if uin == 0 {
		return FormatAt + "All" + FormatEnd
	}
	return fmt.Sprintf(FormatAt+"%d"+FormatEnd, uin)
}
