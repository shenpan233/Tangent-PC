/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		登录的最后一个封包
* @Creat:   2021/12/10 23:24
 */

package Tangent_PC

import (
	"github.com/shenpan233/Tangent-PC/model"
	"github.com/shenpan233/Tangent-PC/protocal/Tlv"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
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

const (
	ecdhTwice = 0x00_08 //二次ecdh计算的标识
	LoginSuc  = model.LogicSuc
)

//0836解包
func (this *TangentPC) unpack0836(bin []byte) (tgt *model.TgtInfo) {
	pack := GuBuffer.NewGuUnPacket(bin)
	pack.GetUint16()               //是否二次加密
	LoginStatus := pack.GetUint8() //登录状态
	if LoginStatus == LoginSuc {
		tgt = new(model.TgtInfo)
		pack = GuBuffer.NewGuUnPacket(util.Decrypt(this.sig.BufTgTGTKey, util.Decrypt(this.teaKey.ShareKey, pack.GetAll())))
		pack.GetUint8() //不知道什么鬼
		for pack.GetLen() > 0 {
			tlv := pack.GetTlv()
			pack := GuBuffer.NewGuUnPacket(tlv.Value)
			switch tlv.Tag {
			case 0x01_09:
				pack.GetUint16()
				tgt.BufSessionKey = pack.GetBin(16)
				tgt.BufSession = pack.GetToken()
				this.sig.BufPwdForConn = pack.GetToken()
				break
			case 0x01_07:
				pack.GetUint16()
				pack.GetToken()
				tgt.BufTgTgTKey = pack.GetBin(16)
				tgt.BufTgt = pack.GetToken()
				tgt.BufGTKeyST = pack.GetBin(16)
				tgt.BufServiceTicket = pack.GetToken()
				break
			}
		}
	}
	//GuLog.Warm("0836Recv", "%s\nTgt=%X\nbufGTKeyST=%X\nKeySession=%X\nBufSession=%X", util.BinToHex(pack.GetAll()), tgt.bufTgTgTKey, tgt.bufGTKeyST, this.teaKey.SessionKey, this.sig.BufSession)
	return
}
