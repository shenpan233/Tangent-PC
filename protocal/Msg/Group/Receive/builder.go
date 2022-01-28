/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/26 12:27
  @Notice:  消息结构构造器
*/

package Receive

import (
	"github.com/shenpan233/Tangent-PC/protocal/Msg"
)

func buildPic(guid string) string {
	return Msg.FormatPic + guid + Msg.FormatEnd
}
