/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		ComputerGuid
* @Creat:   2021/12/18 12:02
 */

package Tlv

import (
	util "Tangent-PC/utils"
	"Tangent-PC/utils/GuBuffer"
)

func GetTlv1A(tgtKey, tlv15 []byte) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetBytes(util.Encrypt(tgtKey, tlv15))
	return pack.ToTlv(0x1A)
}
