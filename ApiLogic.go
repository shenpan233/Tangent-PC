/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 17:48
  @Notice:  内部业务
*/

package Tangent_PC

import (
	"github.com/shenpan233/Tangent-PC/model"
	GroupMsg "github.com/shenpan233/Tangent-PC/protocal/Msg/Group/Receive"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
	"github.com/shenpan233/Tangent-PC/utils/GuLog"
)

//refreshClient 刷新ClientKey
func (this *TangentPC) refreshClient() {
	ssoSeq, buffer := this.pack001D(0x11)
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		//GuLog.Error("refreshClient", util.BinToHex(bin))
	}
}

//refreshHttpConnSig 刷新HttpConnSig
func (this *TangentPC) refreshHttpConnSig() bool {
	ssoSeq, buffer := this.pack001D(subCmd0x001DHttpConn)
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		if gen := this.unpack001D(bin); gen != nil {
			gen := gen.(map[int8][]byte)
			this.teaKey.HttpConn = gen[0]
			this.sig.BufSigHttpConnToken = gen[1]
			return true
		} else {
			return false
		}
	}
	return false
}

//refreshWebKey 刷新WebKey
func (this *TangentPC) refreshWebKey() {
	ssoSeq, buffer := this.pack001D(subCmd0x001DWebKey)
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		ret := this.unpack001D(bin)
		if ret != nil {
			this.info.SelfWebKey = ret.(*model.WebKey)
		}
	}

}

//GenHttpConn
func (this *TangentPC) genHttpConn() {
	ssoSeq, buffer := this.pack01BB(subCmd0x01BBHttpConn)
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
	}
}

//finishLogin 登录完成后的操作
func (this *TangentPC) finishLogin() {
	//刷新token
	this.refreshClient()
	this.refreshHttpConnSig()
	this.refreshWebKey()
	//绑定接收器
	this.udper.UdpRecv = this.Recv
	this.handle = map[uint16]unpack{
		0x00_17: this.unpack0017,
	}
}

//GetServerMsg 读取系统信息
func (this *TangentPC) GetServerMsg(Cmd uint16, seq uint16, MsgInfo, data []byte) {
	go func() {
		buffer := this.pack0017(seq, MsgInfo)
		this.udper.Send(&buffer)
	}()
	switch Cmd {
	case 0x00_52:
		Msg := GroupMsg.GroupMsg(data)
		Msg.Account = this.info.LongUin
		if this.hook.GroupMsg != nil {
			go this.ReadGroupMsg(Msg.GroupUin, Msg.MsgSeq)
			go this.hook.GroupMsg(Msg)
		}
		break

	}
}

//Recv 数据包接收
func (this *TangentPC) Recv(Cmd uint16, seq uint16, pack *GuBuffer.GuUnPacket) {
	pack.GetBin(3)
	pack = GuBuffer.NewGuUnPacket(util.Decrypt(this.teaKey.SessionKey, pack.GetAll()))
	if event := this.handle[Cmd]; event != nil {
		event(seq, pack.GetAll())
	} else {
		GuLog.Info("ReCv", "QQ:[%d],Cmd=0x%X,Buff=%X", this.info.LongUin, Cmd, pack.GetAll())
	}

}
