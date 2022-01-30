/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/27 19:26
  @Notice:  一些结构化的内容
*/

package Msg

const (
	FormatPic   = "[Pic="
	FormatAt    = "[At="
	FormatReply = "[Reply,FromUin=%s,MsgSeq=%s,SendTime=%s"
	FormatEnd   = "]"
)

const (
	TypeText  = 0x01
	CommonMsg = 0x01
	CommonAt  = 0x06
	TypeFace  = 0x02
	TypeReply = 0x19
	TypePic   = 0x03
)
