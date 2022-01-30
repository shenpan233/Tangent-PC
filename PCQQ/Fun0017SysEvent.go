/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/2 19:24
  @Notice:	接收系统消息
*/

package PCQQ

import (
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

func (this *TangentPC) unpack0017(Seq uint16, bin []byte) {
	pack := GuBuffer.NewGuUnPacket(bin)
	MsgInfo := pack.GetBin(16) //可以细拆，但是没必要
	pack.GetInt16()            //服务器端口 8000
	cmd := pack.GetInt16()     //接收命令
	this.GetServerMsg(cmd, Seq, MsgInfo, pack.GetAll())

	return
}
func (this *TangentPC) pack0017(Seq uint16, MsgInfo []byte) (buffer []byte) {
	return GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint8(2)
		pack.SetUint16(this.sdk.CMainVer)
		pack.SetUint16(0x00_17)
		pack.SetUint16(Seq)
		pack.SetUint32(uint32(this.info.LongUin))
		pack.SetBytes([]byte{0x02, 0x00, 0x00})
		pack.SetUint32(0x00_01_01_01)
		pack.SetUint32(this.sdk.DwPubNo)
		pack.SetBytes(util.Encrypt(this.teaKey.SessionKey, MsgInfo))
		pack.SetUint8(3)
	})
}
