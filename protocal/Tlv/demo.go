/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		一个可以抄袭的模板
* @Creat:   2021/10/29 0029 23:07
 */

package Tlv

import (
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

func GetTlv() []byte {
	pack := GuBuffer.NewGuPacket()
	return pack.ToTlv(0x0)
}

func Verify(pack *GuBuffer.GuUnPacket) bool {
	if pack == nil {
		return false
	}
	return true
}
