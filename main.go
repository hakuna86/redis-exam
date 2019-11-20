package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/hakuna86/redis-exam/client"
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

func main() {
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
	SetKeyMustExist(cli, keys)
}
