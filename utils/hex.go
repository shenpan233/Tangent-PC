package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GetRandomBin(i int) (data []byte) {
	rand.Seed(time.Now().Unix())
	for i2 := 0; i2 < i; i2++ {
		data = append(data, byte(rand.Intn(255)))
	}
	return
}

func GetRand32() int32 {
	rand.Seed(GetServerCurTime())
	return int32(rand.Int())
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
	r := regexp.MustCompile("([0-9a-fA-F]+)")
	data := r.FindAllStringSubmatch(HexData, -1)
	tmp := ""
	for _, datum := range data {
		tmp += datum[0]
	}
	decodeString, _ := hex.DecodeString(tmp)
	return decodeString
}

func BinToHex(Bin []byte) string {
	return strings.ToUpper(hex.EncodeToString(Bin))
}

func BinToHex2(Bin *[]byte) string {
	return strings.ToUpper(hex.EncodeToString(*Bin))
}

func ToMd5Bytes(data []byte) (ret []byte) {
	tmp := md5.Sum(data)
	ret = tmp[:]
	return
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

func GetCrc32(data []byte) []byte {
	ieee := crc32.ChecksumIEEE(data)
	return []byte{byte(ieee >> 0), byte(ieee >> 8), byte(ieee >> 16), byte(ieee >> 24)}
}
