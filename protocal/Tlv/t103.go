/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:	    DhPublicKey
* @Creat:   2021/11/26 0026 23:24
 */

package Tlv

import "github.com/shenpan233/Tangent-PC/utils/GuBuffer"

func GetTlv103(PublicKey *[]byte) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetBytes(*PublicKey)
	return pack.ToTlv(0x01_03)
}
