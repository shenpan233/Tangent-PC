/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		验证二维码的参数
* @Creat:   2021/12/10 21:40
 */

package Tlv

import "github.com/shenpan233/Tangent-PC/utils/GuBuffer"

func GetTlv30(sig []byte) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetBytes(sig)
	return pack.ToTlv(0x0030)
}
