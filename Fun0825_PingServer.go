/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:	    0825组包/解包 服务器Ping
* @Creat:   2021/11/26 0026 22:48
 */

package Tangent_PC

import (
	"github.com/shenpan233/Tangent-PC/model"
	"github.com/shenpan233/Tangent-PC/protocal/Tlv"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

const (
	_0825Redirect = 0xFE
	_0825PingSuc  = model.LogicSuc
)

//pack0825 0825组包
func (this *TangentPC) pack0825() (SsoSeq uint16, buffer []byte) {
	this.teaKey.Ping0825Key = util.GetRandomBin(16)
	return this.packetLogin(0x08_25, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetBytes(this.teaKey.Ping0825Key)
		pack.SetBytes(util.Encrypt(this.teaKey.Ping0825Key, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
			pack.SetBytes(Tlv.GetTlv18Ping(this.info.LongUin, this.sdk, uint16(this.info.RedirectIp.Len())))
			pack.SetBytes(Tlv.GetTlv4())
			pack.SetBytes(Tlv.GetTlv309PingStrategy("0.0.0.0", this.info.RedirectIp))
			pack.SetBytes(Tlv.GetTlv114DHParams(&this.teaKey.PublicKey))
			pack.SetBytes(Tlv.GetTlv511())
		})))
	}))
}

//	返回值说明 0xFE->需要重定向
func (this *TangentPC) unpack0825(bin []byte) (result uint8) {
	pack := GuBuffer.NewGuUnPacket(util.Decrypt(this.teaKey.Ping0825Key, bin))
	result = pack.GetUint8()
	/*Tlv解析*/
	this.uCode0825Tlv(pack)
	return
}

func (this TangentPC) uCode0825Tlv(pack *GuBuffer.GuUnPacket) {
	for pack.GetLen() > 0 {
		if tlv := pack.GetTlv(); tlv != nil {
			GuBuffer.NewGuUnPacketFun(tlv.Value, func(tPack *GuBuffer.GuUnPacket) {
				switch tlv.Tag {
				case 0x00_0C: /*Tlv000C*/
					tPack.GetBin(12)
					this.info.ConnectIp = util.IntToIp(int32(tPack.GetUint32()))
					this.info.RedirectIp.PushBack(this.info.ConnectIp)
				case 0x01_12:
					this.sig.BufSigClientAddr = tlv.Value
				case 0x00_17:
					tPack.GetUint16()
					this.info.PingTime = tPack.GetUint32()
					this.info.WlanIp = util.IntToIp(int32(tPack.GetUint32()))
				}
			})
		}
	}
}
