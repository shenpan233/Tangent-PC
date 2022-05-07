package Tlv

import (
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

func GetTlv551(DwSSOVersion uint32) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint32(DwSSOVersion)
	pack.SetBytes(util.HexToBin("EF 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 78 43 31 78 76 71 43 34 66 31 02 7B 5D 71 00 08 00 02 38 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 04 35 31 35 31"))
	return pack.ToTlv(0x05_51)
}
