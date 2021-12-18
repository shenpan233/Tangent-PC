/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		Uin
* @Creat:   2021/12/10 23:31
 */

package Tlv

import "Tangent-PC/utils/GuBuffer"

func GetTlv5Uin(qqUin uint64) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(2)
	pack.SetUint32(uint32(qqUin))
	return pack.ToTlv(0x5)
}
