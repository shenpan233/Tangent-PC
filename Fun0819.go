/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:	    二维码检验
* @Creat:   2021/12/10 21:35
 */

package Tangent_PC

import (
	"github.com/shenpan233/Tangent-PC/model"
	"github.com/shenpan233/Tangent-PC/protocal/Tlv"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/Bytes"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
	"strconv"
)

//组包
func (this *TangentPC) pack0819(resp *QRResp) (SsoSeq uint16, buffer []byte) {
	return this.packetLogin(0x08_19, GuBuffer.NewGuPacketFun(func(packet *GuBuffer.GuPacket) {
		packet.SetBytes(Tlv.GetTlv30(resp.Sig0x30))
		packet.SetBytes(util.Encrypt(resp.BufQRKey, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
			pack.SetBytes(Tlv.GetTlv19SSOInfo(this.sdk))
			pack.SetBytes(Tlv.GetTlv301(resp.SigQRSing))
		}),
		))
	}))
}

//解包

//二维码状态标识
const (
	QRNoAgree = 0x1            /*已扫码但未点击确认*/
	QRNoScan  = 0x2            /*未扫码*/
	QROk      = model.LogicSuc /*已确认登录*/
	QRUnKnow  = 0xFF           /*未响应*/
)

func (this *TangentPC) unpack0819(qrResp *QRResp, bin []byte) (status uint8) {
	pack := GuBuffer.NewGuUnPacket(util.Decrypt(qrResp.BufQRKey, bin[3:]))
	status = pack.GetUint8()
	qrResp.Status = status
	switch status {
	case QRNoScan, QRNoAgree:
		return
	case QROk:
		/*已确认登录*/
		/*Tlv解析*/
		GuBuffer.TlvEnum(pack.GetAll(), map[uint16]func(pack *GuBuffer.GuUnPacket){
			0x00_04: func(pack *GuBuffer.GuUnPacket) {
				pack.GetUint16()
				Account := pack.GetStr(int32(pack.GetUint16()))
				//号码初始化
				{
					//GuLog.Notice("扫码成功", "AtUin=%s", Account)
					this.info.Account = Account
					this.info.LongUin, _ = strconv.ParseUint(Account, 10, 64)
				}
			},
			0x03_03: func(pack *GuBuffer.GuUnPacket) {
				this.sig.BufQR303 = pack.GetAll()
				this.info.PassWord = Bytes.GetMd5Bytes(this.sig.BufQR303)
			},
			0x03_04: func(pack *GuBuffer.GuUnPacket) {
				this.sig.BufTgTGTKey = pack.GetAll()
			},
		})
	}

	return
}
