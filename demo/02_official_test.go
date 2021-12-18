/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:
* @Creat:   2021/12/11 20:28
 */

package demo

import (
	util "Tangent-PC/utils"
	"fmt"
	"testing"
)

func TestCreatOfficial(t *testing.T) {
	//Tlv.CreateOfficial()
	fmt.Println(util.GetCrc32([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))

}
