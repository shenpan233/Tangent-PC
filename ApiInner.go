/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/3 22:05
  @Notice:  内部主动事件
*/

package Tangent_PC

import (
	"bytes"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
	"github.com/shenpan233/Tangent-PC/utils/GuLog"
	"runtime"
	"strconv"
)

//finishLogin 登录完成后的操作
func (this *TangentPC) finishLogin() {
	//刷新token
	this.refreshClient()
	this.refreshHttpConnSig()
	this.refreshWebKey()
	this.getGroupsInfo()
	//绑定接收器
	this.udper.UdpRecv = this.receive
	this.handle = this.eventHandles()
}

//receive 数据包接收
func (this *TangentPC) receive(Cmd uint16, seq uint16, pack *GuBuffer.GuUnPacket) {
	//异常捕获
	defer func() {
		if err := recover(); err != nil {
			lineW := bytes.NewBufferString("")
			for i := 0; i < 2; i++ {
				_, file, line, _ := runtime.Caller(3 + i)
				lineW.WriteString(file + ":" + strconv.Itoa(line))
				lineW.WriteString("\n")
			}
			GuLog.ErrorF("%s\n%v\n", lineW.String(), err)
		}
	}()
	pack.GetBin(3)
	pack = GuBuffer.NewGuUnPacket(util.Decrypt(this.teaKey.SessionKey, pack.GetAll()))
	if event := this.handle[Cmd]; event != nil {
		event(seq, pack.GetAll())
	} else {
		//未定义事件
		GuLog.WarmF("Receive\nQQ:[%d],Cmd=0x%X,Buff=%X", this.info.LongUin, Cmd, pack.GetAll())
	}

}
