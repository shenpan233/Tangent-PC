/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/2/25 21:44
  @Notice:  SigHttpConn
*/

package Tangent_PC

import "github.com/shenpan233/Tangent-PC/utils/GuBuffer"

const (
	subCmd0x01BBHttpConn = 3
)

func (this *TangentPC) pack01BB(subCmd uint8) (SsoSeq uint16, buffer []byte) {
	return this.packetCommonEnc(0x01_BB, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint8(subCmd)
		pack.SetUint32(this.sdk.ServiceId)
	}))
}
