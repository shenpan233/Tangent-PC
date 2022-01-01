/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2021/12/31 23:55
  @Notice: 0828组包和解包
*/

package PCQQ

import (
	"Tangent-PC/protocal/Tlv"
	"Tangent-PC/utils/GuBuffer"
)

func (this *TangentPC) pack0828(tgt *tgtInfo) (SsoSeq uint16, buffer []byte) {
	return this.packetLogin(0x08_28, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetBytes(Tlv.GetTlv7Tgt(&tgt.bufTgt))
		pack.SetBytes(Tlv.GetTlvC(this.info.Computer.ConnectIp))
		pack.SetBytes(Tlv.GetTlv15(&this.info.Computer))
		pack.SetBytes(Tlv.GetTlv36LoginReason())
		pack.SetBytes(Tlv.GetTlv18Ping(this.info.LongUin, this.sdk, uint16(this.info.RedirectIp.Len())))
		pack.SetBytes(Tlv.GetTlv1FDeviceID(this.info.Computer.DeviceID))
		pack.SetBytes(Tlv.GetTlv103())
	}))
}

//
func (this *TangentPC) unpack0828(bin []byte) (result uint8) {
	//pack := GuBuffer.NewGuUnPacket(util.Decrypt(this.teaKey.Ping0825Key, bin))

	return
}
