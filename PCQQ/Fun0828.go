/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2021/12/31 23:55
  @Notice: 0828组包和解包
*/

package PCQQ

import (
	"Tangent-PC/protocal/Tlv"
	util "Tangent-PC/utils"
	"Tangent-PC/utils/GuBuffer"
	"Tangent-PC/utils/GuLog"
)

func (this *TangentPC) pack0828(tgt *tgtInfo) (SsoSeq uint16, buffer []byte) {
	return this.packetCommon(0x08_28, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetBytes([]byte{0x00, 0x30, 0x00, 0x3A, 0x00, 0x38})
		pack.SetBytes(this.sig.BufSession)
		pack.SetBytes(util.Encrypt(this.teaKey.SessionKey, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
			pack.SetBytes(Tlv.GetTlv7Tgt(&tgt.bufTgt))
			pack.SetBytes(Tlv.GetTlvC(this.info.Computer.ConnectIp))
			pack.SetBytes(Tlv.GetTlv15(&this.info.Computer))
			pack.SetBytes(Tlv.GetTlv36LoginReason())
			pack.SetBytes(Tlv.GetTlv18Ping(this.info.LongUin, this.sdk, uint16(this.info.RedirectIp.Len())))
			pack.SetBytes(Tlv.GetTlv1FDeviceID(this.info.Computer.DeviceID))
			pack.SetBytes(Tlv.GetTlv105())
			pack.SetBytes(Tlv.GetTlv10B(true, this.sdk, &tgt.bufTgt))
			pack.SetBytes(Tlv.GetTlv2D())
		})))
	}))
}

//
func (this *TangentPC) unpack0828(bin []byte, tgt *tgtInfo) (result uint8) {
	pack := GuBuffer.NewGuUnPacket(util.Decrypt(tgt.bufTgTgTKey, bin[3:]))
	result = pack.GetUint8()
	if result == 0 {
		for pack.GetLen() > 0 {
			tlv := pack.GetTlv()
			pack := GuBuffer.NewGuUnPacket(tlv.Value)
			switch tlv.Tag {
			case 0x01_0C:
				pack.GetInt16()
				this.teaKey.SessionKey = pack.GetBin(16)
			default:
				GuLog.Warm("un0828", "%X", pack.GetAll())
			}
		}
		GuLog.Info("unpack0828", "NewBufSessionKey=%X\n", this.teaKey.SessionKey)
	}
	return
}
