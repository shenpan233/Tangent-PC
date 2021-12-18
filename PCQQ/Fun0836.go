/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		登录的最后一个封包
* @Creat:   2021/12/10 23:24
 */

package PCQQ

import (
	"Tangent-PC/protocal/Tlv"
	util "Tangent-PC/utils"
	"Tangent-PC/utils/GuBuffer"
)

//组包

func (this *TangentPC) pack0836QrCode() (Ssoseq uint16, data []byte) {
	return this.packetLogin(0x08_36, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint16(1) //01禁用二次ECDH 02启用
		pack.SetBytes(Tlv.GetTlv103(&this.teaKey.PublicKey))
		pack.SetUint16(0)
		pack.SetToken(util.GetRandomBin(16))
		pack.SetBytes(util.Encrypt(this.teaKey.ShareKey, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
			tlv15 := Tlv.GetTlv15(&this.info.Computer)
			pack.SetBytes(Tlv.GetTlv112SigClientAddr(&this.sig.BufSigClientAddr))
			pack.SetBytes(Tlv.GetTlv30FPcName(this.info.ComputerName))
			pack.SetBytes(Tlv.GetTlv5Uin(this.info.LongUin))
			pack.SetBytes(Tlv.GetTl303(this.sig.BufQR303))
			pack.SetBytes(tlv15)
			pack.SetBytes(Tlv.GetTlv1A(this.sig.BufTgTGTKey, tlv15))
			pack.SetBytes(Tlv.GetTlv18Ping(this.info.LongUin, this.sdk, uint16(this.info.RedirectIp.Len())))
			pack.SetBytes(Tlv.GetTlv312MiscFlag())
			pack.SetBytes(Tlv.GetTlv508())
			pack.SetBytes(Tlv.GetTlv313GUIDEx(this.info.Computer.MacGuid))
			pack.SetBytes(Tlv.GetTlv102Official(this.info))
		})))
	}))
}
