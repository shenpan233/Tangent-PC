//goreleaser --snapshot --skip-publish --rm-dist

package main

import (
	"bufio"
	"fmt"
	"github.com/shenpan233/Tangent-PC"
	"github.com/shenpan233/Tangent-PC/model"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuLog"
	"github.com/shenpan233/Tangent-PC/utils/GuStr"
	"os"
	"time"
)

func main() {
	GuLog.Config(true, "", true)
	computer := model.Computer{
		ComputerId:   util.HexToBin("9B 13 EE DA 00 00 00 00 00 00 00 00 00 00 00 00"),
		ComputerIdEx: util.HexToBin("3B AE C1 AD 3C 07 44 EE 29 BE C0 38 4E F2 4A 1A"),
		DeviceID:     util.GetRandomBin(32),
		ComputerName: "Beijing University",
		MacGuid:      util.GetRandomBin(16),
	}
ask:
	GuLog.Notice("请选择一种登录方式:1.二维码登录 2.TGTGT令牌登录")
	loginType := 0
	fmt.Scanf("%d", &loginType)
	switch loginType {
	case 1:
		client := Tangent_PC.New("0", computer).U948()
		QrLogin(client)
	case 2:
		GuLog.Notice("请输入QQ")
		account, _ := bufio.NewReader(os.Stdin).ReadString('\r') // 回车结束
		client := Tangent_PC.New(GuStr.Between(account, "\n", "\r"), computer).U948()
		GuLog.Notice("请输入TGTGT令牌:")
		TgtData, _ := bufio.NewReader(os.Stdin).ReadString(13) // 回车结束

		TgtInfo := (new(model.TgtInfo)).Decode(TgtData)
		if TgtInfo == nil {
			GuLog.Error("TGTToken不符规范 请检查")
		} else {
			TGTGTLogin(client, TgtInfo)
		}
	default:
		goto ask
	}

	select {}
}

func QrLogin(client *Tangent_PC.TangentPC) {
	client.PingServer()
	resp := client.FetchQRCode()
	if err := os.WriteFile("./QRCode.png", resp.QRCode, os.FileMode(0777)); err != nil {
		GuLog.Error(err.Error())
		return
	}
	go func() {
		for client.CheckQRCode(resp) != Tangent_PC.QROk {
			GuLog.Debug("wait for login")
			time.Sleep(3 * time.Second)
		}
		{
			if isLogin, tgt := client.QRLogin(); isLogin == nil {
				GuLog.Debug(tgt.Encode())
				GuLog.Clear()
				Online(client)
			}
		}
	}()
}

func TGTGTLogin(client *Tangent_PC.TangentPC, tgt *model.TgtInfo) {
	if !client.PingServer() {
		GuLog.Error("Fail to PingServer")
		return
	}
	if code, err := client.LoginByToken(tgt); code != 0 {
		GuLog.ErrorF("Fail to login because %s", err.Error())
		return
	} else {
		Online(client)
	}
}

func HeartBoat(client *Tangent_PC.TangentPC) {
	for client.HeatBoat() {
		GuLog.InfoF("[QQ=%d] => HeartBoat SuccessFully", client.GetSelfInfo().LongUin)
		time.Sleep(3 * time.Minute)
	}
}

func Online(client *Tangent_PC.TangentPC) {
	if client.ChangeOnlineStatus(Tangent_PC.Online) {
		client.Hook(func(Msg model.GroupMsg) {
			GuLog.NoticeF("[QQ=%d] <= %s(%d)\nFrom:%s(%d) => %s", Msg.Account, Msg.GroupName, Msg.GroupUin, Msg.FromName, Msg.SenderUin, Msg.Msg)
			if Msg.SenderUin == client.Uin {
				return
			}
		})
		go HeartBoat(client)
	}
}

func GroupMsgHook(Msg model.GroupMsg) {

}
