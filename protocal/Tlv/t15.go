/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		TLV_ComputerGuid 电脑信息
* @Creat:   2021/12/4 0004 15:40
 */

package Tlv

import (
	"github.com/shenpan233/Tangent-PC/model"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
	"hash/crc32"
)

//TLV_ComputerGuid

func GetTlv15(computer *model.Computer) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(wSubVer)
	/*CComputerIDGenerator::Generate*/
	pack.SetUint8(1)
	pack.SetUint32(crc32.ChecksumIEEE(computer.ComputerId))
	pack.SetToken(computer.ComputerId)
	//bufComputerIDEx
	pack.SetUint8(2)
	pack.SetUint32(crc32.ChecksumIEEE(computer.ComputerIdEx))
	pack.SetToken(computer.ComputerIdEx)
	return pack.ToTlv(0x15)
}
