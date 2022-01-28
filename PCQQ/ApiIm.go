/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/24 12:06
  @Notice:  聊天相关的Api
*/

package PCQQ

import (
	GroupMsg "Tangent-PC/protocal/Msg/Group/Receive"
	GroupSend "Tangent-PC/protocal/Msg/Group/Send"
	"Tangent-PC/protocal/Protobuf/im/cmd0x0002"
)

//RevokeGroupMessage 撤回消息	(:要有管理员权限
func (this *TangentPC) RevokeGroupMessage(GroupCode uint64, MsgSeq uint32) {
	ssoSeq, buffer := this.pack0x3f7(GroupCode, MsgSeq)
	this.udper.SendAndGet(ssoSeq, WaitTime, &buffer)
}

//ReadGroupMsg 置群消息已读
//	内部会自动调用不用管
func (this *TangentPC) ReadGroupMsg(GroupCode uint64, MsgSeq uint32) bool {
	ssoSeq, buffer := this.pack0002(GroupMsg.ReadMsg(GroupCode, MsgSeq))
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		isSuc, _ := this.unpack0002(bin)
		return isSuc
	}
	return false
}

//SendGroupMsg 发送群消息
//	GroupCode 群号
//	Msg 	  消息内容
func (this *TangentPC) SendGroupMsg(GroupCode uint64, Msg string) (Code bool, MsgSeq uint32) {
	ssoSeq, buffer := this.pack0002(GroupSend.GroupMsg(GroupCode, Msg))
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		isSuc, Recall := this.unpack0002(bin)
		msg := Recall.(cmd0x0002.SendGroupMsg)
		return isSuc, msg.GetMsgSeq()
	}
	return false, 0
}
