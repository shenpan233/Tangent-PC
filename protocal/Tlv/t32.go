/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/14 22:56
  @Notice:	QD算法
*/

package Tlv

import (
	"Tangent-PC/model"
	util "Tangent-PC/utils"
	"Tangent-PC/utils/GuBuffer"
)

func GetTlv32QDData(Machine []byte, sdk *model.Version) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetBytes([]byte{0x3E, 0x00, 0x63, 0x02})
	pack.SetUint32(sdk.DwQdVersion)
	pack.SetBytes([]byte{0x00, 0x04})
	pack.SetBytes([]byte{0x00})
	pack.SetBytes(util.GetRandomBin(2))
	pack.SetBytes(Machine)
	pack.SetUint16(1)
	pack.SetUint32(sdk.DwPubNo)
	pack.SetUint16(uint16(sdk.ClientVer))
	pack.SetBytes([]byte{00, 00, 00, 00, 00, 00, 00, 00})
	pack.SetBytes([]byte{0x07, 0xDE, 0x00, 0x03, 0x00, 0x06, 0x00, 0x01, 0x00, 0x04, 0x00, 0x04, 0x00, 0x04, 0x24, 0x5A, 0x00})
	pack.SetBytes(util.GetRandomBin(16)) //32轮的Tea懒得弄了
	pack.SetUint8(0x68)
	return pack.ToTlv(0x00_32)
}
