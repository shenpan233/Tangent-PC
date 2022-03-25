/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/3/19 14:49
  @Notice:  notice
*/

package QunInfo

import (
	"github.com/shenpan233/Tangent-PC/model"
	"testing"
)

func TestGetGroupList(t *testing.T) {
	GetGroupList("256783314", model.CommonWebKey{
		Skey:  "@kZpiOmV1S",
		PSkey: "*q7ZTcYLDBPEwm43J5aS2jVzdYKhU5TliraRTfMqPMQ_",
	})
}
