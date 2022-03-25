/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/3/25 22:00
  @Notice:  notice
*/

package test

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
)

func TestErr(t *testing.T) {
	err()
}
func err() {
	defer func() {
		if err := recover(); err != nil {
			num := 3
			for i := 0; i < num; i++ {
				_, file, line, _ := runtime.Caller(3 + i)
				fmt.Println(file + ":" + strconv.Itoa(line))
			}
			fmt.Println(err)
		}
	}()
	mistake()
}

var a *int

func mistake() {
	defer func() {
		*a = 1
	}()
}
