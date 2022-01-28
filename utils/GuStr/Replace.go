/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/28 12:20
  @Notice:  替换器组件
*/

package GuStr

import (
	"strings"
)

type StrReplace struct {
	data string
	old  []string
	new  []string
}

//NewReplace 新建一个替换器
func NewReplace(data string) (this *StrReplace) {
	//什么都没有换个蛇皮怪，浪费性能吗
	if data == "" {
		return
	}
	this = new(StrReplace)
	this.data = data
	return
}

func (this *StrReplace) OldData(olds ...string) (r *StrReplace) {
	this.old = olds
	return this
}

func (this *StrReplace) NewData(news ...string) (r *StrReplace) {
	if len(news) != len(this.old) {
		panic("搞什么鬼呢(╬￣皿￣)？新旧数据数量都不一致")
		return nil
	}
	this.new = news
	return this
}

func (this *StrReplace) MapData(NewOlds map[string]string) (r *StrReplace) {
	for old, New := range NewOlds {
		this.old = append(this.old, old)
		this.new = append(this.new, New)
	}
	return this
}

func (this *StrReplace) Replace() string {
	replaceData := make([]string, 0)
	for i := range this.old {
		replaceData = append(replaceData, this.old[i], this.new[i])
	}
	return strings.NewReplacer(replaceData...).Replace(this.data)
}
