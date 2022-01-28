/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:	    GUID_Ex
* @Creat:   2021/12/18 12:20
 */

package Tlv

import (
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

const (
	tick = 0x00_00_00_1F
)

func GetTlv313GUIDEx(MacGuid []byte) []byte {
	GUIDCount := uint8(1)
	pack := GuBuffer.NewGuPacket()
	pack.SetUint8(1)
	pack.SetUint8(GUIDCount)
	pack.SetUint8(2) //GuidIndex
	pack.SetToken(MacGuid)
	pack.SetUint32(tick)
	return pack.ToTlv(0x313)
}
