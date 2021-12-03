/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		应用接口
* @Creat:   2021/12/3 0003 21:53
 */
package PCQQ

import (
	"Tangent-PC/utils/GuLog"
)

/*连接初始化,Ping服务器*/
func (this *TangentPC) Ping() {
	ssoSeq, buffer := this.pack0825()
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin == nil {
		/*无接收返回*/
		GuLog.Error("[%s]=>Ping失败", this.info.Account)
		return
	} else {
		/*正常接收*/
		switch this.unpack0825(bin[3:]) {
		case _0825redirect:

		}
	}
}
