/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 12:20
  @Notice:	vec0x12c
*/

package Tlv

import (
	util "Tangent-PC/utils"
	"Tangent-PC/utils/GuBuffer"
	"fmt"
)

func GetTlv105(buf0102, buf0202 []byte) []byte {
	fmt.Println(util.BinToHex(buf0102))
	fmt.Println(util.BinToHex(buf0202))

	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(1)
	pack.SetBytes([]byte{0x01, 0x02})
	//TLV_m_vec0x12c(如果0828接收包返回0x0088,下次登陆0828就组上次0828返回的这段0x0088,第一次登陆0828的0105只有0x0030字节)
	if len(buf0102) == 0 || len(buf0202) == 0 {
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
	} else {
		//返回来的还不能直接用我吐了
		pack.SetUint16(64)
		pack.SetUint8(2)
		pack.SetBytes(buf0102[2:])

		pack.SetUint16(64)
		pack.SetUint8(2)
		pack.SetBytes(buf0202[2:])

	}

	return pack.ToTlv(0x01_05)
}
