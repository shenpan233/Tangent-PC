/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/31 19:01
  @Notice:  资料卡详细信息
*/

package Tangent_PC

import (
	"errors"
	"github.com/shenpan233/Tangent-PC/model"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/Bytes"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

const (
	SummaryNick     = 0x4E_22 //qq昵称
	SummaryQQAge    = 0x65_97 //q龄
	SummaryAge      = 0x4E_45 //年龄
	SummaryEmail    = 0x4E_2B //邮箱
	SummaryBirthday = 0x4E_3F //生日
)

func (this *TangentPC) pack003C(uin uint32, subCmd []uint16) (SsoSeq uint16, buffer []byte) {
	return this.packetCommonEnc(0x00_3C, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint16(1)
		pack.SetUint32(uin)
		pack.SetBytes([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) //22个00
		pack.SetUint16(uint16(len(subCmd)))
		for _, i := range subCmd {
			pack.SetUint16(i)
		}
	}))
}

//QQ信息的解包
func (this *TangentPC) unpack003C(bin []byte) (err error, data map[uint16]interface{}) {
	GuBuffer.NewGuUnPacketFun(util.Decrypt(this.teaKey.SessionKey, bin[:3]), func(pack *GuBuffer.GuUnPacket) {
		if pack.GetUint16() != model.LogicSuc {
			err = errors.New(pack.GetAllHex())
			return
		}
		/*
			跳过这些东西
			01 00
			** ** ** ** //qq
			00 00 00 00
		*/
		pack.Skip(10)
		dataNum := int(pack.GetUint16())
		if dataNum == 0 {
			return
		} else {
			data = make(map[uint16]interface{})
		}
		for i := 0; i < dataNum; i++ {
			tlv := pack.GetTlv()
			tag := uint16(tlv.Tag)
			switch tag {
			case SummaryNick:
				data[tag] = string(tlv.Value)
				break
			case SummaryQQAge:
				data[tag] = Bytes.Bytes2Uint8(tlv.Value)
				break
			case SummaryAge:
				data[tag] = Bytes.Bytes2Uint8(tlv.Value)
				break
			case SummaryBirthday:
				tPack := GuBuffer.NewGuUnPacket(tlv.Value)
				birthday := model.Birthday{}
				birthday.Year = tPack.GetUint32()
				birthday.Month = tPack.GetUint8()
				birthday.Day = tPack.GetUint8()

			}
		}
	})
	return
}
