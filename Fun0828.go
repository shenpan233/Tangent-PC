/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2021/12/31 23:55
  @Notice: 0828组包和解包
*/

package Tangent_PC

import (
	"errors"
	"github.com/shenpan233/Tangent-PC/model"
	"github.com/shenpan233/Tangent-PC/protocal/Tlv"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

//pack0828
//	tgt=>tgt参数
func (this *TangentPC) pack0828(tgt *model.TgtInfo) (SsoSeq uint16, buffer []byte) {
	this.sig.BufSession = tgt.BufSession
	this.teaKey.SessionKey = tgt.BufSessionKey
	return this.packetCommon(0x08_28, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetBytes([]byte{0x00, 0x30, 0x00, 0x3A, 0x00, 0x38})
		pack.SetBytes(this.sig.BufSession)
		pack.SetBytes(util.Encrypt(tgt.BufSessionKey, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
			pack.SetBytes(Tlv.GetTlv7Tgt(&tgt.BufTgt))
			pack.SetBytes(Tlv.GetTlvC(this.info.Computer.ConnectIp))
			pack.SetBytes(Tlv.GetTlv15(&this.info.Computer))
			pack.SetBytes(Tlv.GetTlv36LoginReason())
			pack.SetBytes(Tlv.GetTlv18Ping(this.info.LongUin, this.sdk, uint16(this.info.RedirectIp.Len())))
			pack.SetBytes(Tlv.GetTlv1FDeviceID(this.info.Computer.DeviceID))
			pack.SetBytes(Tlv.GetTlv105vec0x12c(tgt.Buf0102, tgt.Buf0202))
			QDData := Tlv.GetTlv32QDData(this.info.Computer.ComputerIdEx, this.sdk)
			pack.SetBytes(Tlv.GetTlv10B(false, this.sdk, &tgt.BufTgt, &QDData))
			pack.SetBytes(Tlv.GetTlv2D())
		})))
	}))
}

func (this *TangentPC) unpack0828(bin []byte, tgt *model.TgtInfo) (result uint8, err error) {
	pack := GuBuffer.NewGuUnPacket(util.Decrypt(tgt.BufTgTgTKey, bin[3:]))
	result = pack.GetUint8()
	if result == 0 {
		//GuLog.InfoF("NewBufSessionKey=%X\n", this.teaKey.SessionKey)
	} else {
		pack = GuBuffer.NewGuUnPacket(util.Decrypt(tgt.BufSessionKey, bin[3:]))
		result = pack.GetUint8()
	}
	GuBuffer.TlvEnum(pack.GetAll(), map[uint16]func(pack *GuBuffer.GuUnPacket){
		0x01_00: func(TlvPack *GuBuffer.GuUnPacket) {
			TlvPack.Skip(8)
			err = errors.New(string(TlvPack.GetToken()))
		},
		0x01_0C: func(TlvPack *GuBuffer.GuUnPacket) {
			TlvPack.Skip(2)
			this.teaKey.SessionKey = TlvPack.GetBin(16)

		},
		0x01_05: func(TlvPack *GuBuffer.GuUnPacket) {
			TlvPack.Skip(4)
			tgt.Buf0102 = TlvPack.GetToken()
			tgt.Buf0202 = TlvPack.GetToken()
		},
	})
	return
}
