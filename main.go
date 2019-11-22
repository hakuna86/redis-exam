package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/hakuna86/redis-exam/client"
	"runtime"
)

type Store struct {
	Key   string
	Value interface{}
}

func GetKeys(cli *client.Client, keys []Store) {
	for _, v := range keys {
		value, err := cli.Get(v.Key)
		if err != nil {
			if err == redis.Nil {
				fmt.Printf("Key = %s not exist \n", v.Key)
				continue
			} else {
				panic(err)
			}
		}
		fmt.Printf("Get Key => %s, value => %s \n", v.Key, value)
	}
}

func SetKeys(cli *client.Client, keys []Store) {
	for _, v := range keys {
		err := cli.Set(v.Key, v.Value, 0)
		if err != nil {
			fmt.Printf("Set key %t : key => %s, value => %v \n", err, v.Key, v.Value)
			continue
		}
	}
}

func SetKeysNX(cli *client.Client, keys []Store) {
	for _, v := range keys {
		ok, err := cli.SetNotExist(v.Key, v.Value, 0)
		if err != nil {
			fmt.Printf("SetKeysNX Error %t : key => %s, value => %v \n", err, v.Key, v.Value)
			continue
		}
		if !ok {
			fmt.Printf("SetKeysNX NOT OK : key => %s, value => %v \n", v.Key, v.Value)
		}
	}
}

func SetKeyMustExist(cli *client.Client, keys []Store) {
	for _, v := range keys {
		ok, err := cli.SetKeyMustExist(v.Key, v.Value, 0)
		if err != nil {
			fmt.Printf("SetKeyMustExist Error %t : key => %s, value => %v \n", err, v.Key, v.Value)
			continue
		}
		if !ok {
			fmt.Printf("SetKeyMustExist NOT OK : key => %s, value => %v \n", v.Key, v.Value)
		}
	}
}

func Publish(cli *client.Client, ch, id string) {
	cnt := 0
	for {
		if err := cli.Publish(ch, fmt.Sprintf("%s=%s", id, cnt)); err != nil {
			fmt.Println("Publish Error", err)
			return
		}
		cnt++
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cli := client.NewClient()
	defer cli.Close()

	if err := cli.Ping(); err != nil {
		panic(err)
	}

	keys := []Store{
		{"key", "hello"}, // ok
		{"key", 100},     // ok
		//{"key-float", 10.095}, // ok
		//{"key-byte", []byte("Hello world")}, // ok
		//{"key-array", []string{"1", "2", "3"}}, // not ok
		//{"key-map", map[string]int{
		// "one" : 1,
		// "two" : 2,
		// "three" : 3,
		//}}, // not ok
	}

	//SetKeys(cli, keys)
	//GetKeys(cli, keys)
	//SetKeysNX(cli, keys)
	//SetKeyMustExist(cli, keys)

	_ = keys

	//k := Store{"key-map", map[string]interface{}{
	//	"one" : 1,
	//	"two" : 2,
	//	"three" : 3,
	//}}
	//if err := cli.SetMap(k.Key, k.Value.(map[string]interface{})); err != nil {
	//	fmt.Println("Set Map Error", err)
	//}

	//v, err := cli.GetMap("key-map")
	//if err != nil {
	//	fmt.Println("Set Map Error", err)
	//}
	//for key, value := range v {
	//	fmt.Printf("Get Map Value Key=> %s, Value=> %s \n", key, value)
	//}

	//sli := Store{"key-array", []string{"1", "2", "3"}}
	//switch t := sli.Value.(type) {
	//case []string:
	//	var data []interface{}
	//	for _, v := range t {
	//		data = append(data, v)
	//	}
	//	if err := cli.SetSlice(sli.Key, data...); err != nil {
	//		fmt.Println("Set Slice Error", err)
	//	}
	//}

	//res, err := cli.GetSlice("key-array")
	//if err != nil {
	//	fmt.Println("Get Slice Error", err)
	//}
	//fmt.Println(res)

	//start := time.Now().String()
	//ch := "test-ch"
	//go func() {
	//	for {
	//		select {
	//			case msg := <- cli.Subscribe(ch):
	//				fmt.Println("Receive message", msg.String())
	//		}
	//	}
	//}()
	//
	//for i := 0; i < 7; i++ {
	//	go Publish(cli, ch, fmt.Sprintf("node%d", i+1))
	//}
	//
	//fmt.Println("=========End=========", "start", start, "end", time.Now().String())
	//<- time.NewTimer(10 * time.Second).C

}
