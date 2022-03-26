/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/3/12 13:02
  @Notice:  Windows的彩色打印 有锁性能较低
*/

package GuLog

import (
	"sync"
	"syscall"
	"time"
)

const (
	foregroundBlue      = 1
	foregroundGreen     = 2
	foregroundRed       = 4
	foregroundIntensity = 8 //前景高亮
	backgroundBlue      = 16
	backgroundGreen     = 32
	backgroundRed       = 64
	backgroundIntensity = 128 //背景高亮
	foregroundWhite     = foregroundBlue | foregroundGreen | foregroundRed | foregroundIntensity
	commonLvbUnderscore = 0x8000 //下划线
)

type linuxColor int

var (
	console     = sync.RWMutex{}
	kernel32    = syscall.NewLazyDLL(`kernel32.dll`)
	proc        = kernel32.NewProc(`SetConsoleTextAttribute`)
	CloseHandle = kernel32.NewProc(`CloseHandle`)
	winColor    = map[linuxColor]int{
		red:    foregroundWhite | foregroundIntensity | backgroundRed,
		blue:   foregroundBlue | foregroundIntensity,
		green:  foregroundGreen | foregroundIntensity,
		grey:   foregroundWhite | foregroundRed | foregroundIntensity,
		yellow: foregroundWhite | foregroundIntensity | backgroundRed | backgroundGreen,
	}
)

func SetConsoleColor(id int) {
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(id))
	CloseHandle.Call(handle)
}

func WinPrint(file, funcName, Tag, msg string, line, color int) {
	console.Lock()
	winPrintTime()
	SetConsoleColor(foregroundWhite | commonLvbUnderscore)
	print(file + ":")
	print(line)

	SetConsoleColor(color)
	print(" [" + Tag + "] ")

	SetConsoleColor(foregroundWhite)
	print(funcName)
	print("\n")

	SetConsoleColor(color)
	print(msg)

	SetConsoleColor(foregroundWhite)
	print("\n\n")

	console.Unlock()
}

func winPrintTime() {
	SetConsoleColor(foregroundWhite | backgroundIntensity)
	print(time.Now().Format(guLog.timeFormat))
	SetConsoleColor(foregroundWhite)
	print(" ")
}
