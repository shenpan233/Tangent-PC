/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/24 12:06
  @Notice:  聊天相关的Api
*/

package PCQQ

import "Tangent-PC/protocal/Msg/GroupMsg"

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
