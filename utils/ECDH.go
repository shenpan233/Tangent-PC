/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/4 13:23
  @Notice:  ECDH
*/

package util

import (
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
)

func GenECDHKey() (PublicKey, ShareKey []byte) {
	pub := []byte{4, 191, 71, 161, 207, 120, 166, 41, 102, 139, 11, 195, 159, 142, 84, 201, 204, 243, 182, 56, 75, 8, 184, 174, 236, 135, 218, 159, 48, 72, 94, 223, 231, 103, 150, 157, 193, 163, 175, 17, 21, 254, 13, 204, 142, 11, 23, 202, 207}
	curve := elliptic.P256()
	key, sx, sy, _ := elliptic.GenerateKey(curve, rand.Reader)
	tx, ty := elliptic.Unmarshal(curve, pub)
	x, _ := curve.ScalarMult(tx, ty, key)
	hash := md5.Sum(x.Bytes()[:16])
	ShareKey = hash[:]
	PublicKey = elliptic.Marshal(curve, sx, sy)
	return
}
