/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/24 12:06
  @Notice:	聊天相关的Api
*/

package PCQQ

func (this *TangentPC) RevokeGroupMessage(GroupCode uint64, MsgSeq uint32) {
	ssoSeq, buffer := this.pack0x3f7(GroupCode, MsgSeq)
	this.udper.SendAndGet(ssoSeq, WaitTime, &buffer)
}
