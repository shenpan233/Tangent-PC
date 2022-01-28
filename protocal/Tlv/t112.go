/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		组0825返回参数
* @Creat:   2021/12/10 23:25
 */

package Tlv

import (
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

func GetTlv112SigClientAddr(BufSigClientAddr *[]byte) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetBytes(*BufSigClientAddr)
	return pack.ToTlv(0x112)
}
