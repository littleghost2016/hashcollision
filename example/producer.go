package main

import (
	// "strconv"
	// "time"
	// "github.com/littleghost2016/hashcollision"
	"github.com/littleghost2016/hashcollision/redis"
)

// timeCost 耗时统计
// func timeCost() func() {
// 	start := time.Now()
// 	return func() {
// 		tc := time.Since(start)
// 		fmt.Printf("time cost = %v\n", tc)
// 	}
// }

// func main() {
// 	defer timeCost()() // 注意，是对 timeCost() 返回的函数进行延迟调用，因此需要加两对小括号
// 	// _ = redis.RPush("key", 1)
// 	for i := 0; i < 1000000; i++ {
// 		eachHash, _ := hashcollision.GetHashCode("sha1", strconv.Itoa(i))
// 		// fmt.Println(i, eachHash)
// 		_ = redis.LPush("key", eachHash)
// 	}
// }

func main() {
	keyName := "key"
	baseString := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	length := 4

	// 初始化location_counter
	locationCounter := make([]int, length)
	for i := 9; i < length; i++ {
		locationCounter = append(locationCounter, 0)
	}

	overflowFlag := false
	for !overflowFlag {
		resultString := getStringFromLocationCounter(locationCounter, baseString)
		// fmt.Println(resultString)
		redis.LPush(keyName, resultString)
		overflowFlag = locationCounterAddOne(&locationCounter, baseString)
	}

}

func getStringFromLocationCounter(locationCounter []int, baseString string) string {
	inByteSlice := []byte(baseString)
	result := make([]byte, len(locationCounter))
	for index := range result {
		result[index] = inByteSlice[locationCounter[index]]
	}

	return string(result)
}

func reverseByteSlice(in *[]byte) {
	for i, j := 0, len(*in)-1; i < j; i, j = i+1, j-1 {
		(*in)[i], (*in)[j] = (*in)[j], (*in)[i]
	}
}

func locationCounterAddOne(locationCounter *[]int, baseString string) bool {
	overflowFlag := false
	carryFlag := false
	locationCounterLength := len(*locationCounter)
	baseStringLength := len(baseString)

	// 最后一个位置+1
	// 如果最后一位+1后需要进位
	if (*locationCounter)[locationCounterLength-1] == baseStringLength-1 {
		// 设置前一位的进位标志
		carryFlag = true
		// 将最后一位置为0
		(*locationCounter)[locationCounterLength-1] = 0
		// 判断除了第一位和最后一位的所有位置是否需要+1
		for i := locationCounterLength - 2; i >= 0; i-- {
			if carryFlag {
				if (*locationCounter)[i] == baseStringLength-1 {
					carryFlag = true
					(*locationCounter)[i] = 0
				} else {
					carryFlag = false
					(*locationCounter)[i]++
				}
			} else {
				break
			}
		}
		// 最后一位+1后不需要进位
	} else {
		(*locationCounter)[locationCounterLength-1]++
	}

	if carryFlag {
		overflowFlag = true
	}

	return overflowFlag
}
