/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		令牌
* @Creat:   2021/12/4 0004 15:28
 */
package PCQQ

import (
	"Tangent-PC/protocal/Tlv"
	util "Tangent-PC/utils"
	"Tangent-PC/utils/GuBuffer"
)

func (this *TangentPC) pack0818() (SsoSeq uint16, buffer []byte) {
	this.teaKey.Ping0818Key = util.GetRandomBin(16)
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

func (this *TangentPC) unpack0818(bin []byte) (ret *QRResp) {
	pack := GuBuffer.NewGuUnPacket(util.Decrypt(this.teaKey.ShareKey, bin[3:]))
	ret = new(QRResp)
	ret.Status = pack.GetUint8()
	/*Tlv解析*/
	for pack.GetLen() > 0 {
		if tlv := pack.GetTlv(); tlv != nil {
			GuBuffer.NewGuUnPacketFun(tlv.Value, func(tPack *GuBuffer.GuUnPacket) {
				switch tlv.Tag {
				case 0x302:
					ret.QRCode = tPack.GetToken()
					break
				case 0x30:
					ret.sig0x30 = tPack.GetAll()
					break
				case 0x301:
					ret.sigQRsing = tPack.GetAll()
					break
				case 0x9:
					tPack.GetBin(2)
					ret.key = tPack.GetAll()
					break
				}
			})
		}
	}

	return
}
