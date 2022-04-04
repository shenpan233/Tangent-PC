/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:	    令牌
* @Creat:   2021/12/4 0004 15:28
 */

package Tangent_PC

import (
	"github.com/shenpan233/Tangent-PC/model"
	"github.com/shenpan233/Tangent-PC/protocal/Tlv"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
	"github.com/shenpan233/Tangent-PC/utils/GuLog"
)

func (this *TangentPC) pack0818() (SsoSeq uint16, buffer []byte) {
	this.teaKey.Ping0818Key = util.GetRandomBin(16)
	GuLog.DebugF("BufTgtKey:%X\nPublicKey:%X\nShareKey:%X", this.teaKey.Ping0818Key, this.teaKey.PublicKey, this.teaKey.ShareKey)

	return this.packetLogin(0x08_18, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetBytes(this.teaKey.Ping0818Key)
		pack.SetBytes(util.Encrypt(this.teaKey.Ping0818Key, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
			pack.SetBytes(Tlv.GetTlv19SSOInfo(this.sdk))
			pack.SetBytes(Tlv.GetTlv114DHParams(&this.teaKey.PublicKey))
			pack.SetBytes(Tlv.GetTlv305())
			pack.SetBytes(Tlv.GetTlv15(&this.info.Computer))
		})))
	}))
}

func (this *TangentPC) unpack0818(bin []byte) (ret *model.QRResp) {
	pack := GuBuffer.NewGuUnPacket(util.Decrypt(this.teaKey.ShareKey, bin[3:]))
	ret = new(model.QRResp)
	ret.Status = pack.GetUint8()
	for pack.GetLen() > 0 {
		if tlv := pack.GetTlv(); tlv != nil {
			GuBuffer.NewGuUnPacketFun(tlv.Value, func(pack *GuBuffer.GuUnPacket) {
				switch tlv.Tag {
				case 0x302:
					ret.QRCode = pack.GetToken()
					break
				case 0x30:
					ret.Sig0x30 = pack.GetAll()
					break
				case 0x301:
					ret.SigQRSing = pack.GetAll()
					break
				case 0x9:
					pack.Skip(2)
					ret.BufQRKey = pack.GetAll()
					break
				}
			})
		}
	}
	//if ret.Status == model.LogicSuc {
	//	/*Tlv解析*/
	//
	//} else {
	//
	//}

	return
}
