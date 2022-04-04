/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 15:15
  @Notice:  心跳包
*/

package Tangent_PC

import (
	"github.com/shenpan233/Tangent-PC/model"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
	"github.com/shenpan233/Tangent-PC/utils/GuLog"
)

func (this *TangentPC) pack0058HeatBoat() (SsoSeq uint16, buffer []byte) {
	return this.packetCommonEnc(0x00_58, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetBytes([]byte{0x0, 0x1, 0x0, 0x1})
	}))
}

func (this *TangentPC) unpack0058(bin []byte) (heartBoatSe uint8) {
	GuBuffer.NewGuUnPacketFun(util.Decrypt(this.teaKey.SessionKey, bin), func(pack *GuBuffer.GuUnPacket) {
		heartBoatSe = pack.GetUint8()
		switch heartBoatSe {
		case LoginSuc:
			return
		case 0x11: //掉线
			tgtInfo := model.TgtInfo{
				BufTgTgTKey:   this.sig.BufTgTGTKey,
				BufTgt:        this.sig.BufTgt,
				BufGTKeyST:    this.sig.BufTgTGTKey,
				BufSessionKey: this.teaKey.SessionKey,
				BufSession:    this.sig.BufSession,
				Buf0102:       this.sig.Buf0102,
				Buf0202:       this.sig.Buf0202,
			}
			this.udper.Break()
			this.PingServer()
			this.LoginByToken(&tgtInfo)
			this.ChangeOnlineStatus(Online)
		}
		GuLog.Info("heartBoat heartBoat:", heartBoatSe, "\n", pack.GetAllHex())
	})
	return
}
