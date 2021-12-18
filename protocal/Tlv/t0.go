/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		机器码
* @Creat:   2021/12/10 23:34
 */

package Tlv

import "Tangent-PC/utils/GuBuffer"

func GetTlv0(passwordMD5 []byte) []byte {
	pack := GuBuffer.NewGuPacket()
	return pack.ToTlv(0x0)
}
