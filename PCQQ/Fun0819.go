/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:	    二维码检验
* @Creat:   2021/12/10 21:35
 */

package PCQQ

import (
	"Tangent-PC/protocal/Tlv"
	util "Tangent-PC/utils"
	"Tangent-PC/utils/GuBuffer"
	"Tangent-PC/utils/GuLog"
	"strconv"
)

//组包
func (this *TangentPC) pack0819(resp *QRResp) (SsoSeq uint16, buffer []byte) {
	return this.packetLogin(0x08_19, GuBuffer.NewGuPacketFun(func(packet *GuBuffer.GuPacket) {
		packet.SetBytes(Tlv.GetTlv30(resp.sig0x30))
		packet.SetBytes(util.Encrypt(resp.key, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
			pack.SetBytes(Tlv.GetTlv19SSOInfo(this.sdk))
			pack.SetBytes(Tlv.GetTlv301(resp.sigQRSing))
		}),
		))
	}))
}

//解包

//二维码状态标识
const (
	QRNoAgree = 0x1  /*已扫码但未点击确认*/
	QRNoScan  = 0x2  /*未扫码*/
	QROk      = 0x0  /*已确认登录*/
	QRUnKnow  = 0xFF /*未响应*/
)

func (this *TangentPC) unpack0819(qrResp *QRResp, bin []byte) (status uint8) {
	pack := GuBuffer.NewGuUnPacket(util.Decrypt(qrResp.key, bin[3:]))
	status = pack.GetUint8()
	qrResp.Status = status
	switch status {
	case QRNoScan, QRNoAgree:
		return
	case QROk:
		/*已确认登录*/
		/*Tlv解析*/
		for pack.GetLen() > 0 {
			if tlv := pack.GetTlv(); tlv != nil {
				GuBuffer.NewGuUnPacketFun(tlv.Value, func(tPack *GuBuffer.GuUnPacket) {
					switch tlv.Tag {
					case 0x00_04: //扫码QQ号
						tPack.GetInt16()
						Account := tPack.GetStr(int32(tPack.GetInt16()))
						//号码初始化
						{
							GuLog.Notice("扫码成功", "Uin=%s", Account)
							this.info.Account = Account
							this.info.LongUin, _ = strconv.ParseUint(Account, 10, 64)
						}
						break
					case 0x03_03: //一种临时密码
						this.sig.BufQR303 = tlv.Value
						this.info.PassWord = util.ToMd5Bytes(tPack.GetAll())
						break
					case 0x03_04:
						this.sig.BufTgTGTKey = tPack.GetAll()
					}

				})
			}
		}
		break
	}

	return
}
