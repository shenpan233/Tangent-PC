/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 12:24
  @Notice:	QDLoginFlag
*/

package Tlv

import (
	"Tangent-PC/model"
	"Tangent-PC/utils/GuBuffer"
)

func GetTlv10B(LoginByQRCode bool, version *model.Version, bufTgt *[]byte) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(2)
	pack.SetBytes(version.ClientMd5)
	pack.SetUint8(CreateQDFlag(1, version.ClientMd5, *bufTgt))
	pack.SetBytes([]byte{0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02})
	if LoginByQRCode {
		pack.SetBytes([]byte{0, 0})
	}
	pack.SetUint32(0)
	return pack.ToTlv(0x01_0B)
}

func CreateQDFlag(VerType byte, clientMd5, tgt []byte) byte {
	VerType = VerType % 100
	for i := 0; i < len(tgt); i++ {
		VerType = VerType ^ (tgt[i])
	}
	for i := 0; i < len(clientMd5); i++ {
		VerType = VerType ^ (clientMd5[i])
	}
	return VerType
}
