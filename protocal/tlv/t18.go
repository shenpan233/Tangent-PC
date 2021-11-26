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

func GetTlv18(sdk *model.Sdk) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(wSubVer)
	pack.SetUint32(sdk.DwSSOVersion)
	pack.SetUint32(sdk.ServiceId)
	pack.SetUint32(sdk.ClientVer)

	return pack.ToTlv(0x00_18)
}
