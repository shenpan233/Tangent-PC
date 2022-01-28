/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 17:48
  @Notice:  内部业务
*/

package PCQQ

import (
	"fmt"
	GroupMsg "github.com/shenpan233/Tangent-PC/protocal/Msg/Group/Receive"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
	"github.com/shenpan233/Tangent-PC/utils/GuLog"
)

//refreshClient 刷新ClientKey
func (this *TangentPC) refreshClient() {
	ssoSeq, buffer := this.pack001D(0x11)
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		GuLog.Error("refreshClient", util.BinToHex(bin))
	}
}

//refresh26 刷新Token26
func (this *TangentPC) refresh26() {
	ssoSeq, buffer := this.pack001D(0x26)
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		GuLog.Error("refresh26", util.BinToHex(bin))
	}
}

//finishLogin 登录完成后的操作
func (this *TangentPC) finishLogin() {
	//刷新token
	this.refreshClient()
	this.refresh26()
	//绑定接收器
	this.udper.UdpRecv = this.Recv
	this.handle = map[int16]unpack{
		0x00_17: this.unpack0017,
	}
}

//GetServerMsg 读取系统信息
func (this *TangentPC) GetServerMsg(Cmd int16, seq uint16, MsgInfo, data []byte) {
	go func() {
		buffer := this.pack0017(seq, MsgInfo)
		this.udper.Send(&buffer)
	}()
	switch Cmd {
	case 0x00_52:
		Msg := GroupMsg.GroupMsg(data)
		Msg.Account = this.info.LongUin
		if this.hook.GroupMsg != nil {
			fmt.Println(Msg)
			go this.ReadGroupMsg(Msg.GroupUin, Msg.MsgSeq)
			go this.hook.GroupMsg(Msg)
		}
		break

	}
}

//Recv 数据包接收
func (this *TangentPC) Recv(Cmd int16, seq uint16, pack *GuBuffer.GuUnPacket) {
	pack.GetBin(3)
	pack = GuBuffer.NewGuUnPacket(util.Decrypt(this.teaKey.SessionKey, pack.GetAll()))
	if event := this.handle[Cmd]; event != nil {
		event(seq, pack.GetAll())
	} else {
		GuLog.Info("ReCv", "QQ:[%d],Cmd=0x%X,Buff=%X", this.info.LongUin, Cmd, pack.GetAll())
	}

}
