/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/26 13:31
  @Notice:	群消息处理/发送
*/

package PCQQ

import (
	GroupMsgModel "Tangent-PC/protocal/Msg/Group"
	GroupMsg "Tangent-PC/protocal/Msg/Group/Receive"
	util "Tangent-PC/utils"
	"Tangent-PC/utils/GuBuffer"
)

func (this *TangentPC) pack0002(bin []byte) (SsoSeq uint16, buffer []byte) {
	return this.packetCommonEnc(0x00_02, bin)
}

func (this *TangentPC) unpack0002(bin []byte) (isSuc bool) {
	GuBuffer.NewGuUnPacketFun(util.Decrypt(this.teaKey.SessionKey, bin[3:]), func(pack *GuBuffer.GuUnPacket) {
		switch pack.GetUint8() { //事件类型
		case GroupMsgModel.Receipt:
			isSuc = GroupMsg.UnReadMsg(pack.GetAll())
			return
		}
	})
	return false
}
