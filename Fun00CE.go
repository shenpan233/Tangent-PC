/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/3 22:28
  @Notice:  00CE接收好友消息
*/

package Tangent_PC

import "github.com/shenpan233/Tangent-PC/utils/GuBuffer"

func (this *TangentPC) unpack00CE(Seq uint16, bin []byte) {
	pack := GuBuffer.NewGuUnPacket(bin)
	MsgInfo := pack.GetBin(16)
	pack.Skip(2)
	cmd := pack.GetUint16()
	this.GetServerMsg(0x00_CE, cmd, Seq, MsgInfo, pack.GetAll())
	return
}
