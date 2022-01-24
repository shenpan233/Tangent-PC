/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 15:15
  @Notice:心跳
*/

package PCQQ

import "Tangent-PC/utils/GuBuffer"

func (this *TangentPC) pack0058HeatBoat() (SsoSeq uint16, buffer []byte) {
	return this.packetCommonEnc(0x00_58, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetBytes([]byte{0x0, 0x1, 0x0, 0x1})
	}))
}

func (this *TangentPC) unpack0058(bin []byte) (status uint8) {
	return
}
