/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 12:43
  @Notice:	LocalIP本地Ip
*/

package Tlv

import (
	util "Tangent-PC/utils"
	"Tangent-PC/utils/GuBuffer"
)

func GetTlv2D() []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(wSubVer)
	pack.SetUint32(uint32(util.IpToInt("192.168.0.110"))) //反正也查不出来，乱来就行
	return pack.ToTlv(0x00_2D)
}
