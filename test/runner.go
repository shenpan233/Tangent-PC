//goreleaser --snapshot --skip-publish --rm-dist

package test

import (
	client2 "github.com/shenpan233/Tangent-PC/PCQQ"
	"github.com/shenpan233/Tangent-PC/model"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuLog"
	"os"
	"time"
)

func main() {
	GuLog.Config(true, "")
	client := client2.New("0", model.Computer{
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
		os.WriteFile("./QRCode.png", resp.QRCode, os.FileMode(0777))
		go func() {
			for client.CheckQRCode(resp) != client2.QROk {
				time.Sleep(3 * time.Second)
			}
			if client.QRLogin() {
				if client.ChangeOnlineStatus(client2.Online) {
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
