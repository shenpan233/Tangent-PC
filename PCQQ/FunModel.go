/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 17:46
  @Notice:
*/

package PCQQ

import "Tangent-PC/utils/GuBuffer"

type (
	unpack func(_ uint16, bin []byte)
)

func (this *TangentPC) pack() (SsoSeq uint16, buffer []byte) {
	return this.packetCommonEnc(0x00, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {

	}))
}

func (this *TangentPC) unpack(bin []byte) {
	GuBuffer.NewGuUnPacketFun(bin, func(pack *GuBuffer.GuUnPacket) {

	})
}
