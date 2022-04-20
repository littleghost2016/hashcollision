package hashcollision

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"

	"github.com/littleghost2016/hashcollision/errors"
)

// GetHashCode: 获取哈希值
func GetHashCode(hashType, message string) (string, error) {
	// messageB := s2b(message)
	messageB := []byte(message)

	var hashObject hash.Hash

	switch hashType {
	case "md5":
		hashObject = md5.New()
	case "sha1":
		hashObject = sha1.New()
	case "sha256":
		//创建一个基于SHA256算法的hash.Hash接口的对象
		hashObject = sha256.New()
	case "sha512":
		hashObject = sha512.New()
	default:
		return "", errors.ErrNotFoundHashFunction
	}

	//输入数据
	hashObject.Write(messageB)
	//计算哈希值
	hashB := hashObject.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(hashB)

	return hashCode, nil
}

func CompareHashCode(inString, expectedString string) bool {
	return inString[:6] == expectedString
}
