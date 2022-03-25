package Tangent_PC

import (
	"container/list"
	"github.com/shenpan233/Tangent-PC/model"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/udper"
	"strconv"
)

var WaitTime = uint32(3) /*三秒延迟*/

type TangentPC struct {
	Uin    uint64
	sdk    *model.Version
	info   *model.Information
	sig    *model.Sig
	udper  *udper.Udper
	teaKey *model.TeaKey
	cache  struct {
		groupList map[uint64]string
		member    map[uint64]map[uint64]*model.GroupMember
	}
	handle map[uint16]unpack
	hook   HOOK
}

//New 新建PC_QQ协议类
func New(Account string, Computer model.Computer) (this *TangentPC) {
	this = new(TangentPC)
	/*通讯器部分*/
	{
		if this.udper = udper.New(model.TxServer[1]+":8000", &udper.Set{
			BuffMaxSize: 2048,
			UdpRecv:     nil,
		}); this.udper == nil {
			//GuLog.Error("New", "服务器连接失败")
			return nil
		}
	}

	/*硬件信息等部分*/
	{
		this.sdk = new(model.Version)
		Computer.RedirectIp = list.New()
		this.info = &model.Information{
			LongUin: func() uint64 {
				Uint, _ := strconv.ParseUint(Account, 10, 64)
				return Uint
			}(),
			Account:  Account,
			PassWord: nil,
			Computer: Computer,
		}
		this.Uin = this.info.LongUin
		this.info.ComputerId = append(this.info.ComputerId[:4], 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)

	}

	/*Tea秘钥申请*/
	{
		this.sig = new(model.Sig)
		this.teaKey = new(model.TeaKey)
		this.teaKey.PublicKey = util.HexToBin("03 1F 06 FA 3B 19 BF F9 2C 7C 02 7D 5D EA C5 60 83 52 86 C1 BF 75 CA 2A 96")
		this.teaKey.ShareKey = util.HexToBin("6C 3E 9F 64 1C 27 F9 CA D6 B6 37 8A A7 74 D0 04")
	}
	/*缓存空间申请*/
	{
		this.info.SelfWebKey = new(model.WebKey)
		this.cache.groupList = make(map[uint64]string)
		this.cache.member = make(map[uint64]map[uint64]*model.GroupMember)
	}

	return
}

//GetSelfInfo 获取账号信息
func (this *TangentPC) GetSelfInfo() model.Information {
	return *this.info
}