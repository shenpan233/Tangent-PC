/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		Udp组件
* @Creat:   2021/11/27 0027 12:03
 */
package udper

import (
	"context"
	"github.com/shenpan233/Tangent-PC/utils/GuLog"
	"net"
	"sync"
	"sync/atomic"
)

func New(host string, set *Set) (udper *Udper) {
	udper = new(Udper)
	if set != nil {
		udper.Set = set
	} else {
		udper.Set = new(Set)
	}
	if udper.BuffMaxSize == 0 {
		udper.BuffMaxSize = 1024
	}
	udper.pull = new(sync.Map)
	udper.Context, udper.CancelFunc = context.WithCancel(context.Background())
	udper.seq = 0
	conn, _ := net.ResolveUDPAddr("udp", host)
	if udper.conn, _ = net.Dial("udp", conn.String()); udper.conn == nil {
		return nil
	} else {
		go udper.recv()
	}
	return
}

/*连其他服务器*/
func (this *Udper) ChangeConnect(host string) bool {
	this.Break()
	this.Context, this.CancelFunc = context.WithCancel(context.Background())
	conn, _ := net.ResolveUDPAddr("udp", host)
	var err error
	if this.conn, err = net.Dial("udp", conn.String()); err != nil {
		GuLog.Error("ChangeConnect", "重连失败,%s", err)
		return false
	} else {
		go this.recv()
		return true
	}
}

/*销毁组件*/
func (this *Udper) Break() {
	this.CancelFunc() /*清理读取器*/
	this.conn.Close()
}

func (this *Udper) GetSeq() uint16 {
	if seq := atomic.AddUint32(&this.seq, 1); seq > 0xFFFF {
		atomic.StoreUint32(&this.seq, 0)
		return 0
	} else {
		return uint16(seq)
	}
}
