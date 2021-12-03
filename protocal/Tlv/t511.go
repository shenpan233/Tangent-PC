/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		DualStackCheck
* @Creat:   2021/11/26 0026 23:29
 */
package Tlv

import "Tangent-PC/utils/GuBuffer"

func GetTlv511() []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint8(0x0A)
	pack.SetUint32(0)
	pack.SetUint8(3)
	return pack.ToTlv(0x05_11)
}
