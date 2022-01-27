/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:	    qr_login
* @Creat:   2021/11/26 0026 23:05
 */

package Tlv

import "Tangent-PC/utils/GuBuffer"

func GetTlv4() []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(0)
	pack.SetSToken("qr_login")
	return pack.ToTlv(0x04)
}
