/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		组包头+尾
* @Creat:   2021/11/27 0027 12:31
 */

package Tangent_PC

import (
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

func (this *TangentPC) packetLogin(cmd uint16, bin []byte) (uint16, []byte) {
	seq := this.udper.GetSeq()
	return seq, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint8(2)
		pack.SetUint16(this.sdk.CMainVer)
		pack.SetUint16(cmd)
		pack.SetUint16(seq)
		pack.SetUint32(uint32(this.info.LongUin))
		pack.SetUint8(03)
		pack.SetUint16(0)
		pack.SetUint32(this.sdk.DwClientType)
		pack.SetUint32(this.sdk.DwPubNo)
		pack.SetUint32(0)
		pack.SetBytes(bin)
		pack.SetUint8(3)
	})
}

//无加密
func (this *TangentPC) packetCommon(cmd uint16, bin []byte) (uint16, []byte) {
	seq := this.udper.GetSeq()
	return seq, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint8(2)
		pack.SetUint16(this.sdk.CMainVer)
		pack.SetUint16(cmd)
		pack.SetUint16(seq)
		pack.SetUint32(uint32(this.info.LongUin))
		pack.SetUint8(02)
		pack.SetUint16(0)
		pack.SetUint32(this.sdk.DwClientType)
		pack.SetUint32(this.sdk.DwPubNo)
		pack.SetBytes(bin)
		pack.SetUint8(3)
	})
}

//自动使用SessionKey加密的packetCommon
func (this *TangentPC) packetCommonEnc(cmd uint16, bin []byte) (uint16, []byte) {
	seq := this.udper.GetSeq()
	return seq, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint8(2)
		pack.SetUint16(this.sdk.CMainVer)
		pack.SetUint16(cmd)
		pack.SetUint16(seq)
		pack.SetUint32(uint32(this.info.LongUin))
		pack.SetBytes([]byte{0x02, 0x00, 0x00})
		pack.SetUint32(this.sdk.DwClientType)
		pack.SetUint32(this.sdk.DwPubNo)
		pack.SetBytes(util.Encrypt(this.teaKey.SessionKey, bin))
		pack.SetUint8(3)
	})
}

//自动使用SessionKey加密的 im功能包
func (this *TangentPC) packetIMEnc(cmd uint16, bin []byte) (uint16, []byte) {
	seq := this.udper.GetSeq()
	return seq, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint8(2)
		pack.SetUint16(this.sdk.CMainVer)
		pack.SetUint16(cmd)
		pack.SetUint16(seq)
		pack.SetUint32(uint32(this.info.LongUin))
		pack.SetBytes([]byte{0x04, 0x00, 0x00})
		pack.SetUint32(this.sdk.DwClientType)
		pack.SetUint32(this.sdk.DwPubNo)
		pack.SetBytes([]byte{0, 0, 0, 0})
		pack.SetBytes([]byte{0, 0, 0, 0})
		pack.SetBytes(util.Encrypt(this.teaKey.SessionKey, bin))
		pack.SetUint8(3)
	})
}

func (this *TangentPC) packetCommonNoSeq(Cmd, Seq uint16, bin []byte) []byte {
	return GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint8(2)
		pack.SetUint16(this.sdk.CMainVer)
		pack.SetUint16(Cmd)
		pack.SetUint16(Seq)
		pack.SetUint32(uint32(this.info.LongUin))
		pack.SetBytes([]byte{0x02, 0x00, 0x00})
		pack.SetUint32(this.sdk.DwClientType)
		pack.SetUint32(this.sdk.DwPubNo)
		pack.SetBytes(util.Encrypt(this.teaKey.SessionKey, bin))
		pack.SetUint8(3)
	})
}

func (this *TangentPC) PacketHttpConn(bin ...[]byte) []byte {
	return GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetBytes([]byte{0x28})
		for _, bytes := range bin {
			pack.SetUint32(uint32(len(bytes)))
		}
		for _, bytes := range bin {
			pack.SetBytes(bytes)
		}
		pack.SetBytes([]byte{0x29})
	})
}
