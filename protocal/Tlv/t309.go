/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		TLV_Ping_Strategy_0x309
* @Creat:   2021/11/26 0026 23:09
 */
package Tlv

import (
	util "Tangent-PC/utils"
	"Tangent-PC/utils/GuBuffer"
)

func GetTlv309PingStrategy(IP string) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(0xA)
	pack.SetUint16(0x4)
	ip := uint32(util.IpToInt(IP))
	pack.SetUint32(ip)
	//先暂时写成这样,后续修改
	pack.SetUint8(0)
	pack.SetUint8(4)
	return pack.ToTlv(0x03_09)
}
