/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		应用接口
* @Creat:   2021/12/3 0003 21:53
 */
package PCQQ

/*连接初始化,Ping服务器*/
func (this *TangentPC) PingServer() bool {
	ssoSeq, buffer := this.pack0825()
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin == nil {
		/*无接收返回*/
		return false
	} else {
		/*正常接收*/
		switch this.unpack0825(bin[3:]) {
		case _0825Redirect: /*需要重定向*/
			/*重建udp组件*/
			if this.udper.ChangeConnect(this.info.ConnectIp + ":8000") {
				if this.PingServer() {
					return true
				}
			}
			break
		case _0825PingSuc:
			return true
		}
	}
	return false
}

/*获取登录二维码*/
func (this *TangentPC) FetchQRCode() *QRResp {
	ssoSeq, buffer := this.pack0818()
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		/*正常接收*/
		return this.unpack0818(bin)
	}
	return nil
}

/*检测二维码状态*/
func (this TangentPC) CheckQRCode(resp *QRResp) {
	ssoSeq, buffer := this.pack0819(resp)
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		this.unpack0819(resp, bin)
	}
	return
}
