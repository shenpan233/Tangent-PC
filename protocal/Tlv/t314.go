/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		二维码验证
* @Creat:   2021/12/10 22:27
 */

package Tlv

import "github.com/shenpan233/Tangent-PC/utils/GuBuffer"

func GetTlv314() []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(0)
	return pack.ToTlv(0x314)
}
