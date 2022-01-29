/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/23 22:01
  @Notice:	群消息撤回
*/

package PCQQ

import (
	"errors"
	"github.com/shenpan233/Tangent-PC/model"
	"github.com/shenpan233/Tangent-PC/protocal/Protobuf/im/cs/cmd0x3f7"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

func (this *TangentPC) pack0x3f7(GroupCode uint64, MsgSeq, MsgID uint32) (SsoSeq uint16, buffer []byte) {
	return this.packetIMEnc(0x03_F7, cmd0x3f7.GetBuffer(GroupCode, MsgSeq, MsgID))
}

func (this *TangentPC) unpack0x3f7(bin []byte) (err error) {
	GuBuffer.NewGuUnPacketFun(util.Decrypt(this.teaKey.SessionKey, bin[3:]), func(pack *GuBuffer.GuUnPacket) {
		len1 := int(pack.GetUint32())
		len2 := int(pack.GetUint32())
		pack.GetBin(len1) //暂时无用
		resp := cmd0x3f7.Decode(pack.GetBin(len2))
		switch resp.GetCode() {
		case model.LogicSuc:
			break
		case 1001:
			err = errors.New("you cannot do so,because the message have been revoked")
			break
		default:
			err = errors.New(resp.GetMessage())
		}
	})
	return
}
