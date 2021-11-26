/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		tlv018
* @Creat:   2021/11/26 0026 22:15
 */
package tlv

import (
	"Tangent-PC/model"
	"Tangent-PC/protocal/GuBuffer"
)

func GetTlv18(Uin uint64, sdk *model.Sdk, RedirectCount uint16) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(wSubVer)
	pack.SetUint32(sdk.DwSSOVersion)
	pack.SetUint32(sdk.ServiceId)
	pack.SetUint32(sdk.ClientVer)
	pack.SetUint32(uint32(Uin))
	pack.SetUint16(RedirectCount)
	pack.SetUint16(0)
	return pack.ToTlv(0x00_18)
}
