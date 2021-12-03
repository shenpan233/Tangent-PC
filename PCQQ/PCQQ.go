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
	"strconv"
)

var WaitTime = uint32(3000) /*三秒延迟*/

type TangentPC struct {
	sdk    *model.Version
	info   *model.Information
	udper  *udper.Udper
	teaKey *model.TeaKey
}

func New(Account string) (this *TangentPC) {
	this = new(TangentPC)
	/*通讯器部分*/
	{
		if this.udper = udper.New(model.TxServer[1], &udper.Set{
			BuffMaxSize: 1024,
			UdpRecv:     nil,
		}); this.udper == nil {
			GuLog.Error("New", "服务器连接失败")
			return
		}
	}

	/*硬件信息等部分*/
	{
		this.sdk = new(model.Version)
		this.info = &model.Information{
			LongUin: func() uint64 {
				uint, _ := strconv.ParseUint(Account, 10, 64)
				return uint
			}(),
			Account:  Account,
			PassWord: nil,
		}
	}

	/*Tea秘钥申请*/
	{
		this.teaKey = new(model.TeaKey)
		this.teaKey.PublicKey = util.HexToBin("03 08 6D E3 9B B7 88 DF E1 4A 33 4D E3 65 4D 6B CE 2A 6D A9 DA AA 52 5F 02")
	}

	return
}
