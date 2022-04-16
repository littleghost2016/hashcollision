package hashcollision

import (
	"testing"
	
	"github.com/littleghost2016/hashcollision/errors"
	"github.com/stretchr/testify/assert"
)

type HashTestCase struct {
	hashType          string
	correctTestResult string
	err               error
}

func TestGetHashCode(t *testing.T) {
	myAssert := assert.New(t)

	message := "123"

	testCase := []HashTestCase{
		{
			"md5",
			"202cb962ac59075b964b07152d234b70",
			nil,
		},
		{
			"sha1",
			"40bd001563085fc35165329ea1ff5c5ecbdbbeef",
			nil,
		},
		{
			"sha256",
			"a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3",
			nil,
		},
		{
			"sha512",
			"3c9909afec25354d551dae21590bb26e38d53f2173b8d3dc3eee4c047e7ab1c1eb8b85103e3be7ba613b31bb5c9c36214dc9f14a42fd7a2fdb84856bca5c44c2",
			nil,
		},
		{
			"someHashFunction",
			"321",
			errors.NotFoundHashFunctionError,
		},
	}

	for _, eachTestCase := range testCase {
		result, err := GetHashCode(eachTestCase.hashType, message)
		if err != nil {
			myAssert.EqualError(errors.NotFoundHashFunctionError, err.Error())
		} else {
			myAssert.Equalf(eachTestCase.correctTestResult, result, "GetHashCode %s 结果错误", eachTestCase.hashType)
		}
	}
}
