/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		AtUin
* @Creat:   2021/12/10 23:31
 */

package Tlv

import "github.com/shenpan233/Tangent-PC/utils/GuBuffer"

const (
	tlvVersion2 = 0x2
)

func GetTlv5Uin(qqUin uint64) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(tlvVersion2)
	pack.SetUint32(uint32(qqUin))
	return pack.ToTlv(0x5)
}
