/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		Misc_Flag
* @Creat:   2021/12/18 12:11
 */

package Tlv

import "Tangent-PC/utils/GuBuffer"

func GetTlv312MiscFlag() []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetBytes([]byte{1, 0, 0, 0, 0})
	return pack.ToTlv(0x312)
}
