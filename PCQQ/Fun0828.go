/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2021/12/31 23:55
  @Notice:	0828组包和解包
*/

package PCQQ

import (
	"Tangent-PC/protocal/Tlv"
	"Tangent-PC/utils/GuBuffer"
)

func (this *TangentPC) pack0828(tgt *tgtInfo) (SsoSeq uint16, buffer []byte) {
	return this.packetLogin(0x08_28, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetBytes(Tlv.GetTlv7Tgt(&tgt.bufTgt))
		pack.SetBytes()

	}))
}

//
func (this *TangentPC) unpack0828(bin []byte) (result uint8) {
	//pack := GuBuffer.NewGuUnPacket(util.Decrypt(this.teaKey.Ping0825Key, bin))

	return
}
