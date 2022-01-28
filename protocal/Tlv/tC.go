/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 12:02
  @Notice:	PingRedirect
*/

package Tlv

import (
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

func GetTlvC(IP string) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(2)
	pack.SetBytes([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	pack.SetUint32(uint32(util.IpToInt(IP)))
	pack.SetUint16(8000)
	pack.SetUint32(0)
	return pack.ToTlv(0x00_0C)
}
