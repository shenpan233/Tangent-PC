/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/31 20:36
  @Notice:  不知道什么离谱的
*/

package Tangent_PC

import (
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

const (
	Remark = 0x0D //备注信息
)

func (this *TangentPC) pack003E(subCmd uint8, uin uint32) (SsoSeq uint16, buffer []byte) {
	return this.packetCommonEnc(0x00_3E, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint8(subCmd)
		pack.SetUint32(uin)
		pack.SetBytes([]byte{0x00})
	}))
}

func (this *TangentPC) unpack003E(bin []byte) (ret interface{}) {
	GuBuffer.NewGuUnPacketFun(util.Decrypt(this.teaKey.SessionKey, bin), func(pack *GuBuffer.GuUnPacket) {
		subCmd := pack.GetUint8()
		_ = pack.GetUint8() == 0x01 //是否查询自己
		pack.Skip(4)                //跳过Uin
		switch subCmd {
		case Remark:
			ret = string(pack.GetToken())
			return
		}
	})
	return nil
}
