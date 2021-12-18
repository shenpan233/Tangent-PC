/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		PCName
* @Creat:   2021/12/10 23:28
 */

package Tlv

import (
	"Tangent-PC/utils/GuBuffer"
)

func GetTlv30FPcName(ComputerName string) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(uint16(uint32(len(ComputerName) + 2)))
	pack.SetString(ComputerName)
	return pack.ToTlv(0x30f)
}
