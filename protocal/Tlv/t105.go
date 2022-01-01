/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 12:20
  @Notice:vec0x12c
*/

package Tlv

import (
	util "Tangent-PC/utils"
	"Tangent-PC/utils/GuBuffer"
)

func GetTlv105() []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(1)
	pack.SetBytes([]byte{0x01, 0x02})
	{
		pack.SetUint16(0x00_14)
		pack.SetUint8(0x01)
		pack.SetUint8(0x01)
		pack.SetToken(util.GetRandomBin(16))
	}
	{
		pack.SetUint16(0x00_14)
		pack.SetUint8(0x01)
		pack.SetUint8(0x02)
		pack.SetToken(util.GetRandomBin(16))
	}

	return pack.ToTlv(0x01_05)
}
