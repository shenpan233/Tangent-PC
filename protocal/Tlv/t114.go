/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:	    DHParams
* @Creat:   2021/11/26 0026 23:23
 */

package Tlv

import (
	"Tangent-PC/utils/GuBuffer"
)

func GetTlv114DHParams(PublicKey *[]byte) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetBytes(GetTlv103(PublicKey))
	return pack.ToTlv(0x01_14)
}
