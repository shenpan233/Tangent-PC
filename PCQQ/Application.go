/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		应用接口(对外)
* @Creat:   2021/12/3 0003 21:53
 */

package PCQQ

import (
	util "Tangent-PC/utils"
	"Tangent-PC/utils/GuLog"
)

// PingServer 连接初始化,Ping服务器
func (this *TangentPC) PingServer() bool {
	ssoSeq, buffer := this.pack0825()
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin == nil {
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

// FetchQRCode 获取登录二维码
func (this *TangentPC) FetchQRCode() *QRResp {
	ssoSeq, buffer := this.pack0818()
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		/*正常接收*/
		return this.unpack0818(bin)
	}
	return nil
}

// CheckQRCode 检测二维码状态
func (this TangentPC) CheckQRCode(resp *QRResp) uint8 {
	ssoSeq, buffer := this.pack0819(resp)
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		return this.unpack0819(resp, bin)
	}
	return QRUnKnow
}

func (this *TangentPC) QRLogin() bool {
	ssoSeq, buffer := this.pack0836QrCode()
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		tgt := this.unpack0836(bin)
		if tgt != nil {
			ssoSeq, buffer := this.pack0828(tgt)
			if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
				GuLog.Warm("QRLogin", "%s", util.BinToHex(bin[3:]))
				if this.unpack0828(bin, tgt) == 0 {
					this.finishLogin()
					return true
				}
			}
		}
	}
	return false
}

//ChangeOnlineStatus 修改在线状态
func (this *TangentPC) ChangeOnlineStatus(OnLineSts uint16) bool {
	ssoSeq, buffer := this.pack00EC(OnLineSts)
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		return true
	}
	return false
}

// HeatBoat 心跳
func (this *TangentPC) HeatBoat() bool {
	ssoSeq, buffer := this.pack0058HeatBoat()
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		return true
	}
	return false
}
