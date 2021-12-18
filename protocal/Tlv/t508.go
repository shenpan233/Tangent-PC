/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		验证码之类的
* @Creat:   2021/12/18 11:53
 */

package Tlv

import "Tangent-PC/utils/GuBuffer"

func GetTlv508() []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetBytes([]byte{1, 0, 0, 0, 0})
	return pack.ToTlv(0x508)
}
