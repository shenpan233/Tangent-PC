/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		ComputerGuid
* @Creat:   2021/12/18 12:02
 */

package Tlv

import (
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

func GetTlv1AComputerGuid(tgtKey, tlv15 []byte) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetBytes(util.Encrypt(tgtKey, tlv15))
	return pack.ToTlv(0x1A)
}
