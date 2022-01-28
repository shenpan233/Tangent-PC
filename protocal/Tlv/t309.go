/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:	    TLV_Ping_Strategy_0x309
* @Creat:   2021/11/26 0026 23:09
 */

package Tlv

import (
	"container/list"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

//GetTlv309PingStrategy
//	ConnectIp 当前连接的IP
//	RedirectIp 重定向的IP列表
func GetTlv309PingStrategy(ConnectIp string, RedirectIp *list.List) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(0xA)
	pack.SetUint16(0x4)
	pack.SetUint32(uint32(util.IpToInt(ConnectIp))) /*当前连接IP*/
	if RedirectIp != nil {
		if RedirectIp.Len() > 0 {
			pack.SetUint8(1)
			pack.SetUint16(uint16(RedirectIp.Len())) /*重定向IP个数*/
			for element := RedirectIp.Front(); element != nil; element = element.Next() {
				pack.SetUint32(uint32(util.IpToInt(element.Value.(string))))
			}
		} else {
			pack.SetUint8(0) /*没有向其他服务器发0825*/
		}
	}
	pack.SetUint8(2)
	return pack.ToTlv(0x03_09)
}
