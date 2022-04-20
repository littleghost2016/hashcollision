package main

import (
	"fmt"
	"sync"
	// "time"

	"github.com/littleghost2016/hashcollision"
	"github.com/littleghost2016/hashcollision/redis"
	"github.com/panjf2000/ants/v2"
)

const (
	expectedString = "6b51ea"
)

func main() {
	keyName := "key"
	// value, _ := redis.RPop(keyName)
	// fmt.Println(value, hashCode)

	var wg sync.WaitGroup

	p, _ := ants.NewPoolWithFunc(50, func(value interface{}) {
		myFunc(value)
		wg.Done()
	})
	defer p.Release()

	inString, _ := redis.RPop(keyName)
	for inString != "" {
		wg.Add(1)
		_ = p.Invoke(inString)
		inString, _ = redis.RPop(keyName)
	}

	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	// fmt.Printf("finish all tasks, result is %d\n", sum)
}

func myFunc(value interface{}) {
	message := value.(string)
	hashCode, _ := hashcollision.GetHashCode("sha256", message)
	if equal := hashcollision.CompareHashCode(hashCode, expectedString); equal {
		fmt.Println(message, hashCode)
	}
}
