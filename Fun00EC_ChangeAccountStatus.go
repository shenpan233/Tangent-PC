/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 14:39
  @Notice:	修改在线状态
*/

package Tangent_PC

import (
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

// Online 0A=在线，3C=Q我吧，1E=离开，32=忙碌，46=请勿打扰，28=隐身
const (
	Online    = 0x0A //在线
	QMe       = 0x3C //Q我吧
	Leave     = 0x1E //离开
	Busy      = 0x32 //忙碌
	DND       = 0x46 //勿扰
	Invisible = 0x28 //隐身
)

func (this *TangentPC) pack00EC(Status uint16) (SsoSeq uint16, buffer []byte) {
	return this.packetCommonEnc(0x00_EC, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint8(01)
		pack.SetUint16(Status)
		pack.SetBytes([]byte{0, 1, 0, 1, 0, 4, 0, 0, 0, 0})
	}))
}
