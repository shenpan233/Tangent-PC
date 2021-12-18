/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		0819_tlv303
* @Creat:   2021/12/18 11:58
 */

package Tlv

import "Tangent-PC/utils/GuBuffer"

func GetTl303(sigQR303 []byte) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetBytes(sigQR303)
	return pack.ToTlv(0x303)
}
