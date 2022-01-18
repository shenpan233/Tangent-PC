/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/8 10:12
  @Notice:
*/

package test

import (
	"Tangent-PC/PCQQ"
	"Tangent-PC/model"
	util "Tangent-PC/utils"
	"Tangent-PC/utils/GuLog"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	GuLog.Config(true, "")
	tgt := new(model.TgtInfo)
	//这行有我token省略了
	qq := PCQQ.New("", model.Computer{
		ComputerId:   util.HexToBin("9B 13 EE DA 00 00 00 00 00 00 00 00 00 00 00 00"),
		ComputerIdEx: util.HexToBin("3B AE C1 AD 3C 07 44 EE 29 BE C0 38 4E F2 4A 1A"),
		DeviceID:     util.GetRandomBin(32),
		ComputerName: "Beijing University",
		MacGuid:      util.GetRandomBin(16),
	})
	qq.U948()
	if qq.PingServer() {
		qq.LoginByToken(tgt)
		if qq.ChangeOnlineStatus(PCQQ.Online) {
			for {
				qq.HeatBoat()
				time.Sleep(time.Minute)
			}
		}
	}
	select {}
}
