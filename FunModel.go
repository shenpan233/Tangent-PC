/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 17:46
  @Notice:  用于抄袭的组包解包模板
*/

package Tangent_PC

import (
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

type (
	unpack func(_ uint16, bin []byte)
)

func (this *TangentPC) pack() (SsoSeq uint16, buffer []byte) {
	return this.packetCommonEnc(0x00, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {

	}))
}

func (this *TangentPC) unpack(bin []byte) {
	GuBuffer.NewGuUnPacketFun(util.Decrypt(this.teaKey.SessionKey, bin), func(pack *GuBuffer.GuUnPacket) {

	})
}
