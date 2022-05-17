package util

import (
	"crypto/md5"
)

// this is password encrypto style
//return string which is decoded

func MD5EnCrypto(password string) string {
	sum := md5.Sum([]byte(password))
	sliceSum := sum[:]
	//fmt.Println(string(sliceSum))
	return string(sliceSum)
}
