package util

// 根据项目 https://github.com/sun8911879/qqtea 的改进

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math/rand"
	"time"
)

const (
	delta   = uint32(0x9E3779B9)
	fillNor = 0xF8
)

type teaCipher struct {
	keys       [4]uint32
	value      []byte
	byte8      [8]byte
	uByte32    [2]uint32
	xor        [8]byte
	fXor       [8]byte
	lXor       [8]byte
	nXor       [8]byte
	baleBuffer *bytes.Buffer
	seedRand   *rand.Rand
}

func NewCipher(key []byte) *teaCipher {
	if len(key) != 16 {
		panic(errors.New("invalid key size error"))
		return nil
	}
	cipher := &teaCipher{
		baleBuffer: bytes.NewBuffer(nil),
	}
	for i := 0; i < 4; i++ {
		cipher.keys[i] = binary.BigEndian.Uint32(key[i*4:])
	}
	cipher.seedRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	return cipher
}

func (c *teaCipher) Encrypt(value []byte) []byte {
	c.baleBuffer.Reset()
	vl := len(value)
	filLn := (8 - (vl + 2)) % 8
	if filLn < 0 {
		filLn += 2 + 8
	} else {
		filLn += 2
	}
	binDex := filLn + 1
	if binDex <= 0 {
		return nil
	}
	rands := make([]byte, binDex)
	for i := 1; i < binDex; i++ {
		rands[i] = byte(c.seedRand.Int() & fillNor)
	}
	rands[0] = byte((filLn - 2) | fillNor)
	c.baleBuffer.Write(rands)
	c.baleBuffer.Write(value)
	c.baleBuffer.Write([]byte{00, 00, 00, 00, 00, 00, 00})
	vl = c.baleBuffer.Len()
	c.value = c.baleBuffer.Bytes()
	c.baleBuffer.Reset()
	for i := 0; i < vl; i += 8 {
		c.xor = xor(c.value[i:i+8], c.fXor[0:8])
		c.uByte32[0] = binary.BigEndian.Uint32(c.xor[0:4])
		c.uByte32[1] = binary.BigEndian.Uint32(c.xor[4:8])
		c.encipher()
		c.fXor = xor(c.byte8[0:8], c.lXor[0:8])
		c.baleBuffer.Write(c.fXor[0:8])
		c.lXor = c.xor

	}
	return c.baleBuffer.Bytes()
}

func (c *teaCipher) Decrypt(value []byte) []byte {
	vl := len(value)
	if vl <= 0 || (vl%8) != 0 {
		return nil
	}
	c.baleBuffer.Reset()
	c.uByte32[0] = binary.BigEndian.Uint32(value[0:4])
	c.uByte32[1] = binary.BigEndian.Uint32(value[4:8])
	c.decipher()
	copy(c.lXor[0:8], value[0:8])
	c.fXor = c.byte8
	pos := int((c.byte8[0] & 0x7) + 2)
	c.baleBuffer.Write(c.byte8[0:8])
	for i := 8; i < vl; i += 8 {
		c.xor = xor(value[i:i+8], c.fXor[0:8])
		c.uByte32[0] = binary.BigEndian.Uint32(c.xor[0:4])
		c.uByte32[1] = binary.BigEndian.Uint32(c.xor[4:8])
		c.decipher()
		c.nXor = xor(c.byte8[0:8], c.lXor[0:8])
		c.baleBuffer.Write(c.nXor[0:8])
		c.fXor = xor(c.nXor[0:8], c.lXor[0:8])
		copy(c.lXor[0:8], value[i:i+8])
	}
	pos++
	c.value = c.baleBuffer.Bytes()
	nl := c.baleBuffer.Len()
	if pos >= c.baleBuffer.Len() || (nl-7) <= pos {
		return nil
	}
	return c.value[pos : nl-7]
}

func (c *teaCipher) encipher() {
	sum := delta
	for i := 0x10; i > 0; i-- {
		c.uByte32[0] += ((c.uByte32[1] << 4 & 0xFFFFFFF0) + c.keys[0]) ^ (c.uByte32[1] + sum) ^ ((c.uByte32[1] >> 5 & 0x07ffffff) + c.keys[1])
		c.uByte32[1] += ((c.uByte32[0] << 4 & 0xFFFFFFF0) + c.keys[2]) ^ (c.uByte32[0] + sum) ^ ((c.uByte32[0] >> 5 & 0x07ffffff) + c.keys[3])
		sum += delta
	}
	binary.BigEndian.PutUint32(c.byte8[0:4], c.uByte32[0])
	binary.BigEndian.PutUint32(c.byte8[4:8], c.uByte32[1])
}

func (c *teaCipher) decipher() {
	sum := delta
	sum = (sum << 4) & 0xffffffff

	for i := 0x10; i > 0; i-- {
		c.uByte32[1] -=
			((c.uByte32[0] << 4 & 0xFFFFFFF0) + c.keys[2]) ^ (c.uByte32[0] + sum) ^ ((c.uByte32[0] >> 5 & 0x07ffffff) + c.keys[3])
		c.uByte32[1] &= 0xffffffff
		c.uByte32[0] -=
			((c.uByte32[1] << 4 & 0xFFFFFFF0) + c.keys[0]) ^ (c.uByte32[1] + sum) ^ ((c.uByte32[1] >> 5 & 0x07ffffff) + c.keys[1])
		c.uByte32[0] &= 0xffffffff
		sum -= delta
	}
	binary.BigEndian.PutUint32(c.byte8[0:4], c.uByte32[0])
	binary.BigEndian.PutUint32(c.byte8[4:8], c.uByte32[1])
}

func xor(a, b []byte) (bts [8]byte) {
	l := len(a)
	for i := 0; i < l; i += 4 {
		binary.BigEndian.PutUint32(bts[i:i+4], binary.BigEndian.Uint32(a[i:i+4])^binary.BigEndian.Uint32(b[i:i+4]))
	}
	return bts
}

func Encrypt(key, bin []byte) []byte {
	tea := NewCipher(key)
	return tea.Encrypt(bin)
}

func Decrypt(key, bin []byte) []byte {
	tea := NewCipher(key)
	return tea.Decrypt(bin)
}
