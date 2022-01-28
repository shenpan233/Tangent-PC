/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/23 22:01
  @Notice:	群消息撤回
*/

package PCQQ

import (
	"fmt"
	"github.com/shenpan233/Tangent-PC/protocal/Protobuf/im/cs/cmd0x3f7"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

func (this *TangentPC) pack0x3f7(GroupCode uint64, MsgSeq uint32) (SsoSeq uint16, buffer []byte) {
	return this.packetIMEnc(0x03_F7, cmd0x3f7.GetBuffer(GroupCode, MsgSeq))
}

func (this *TangentPC) unpack0x3f7(bin []byte) {
	GuBuffer.NewGuUnPacketFun(bin, func(pack *GuBuffer.GuUnPacket) {
		len1 := int(pack.GetUint32())
		len2 := int(pack.GetUint32())
		pack.GetBin(len1) //暂时无用
		resp := cmd0x3f7.Decode(pack.GetBin(len2))
		fmt.Println(resp)
	})
}
