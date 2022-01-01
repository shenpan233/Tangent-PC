/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 12:15
  @Notice:DeviceID
*/

package Tlv

import "Tangent-PC/utils/GuBuffer"

//DeviceID

func GetTlv1FDeviceID(DeviceID []byte) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(wSubVer)
	pack.SetBytes(DeviceID)
	return pack.ToTlv(0x00_1F)
}
