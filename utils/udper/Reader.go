/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		接收
* @Creat:   2021/11/27 0027 13:17
 */
package udper

import (
	"bytes"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
	"github.com/shenpan233/Tangent-PC/utils/GuLog"
)

func (this *Udper) recv() {
	select {
	case <-this.Context.Done():
		break
	default:
		for {
			bin := make([]byte, this.BuffMaxSize)
			if read, err := this.conn.Read(bin[:]); err == nil {
				receiver := bin[:read]
				/*报文判断*/
				if bytes.Contains(receiver[:1], []byte{2}) && bytes.Contains(receiver[read-1:], []byte{3}) {
					receiver = receiver[1 : read-1]
				}
				/*普通处理*/
				pack := GuBuffer.NewGuUnPacket(receiver) //申请缓冲器
				pack.GetInt16()                          //Version
				Cmd := pack.GetInt16()                   //命令
				SsoSeq := uint16(pack.GetInt16())
				pack.GetUint32() //QQUin
				/*检查是否要拉取*/
				if value, exits := this.pull.LoadAndDelete(SsoSeq); exits && value != nil {
					puller := value.(chan []byte)
					puller <- pack.GetAll()
					continue
				}

				/*无需拉取*/
				if this.UdpRecv != nil {
					go this.UdpRecv(Cmd, SsoSeq, pack)
				}
			} else {
				GuLog.Error("UdpEr", "%s", err.Error())
			}
		}

	}
}
