/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:	    登录相关的Api
* @Creat:   2021/12/3 0003 21:53
 */

package Tangent_PC

import (
	"errors"
	"github.com/shenpan233/Tangent-PC/model"
	"github.com/shenpan233/Tangent-PC/utils/Bytes"
	"strconv"
)

//ResetAccount 账号重新设置
func (this *TangentPC) ResetAccount(account string) {
	this.info.LongUin, _ = strconv.ParseUint(account, 10, 64)
	this.info.Account = account
}

// PingServer 连接初始化,Ping服务器
func (this *TangentPC) PingServer() bool {
	ssoSeq, buffer := this.pack0825()
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin == nil {
		//GuLog.Error("返回空字节")
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
func (this *TangentPC) FetchQRCode() *model.QRResp {
	ssoSeq, buffer := this.pack0818()
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		/*正常接收*/
		return this.unpack0818(bin)
	}
	return nil
}

// CheckQRCode 检测二维码状态
func (this TangentPC) CheckQRCode(resp *model.QRResp) uint8 {
	ssoSeq, buffer := this.pack0819(resp)
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		return this.unpack0819(resp, bin)
	}
	return QRUnKnow
}

//QRLogin 二维码登录
func (this *TangentPC) QRLogin() (err error, tgt *model.TgtInfo) {
	ssoSeq, buffer := this.pack0836QrCode()
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		tgt = this.unpack0836QrCode(bin)
		if tgt != nil {
			ssoSeq, buffer := this.pack0828(tgt)
			if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
				result := uint8(0)
				if result, err = this.unpack0828(bin, tgt); result == 0 {
					this.finishLogin()
					return
				} else {
					//GuLog.Error("QRLogin", "%s", err.Error())
				}
			}
		}
	}
	return
}

//LoginByToken	令牌登录
func (this *TangentPC) LoginByToken(tgt *model.TgtInfo) (code uint8, err error) {
	if tgt != nil {
		ssoSeq, buffer := this.pack0828(tgt)
		if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
			code, err = this.unpack0828(bin, tgt)
			if code == 0 {
				this.finishLogin()
				return
			} else {
				return
			}
		} else {
			err = errors.New("TGTGT失效")
		}
	}
	return
}

//Login 账号密码登录
func (this *TangentPC) Login(Password string) {
	this.info.PassWord = Bytes.GetMd5Bytes([]byte(Password))
	ssoSeq, buffer := this.pack0836Common()
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		this.unpack0836Login(bin)
	}
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
		this.unpack0058(bin[3:])
		return true
	}
	return false
}

//Hook 事件回调
func (this *TangentPC) Hook(GroupMsg func(Msg model.GroupMsg)) {
	this.hook.GroupMsg = GroupMsg
}
