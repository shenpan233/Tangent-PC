/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/24 12:06
  @Notice:  聊天相关的Api
*/

package PCQQ

import (
	GroupMsg "Tangent-PC/protocal/Msg/Group/Receive"
	GroupSend "Tangent-PC/protocal/Msg/Group/Send"
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
		return this.unpack0002(bin)
	}
	return false
}

//SendGroupMsg 发送群消息
//	GroupCode 群号
//	Msg 	  消息内容
func (this *TangentPC) SendGroupMsg(GroupCode uint64, Msg string) bool {
	ssoSeq, buffer := this.pack0002(GroupSend.GroupMsg(GroupCode, Msg))
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		return this.unpack0002(bin)
	}
	return false
}
