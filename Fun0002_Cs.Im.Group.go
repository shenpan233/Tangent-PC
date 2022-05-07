/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/26 13:31
  @Notice:	群消息处理/发送
*/

package Tangent_PC

import (
	GroupMsgModel "github.com/shenpan233/Tangent-PC/protocal/Msg/Group"
	GroupMsg "github.com/shenpan233/Tangent-PC/protocal/Msg/Group/Receive"
	GroupSend "github.com/shenpan233/Tangent-PC/protocal/Msg/Group/Send"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

const (
	_0x0002Send = 0x2A
)

func (this *TangentPC) pack0002(bin []byte) (SsoSeq uint16, buffer []byte) {
	return this.packetCommonEnc(0x00_02, bin)
}

func (this *TangentPC) unpack0002(bin []byte) (isSuc bool, Recall interface{}) {
	GuBuffer.NewGuUnPacketFun(util.Decrypt(this.teaKey.SessionKey, bin[3:]), func(pack *GuBuffer.GuUnPacket) {
		switch pack.GetUint8() { //事件类型
		case GroupMsgModel.Receipt:
			isSuc = GroupMsg.UnReadMsg(pack.GetAll())
			return
		case GroupMsgModel.Send:
			isSuc, Recall = GroupSend.Recall(pack.GetAll())
			return
		}
	})
	return false, nil
}
