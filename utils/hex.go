package util

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/shenpan233/Tangent-PC/utils/GuStr"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	matchHexData = regexp.MustCompile("[0-9a-fA-F]+")
)

func GetRandomBin(len int) []byte {
	rand.Seed(time.Now().Unix())
	buffer := bytes.NewBuffer(nil)
	for i := 0; i < len; i++ {
		buffer.WriteByte(byte(rand.Intn(255)))
	}
	return buffer.Bytes()
}

func GetRand32() uint32 {
	rand.Seed(GetServerCurTime())
	return uint32(rand.Int())
}

func RandUint32(min, max uint32) uint32 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return uint32(rand.Int31n(int32(max-min)) + int32(min))
}

func GetServerCurTime() int64 {
	return time.Now().Unix()
}

func HexToBin(HexData string) []byte {
	data := matchHexData.FindAllString(HexData, -1)
	tmp := bytes.NewBuffer(nil)
	for _, datum := range data {
		tmp.WriteString(datum)
	}
	decodeString, _ := hex.DecodeString(tmp.String())
	return decodeString
}

func BinToHex(Bin []byte) string {
	return strings.ToUpper(hex.EncodeToString(Bin))
}

func BinToHex2(Bin *[]byte) string {
	return strings.ToUpper(hex.EncodeToString(*Bin))
}

func IpToInt(Ip string) int64 {
	split := strings.Split(Ip, ".")
	if len(split) != 4 {
		return 0
	}
	var intIp []int
	for _, s := range split {
		tmp, _ := strconv.Atoi(s)
		intIp = append(intIp, tmp)
	}
	//JAVA反编译↓
	//Long.parseLong(split[0]) + (Long.parseLong(split[3]) << 24) + (Long.parseLong(split[2]) << 16) + (Long.parseLong(split[1]) << 8);

	return int64(intIp[1] + (intIp[2] << 24) + (intIp[3] << 16) + (intIp[0] << 8))
}

func IntToIp(i int32) string {
	return fmt.Sprintf("%d.%d.%d.%d", (byte)(i>>24), (byte)(i>>16), (byte)(i>>8), (byte)(i>>0))
}

func Guid2Md5Bytes(GuidData string) []byte {
	Guid := GuStr.Between(GuidData, "{", "}")
	if Guid == "" {
		return nil
	} else {
		return HexToBin(strings.ReplaceAll(Guid, "-", ""))
	}
}
