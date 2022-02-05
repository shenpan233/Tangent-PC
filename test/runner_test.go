//goreleaser --snapshot --skip-publish --rm-dist

package test

import (
	"github.com/shenpan233/Tangent-PC"
	"github.com/shenpan233/Tangent-PC/model"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuLog"
	"os"
	"testing"
	"time"
)

func TestQRCode_Login(t *testing.T) {
	GuLog.Config(true, "")
	client := Tangent_PC.New("0", model.Computer{
		ComputerId:   util.HexToBin("9B 13 EE DA 00 00 00 00 00 00 00 00 00 00 00 00"),
		ComputerIdEx: util.HexToBin("3B AE C1 AD 3C 07 44 EE 29 BE C0 38 4E F2 4A 1A"),
		DeviceID:     util.GetRandomBin(32),
		ComputerName: "Beijing University",
		MacGuid:      util.GetRandomBin(16),
	})
	if client == nil {
		return
	}
	client.U948()

	if client.PingServer() {
		GuLog.Notice("System", "[QQ=%s]Ping成功", "0")
		resp := client.FetchQRCode()
		if err := os.WriteFile("./QRCode.png", resp.QRCode, os.FileMode(0777)); err != nil {
			GuLog.Error("TestQRCode", err.Error())
			return
		}
		go func() {
			for client.CheckQRCode(resp) != Tangent_PC.QROk {
				time.Sleep(3 * time.Second)
			}
			if client.QRLogin() == nil {
				if client.ChangeOnlineStatus(Tangent_PC.Online) {
					for {
						client.HeatBoat()
						time.Sleep(time.Minute)
					}
				}

			}
		}()
	}

	select {}
}
