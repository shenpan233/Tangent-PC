/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/2 19:24
  @Notice:	接收系统消息
*/

package Tangent_PC

import (
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

func (this *TangentPC) unpack0017(Seq uint16, bin []byte) {
	pack := GuBuffer.NewGuUnPacket(bin)
	MsgInfo := pack.GetBin(16) //可以细拆，但是没必要
	pack.GetUint16()           //服务器端口 8000
	cmd := pack.GetUint16()    //接收命令
	this.GetServerMsg(0x00_17, cmd, Seq, MsgInfo, pack.GetAll())
	return
}
