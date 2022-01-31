/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 17:46
  @Notice:	刷新token
*/

package Tangent_PC

import "github.com/shenpan233/Tangent-PC/utils/GuBuffer"

func (this *TangentPC) pack001D(t uint8) (SsoSeq uint16, buffer []byte) {
	return this.packetCommonEnc(0x00_1D, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint8(t)
	}))
}
