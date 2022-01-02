/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 17:48
  @Notice:	内部业务接口
*/

package PCQQ

import (
	util "Tangent-PC/utils"
	"Tangent-PC/utils/GuLog"
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
