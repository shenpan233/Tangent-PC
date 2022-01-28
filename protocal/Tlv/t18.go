/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		tlv018
* @Creat:   2021/11/26 0026 22:15
 */

package Tlv

import (
	"github.com/shenpan233/Tangent-PC/model"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

func GetTlv18Ping(Uin uint64, sdk *model.Version, RedirectCount uint16) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(wSubVer)
	pack.SetUint32(sdk.DwSSOVersion)
	pack.SetUint32(sdk.ServiceId)
	pack.SetUint32(sdk.ClientVer)
	pack.SetUint32(uint32(Uin))
	pack.SetUint16(RedirectCount)
	pack.SetUint16(wUnknown1)
	return pack.ToTlv(0x00_18)
}
