/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		协议主类
* @Creat:   2021/11/26 0026 22:45
 */
package PCQQ

import (
	"Tangent-PC/model"
	util "Tangent-PC/utils"
	"Tangent-PC/utils/GuLog"
	"Tangent-PC/utils/udper"
	"container/list"
	"strconv"
)

var WaitTime = uint32(3) /*三秒延迟*/

type TangentPC struct {
	sdk    *model.Version
	info   *model.Information
	sig    *model.Sig
	udper  *udper.Udper
	teaKey *model.TeaKey
}

func New(Account string, Computer model.Computer) (this *TangentPC) {
	this = new(TangentPC)
	/*通讯器部分*/
	{
		if this.udper = udper.New(model.TxServer[1]+":8000", &udper.Set{
			BuffMaxSize: 1024,
			UdpRecv:     nil,
		}); this.udper == nil {
			GuLog.Error("New", "服务器连接失败")
			return nil
		}
	}

	/*硬件信息等部分*/
	{
		this.sdk = new(model.Version)
		Computer.RedirectIp = list.New()
		this.info = &model.Information{
			LongUin: func() uint64 {
				Uint, _ := strconv.ParseUint(Account, 10, 64)
				return Uint
			}(),
			Account:  Account,
			PassWord: nil,
			Computer: Computer,
		}
		this.info.ComputerId = append(this.info.ComputerId[:4], 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)

	}

	/*Tea秘钥申请*/
	{
		this.sig = new(model.Sig)
		this.teaKey = new(model.TeaKey)
		this.teaKey.PublicKey = util.HexToBin("03 1F 06 FA 3B 19 BF F9 2C 7C 02 7D 5D EA C5 60 83 52 86 C1 BF 75 CA 2A 96")
		this.teaKey.ShareKey = util.HexToBin("6C 3E 9F 64 1C 27 F9 CA D6 B6 37 8A A7 74 D0 04")
	}

	return
}
