/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 11:49
  @Notice:	Tgt
*/

package Tlv

import "Tangent-PC/utils/GuBuffer"

func GetTlv7Tgt(bufTgt *[]byte) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetBytes(*bufTgt)
	return pack.ToTlv(0x00_07)
}
