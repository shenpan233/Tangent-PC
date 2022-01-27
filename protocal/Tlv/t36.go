/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 12:10
  @Notice:  LoginReason
*/

package Tlv

import "Tangent-PC/utils/GuBuffer"

func GetTlv36LoginReason() []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(2)
	pack.SetUint16(1)
	pack.SetUint32(5)
	pack.SetBytes([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	return pack.ToTlv(0x00_36)
}
