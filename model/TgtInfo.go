/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/8 9:59
  @Notice:  0828令牌验证
*/

package model

import (
	"encoding/json"
)

type (
	TgtInfo struct {
		BufTgTgTKey      []byte `json:"bufTgTgTKey"`
		BufTgt           []byte `json:"bufTgt"`
		BufGTKeyST       []byte `json:"bufGTKeyST"`
		BufServiceTicket []byte `json:"bufServiceTicket"`
		BufSessionKey    []byte `json:"bufSessionKey"`
		BufSession       []byte `json:"bufSession"`
		Buf0102          []byte `json:"buf0102"`
		Buf0202          []byte `json:"buf0202"`
		dataJson
	}
)

func (this *TgtInfo) Encode() string {
	if data, err := json.Marshal(this); err != nil {
		return ""
	} else {
		return string(data)
	}
}
func (this *TgtInfo) Decode(data string) bool {
	if err := json.Unmarshal([]byte(data), this); err != nil {
		return false
	}
	return true
}
