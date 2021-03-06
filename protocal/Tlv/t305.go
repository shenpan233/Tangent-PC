/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		二维码参数
* @Creat:   2021/12/4 0004 15:32
 */

package Tlv

import (
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

//GetTlv305 QRCodeParams
func GetTlv305() []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetBytes(util.HexToBin("00 00 00 00 00 00 00 05 00 00 00 04 00 00 00 00 00 00 00 48 00 00 00 02 00 00 00 02 00 00"))
	return pack.ToTlv(0x305)
}
