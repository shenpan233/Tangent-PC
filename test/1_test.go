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
	tgt.Decode(`{"bufTgTgTKey":"PilUamYqJUVaVklXNlVBZg==","bufTgt":"AARh5S8zeOdI2QAAAAAAeCVfxk0M9t9BoP8aqB8M47F8KiqJM0LZjM9QGfYykRoI8Vw25ZngChoVEFLvkzP3qPJmusUpG3UpgqOvw+eSV8f4gEq+qER6bGvxFEB1ZpG2pq1BANlCFFI67BtA30TqfHsEmmiFOh9M7w/Xx3EfchGf5WoK4DuDsA==","bufGTKeyST":"fik2KGlSJCksc0Y7K1YqWA==","bufServiceTicket":"AAFh5S8zAHAQD2WRwiTqJKoNqhEc5z0Nk8K0humzJos2pf0d75SjaX5CTWGu5KyzE20qSO5ZC943GFwzbaYeBo8B48W7c9UVV2YA9teOj97/ph0NtZdCvWs3sla0fILme3csCQ49ix7TsWBfRNe9qNcAuMTIz60D","bufSessionKey":"B/qivV+5icJnQS2imh/vdA==","bufSession":"0W+61ol2Y0m08snD/R61rJnKn4YLNS+G/x44YnrXMZ8AqjSoRicnW56WkocEIHG0nMN7GQAZfjo=","buf0102":"AQABAzwBAwAAR5R/bXYbXoRVnXb+vi3PgzMN8R5BRw70TPV1JkOJzRbZhCpQ9Z5jJhBlhDvu3n3aBPnvEHkACco=","buf0202":"AQACAzwBAwAAq7KIoNqPek0zFnx7LfnK6ZRI04kFpUGZx+2j3VxSWE9RCITX3ryLrFYeKB8gUugFPAviZwSzUv8="}`)
	qq := PCQQ.New("2849567593", model.Computer{
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
